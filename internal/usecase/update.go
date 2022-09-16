package usecase

import (
	"context"
	"fmt"
	"io"
	"path"
	"time"

	"github.com/TcMits/ent-clean-template/ent"
	"github.com/TcMits/ent-clean-template/internal/repository"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/TcMits/ent-clean-template/pkg/tool/generic"
)

var (
	_wrapStartUpdateTransactionError = func(err error) error {
		return useCaseModel.NewUseCaseError(
			fmt.Errorf("updateModelInTransactionUseCase - Update - u.transactionRepository.Start: %w", err),
			"internal.usecase.update.updateModelInTransactionUseCase.Update.StartUpdateTransactionError",
			"Can't update now",
			DBError,
		)
	}
	_wrapCommitUpdateError = func(err error) error {
		return useCaseModel.NewUseCaseError(
			fmt.Errorf("updateModelInTransactionUseCase - Update - u.transactionRepository.Commit: %w", err),
			"internal.usecase.update.updateModelInTransactionUseCase.Update.CommitUpdateError",
			"Can't update now",
			DBError,
		)
	}
	_wrapRollbackUpdateError = func(err error) error {
		return useCaseModel.NewUseCaseError(
			fmt.Errorf("updateModelInTransactionUseCase - Update - u.transactionRepository.Rollback: %w", err),
			"internal.usecase.update.updateModelInTransactionUseCase.Update.RollbackUpdateError",
			"Can't update now",
			DBError,
		)
	}
)

type UpdateFile struct {
	Filename string
	Size     int64
	Reader   io.Reader
}

type UpdateValidateFunc[ModelType, UpdateInputType, RepoUpdateInputType any] func(context.Context, ModelType, UpdateInputType) (RepoUpdateInputType, error)
type UpdateInTransactionValidateFunc[ModelType, UpdateInputType, RepoUpdateInputType any] func(context.Context, ModelType, UpdateInputType, *ent.Client) (RepoUpdateInputType, error)
type UpdateWithFileValidateFunc[UpdateInputType, UseCaseUpdateInputType any] func(context.Context, UpdateInputType, UpdateExistFunc) (UseCaseUpdateInputType, []*UpdateFile, error)
type UpdateExistFunc func(context.Context, string) (bool, error)

type updateModelUseCase[ModelType, WhereInputType, UpdateInputType, RepoWhereInputType, RepoUpdateInputType any] struct {
	repository           repository.UpdateModelRepository[ModelType, RepoUpdateInputType]
	getRepository        repository.GetModelRepository[ModelType, RepoWhereInputType]
	toRepoWhereInputFunc ConverFunc[WhereInputType, RepoWhereInputType]
	validateFunc         UpdateValidateFunc[ModelType, UpdateInputType, RepoUpdateInputType]
	wrapGetErrorFunc     func(error) error
	wrapUpdateErrorFunc  func(error) error
}

type updateModelInTransactionUseCase[ModelType, WhereInputType, UpdateInputType, RepoWhereInputType, RepoUpdateInputType any] struct {
	repository            repository.UpdateWithClientModelRepository[ModelType, RepoUpdateInputType]
	getRepository         repository.GetWithClientModelRepository[ModelType, RepoWhereInputType]
	transactionRepository repository.TransactionRepository
	toRepoWhereInputFunc  ConverFunc[WhereInputType, RepoWhereInputType]
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
	basePath            string
}

func (u *updateModelUseCase[ModelType, WhereInputType, UpdateInputType, _, _]) GetAndUpdate(
	ctx context.Context, whereInput WhereInputType, updateInput UpdateInputType,
) (ModelType, error) {
	repoWhereInput, err := u.toRepoWhereInputFunc(ctx, whereInput)
	if err != nil {
		return generic.Zero[ModelType](), err
	}
	instance, err := u.getRepository.Get(ctx, repoWhereInput)
	if err != nil {
		return generic.Zero[ModelType](), u.wrapGetErrorFunc(err)
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
	ctx context.Context, whereInput WhereInputType, updateInput UpdateInputType,
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
	oldInstance, err := u.getRepository.GetWithClient(ctx, client, repoWhereInput, u.selectForUpdate)
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
	ctx context.Context, whereInput WhereInputType, updateInput UpdateInputType,
) (ModelType, error) {
	useCaseUpdateInput, files, err := u.validateFunc(ctx, updateInput, u.existFileRepository.Exist)
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
					fileCtx, path.Join(u.basePath, file.Filename), file.Reader, file.Size,
				)
				if fileErr != nil {
					u.l.Error(fileErr)
					continue
				}
				u.l.Info("getAndUpdateModelWithFileUseCase - GetAndUpdate - u.getAndUpdateUseCase.GetAndUpdate: Upload %d bytes", n)
			}
		}()
	}
	return instance, nil
}
