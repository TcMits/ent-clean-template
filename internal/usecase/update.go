package usecase

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/TcMits/ent-clean-template/ent"
	"github.com/TcMits/ent-clean-template/internal/repository"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/TcMits/ent-clean-template/pkg/tool/generic"
)

var (
	_wrapStartUpdateTransactionError = func(err error) error {
		return model.NewTranslatableError(
			fmt.Errorf(
				"internal.usecase.update.updateModelInTransactionUseCase.GetAndUpdate: %w",
				err,
			),
			_canNotUpdateNowMessage,
			DBError,
			nil,
		)
	}
	_wrapCommitUpdateError = func(err error) error {
		return model.NewTranslatableError(
			fmt.Errorf(
				"internal.usecase.update.updateModelInTransactionUseCase.GetAndUpdate: %w",
				err,
			),
			_canNotUpdateNowMessage,
			DBError,
			nil,
		)
	}
	_wrapRollbackUpdateError = func(err error) error {
		return model.NewTranslatableError(
			fmt.Errorf(
				"internal.usecase.update.updateModelInTransactionUseCase.GetAndUpdate: %w",
				err,
			),
			_canNotUpdateNowMessage,
			DBError,
			nil,
		)
	}
)

type UpdateFile struct {
	Filename string
	Size     int64
	Reader   io.Reader
}

type (
	UpdateValidateFunc[ModelType, UpdateInputType, RepoUpdateInputType any]              func(context.Context, ModelType, UpdateInputType) (RepoUpdateInputType, error)
	UpdateInTransactionValidateFunc[ModelType, UpdateInputType, RepoUpdateInputType any] func(context.Context, ModelType, UpdateInputType, *ent.Client) (RepoUpdateInputType, error)
	UpdateWithFileValidateFunc[UpdateInputType, UseCaseUpdateInputType any]              func(context.Context, UpdateInputType) (UseCaseUpdateInputType, []*UpdateFile, error)
)

type updateModelUseCase[ModelType, WhereInputType, UpdateInputType, RepoUpdateInputType any] struct {
	getUseCase          GetModelUseCase[ModelType, WhereInputType]
	repository          repository.UpdateModelRepository[ModelType, RepoUpdateInputType]
	validateFunc        UpdateValidateFunc[ModelType, UpdateInputType, RepoUpdateInputType]
	wrapUpdateErrorFunc func(error) error
}

type updateModelInTransactionUseCase[ModelType, WhereInputType, UpdateInputType, RepoWhereInputType, RepoUpdateInputType any] struct {
	repository            repository.UpdateWithClientModelRepository[ModelType, RepoUpdateInputType]
	getRepository         repository.GetWithClientModelRepository[ModelType, RepoWhereInputType]
	transactionRepository repository.TransactionRepository
	toRepoWhereInputFunc  ConvertFunc[WhereInputType, RepoWhereInputType]
	validateFunc          UpdateInTransactionValidateFunc[ModelType, UpdateInputType, RepoUpdateInputType]
	selectForUpdate       bool
	wrapGetErrorFunc      func(error) error
	wrapUpdateErrorFunc   func(error) error
}

type getAndUpdateModelWithFileUseCase[ModelType, WhereInputType, UpdateInputType, UseCaseUpdateInputType any] struct {
	getAndUpdateUseCase GetAndUpdateModelUseCase[ModelType, WhereInputType, UseCaseUpdateInputType]
	existFileRepository repository.ExistFileRepository
	writeFileRepository repository.WriteFileRepository
	validateFunc        UpdateWithFileValidateFunc[UpdateInputType, UseCaseUpdateInputType]
	writeFileTimeout    time.Duration
	l                   logger.Interface
}

func (u *updateModelUseCase[ModelType, WhereInputType, UpdateInputType, _]) GetAndUpdate(
	ctx context.Context, whereInput WhereInputType, updateInput UpdateInputType,
) (ModelType, error) {
	instance, err := u.getUseCase.Get(ctx, whereInput)
	if err != nil {
		return generic.Zero[ModelType](), err
	}
	repoUpdateInput, err := u.validateFunc(ctx, instance, updateInput)
	if err != nil {
		return generic.Zero[ModelType](), err
	}
	instance, err = u.repository.Update(ctx, instance, repoUpdateInput)
	if err != nil {
		return generic.Zero[ModelType](), u.wrapUpdateErrorFunc(err)
	}
	return instance, nil
}

func (u *updateModelInTransactionUseCase[ModelType, WhereInputType, UpdateInputType, _, _]) GetAndUpdate(
	ctx context.Context,
	whereInput WhereInputType,
	updateInput UpdateInputType,
) (instance ModelType, err error) {
	client, commitFunc, rollbackFunc, err := u.transactionRepository.Start(ctx)
	if err != nil {
		err = _wrapStartUpdateTransactionError(err)
		return
	}

	// ensure rollback or commit
	defer func() {
		if r := recover(); r != nil {
			rollbackFunc()
			panic(r)
		}
		if err != nil {
			if rerr := rollbackFunc(); rerr != nil {
				err = _wrapRollbackUpdateError(rerr)
			}
			return
		}
		if cerr := commitFunc(); cerr != nil {
			err = _wrapCommitUpdateError(cerr)
			instance = generic.Zero[ModelType]()
		}
	}()

	// get whereInput
	repoWhereInput, err := u.toRepoWhereInputFunc(ctx, whereInput)
	if err != nil {
		return
	}

	// get instance
	oldInstance, err := u.getRepository.GetWithClient(
		ctx,
		client,
		repoWhereInput,
		u.selectForUpdate,
	)
	if err != nil {
		err = u.wrapGetErrorFunc(err)
		return
	}

	// validate input with old instance
	repoUpdateInput, err := u.validateFunc(ctx, oldInstance, updateInput, client)
	if err != nil {
		return
	}

	// update instance
	instance, err = u.repository.UpdateWithClient(ctx, client, oldInstance, repoUpdateInput)
	if err != nil {
		err = u.wrapUpdateErrorFunc(err)
		return
	}
	return
}

func (u *getAndUpdateModelWithFileUseCase[ModelType, WhereInputType, UpdateInputType, _]) GetAndUpdate(
	ctx context.Context,
	whereInput WhereInputType,
	updateInput UpdateInputType,
) (ModelType, error) {
	useCaseUpdateInput, files, err := u.validateFunc(ctx, updateInput)
	if err != nil {
		return generic.Zero[ModelType](), err
	}
	instance, err := u.getAndUpdateUseCase.GetAndUpdate(ctx, whereInput, useCaseUpdateInput)
	if err != nil {
		return generic.Zero[ModelType](), err
	}

	// files upload
	if len(files) > 0 {
		go func() {
			fileCtx, cancelFileCtx := context.WithTimeout(context.Background(), u.writeFileTimeout)
			defer cancelFileCtx()
			for _, file := range files {
				n, fileErr := u.writeFileRepository.Write(
					fileCtx, file.Filename, file.Reader, file.Size,
				)
				if fileErr != nil {
					u.l.Error(fileErr)
					continue
				}
				u.l.Info(
					"getAndUpdateModelWithFileUseCase - GetAndUpdate - u.getAndUpdateUseCase.GetAndUpdate: Upload %d bytes",
					n,
				)
			}
		}()
	}
	return instance, nil
}
