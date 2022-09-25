package usecase

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/TcMits/ent-clean-template/ent"
	"github.com/TcMits/ent-clean-template/internal/repository"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/TcMits/ent-clean-template/pkg/tool/generic"
)

var (
	_wrapStartCreateTransactionError = func(err error) error {
		return useCaseModel.NewUseCaseError(
			fmt.Errorf(
				"createModelInTransactionUseCase - Create - u.transactionRepository.Start: %w",
				err,
			),
			"internal.usecase.create.createModelInTransactionUseCase.Create.StartCreateTransactionError",
			"Can't create now",
			DBError,
		)
	}
	_wrapCommitCreateError = func(err error) error {
		return useCaseModel.NewUseCaseError(
			fmt.Errorf(
				"createModelInTransactionUseCase - Create - u.transactionRepository.Commit: %w",
				err,
			),
			"internal.usecase.create.createModelInTransactionUseCase.Create.CommitCreateError",
			"Can't create now",
			DBError,
		)
	}
	_wrapRollbackCreateError = func(err error) error {
		return useCaseModel.NewUseCaseError(
			fmt.Errorf(
				"createModelInTransactionUseCase - Create - u.transactionRepository.Rollback: %w",
				err,
			),
			"internal.usecase.create.createModelInTransactionUseCase.Create.RollbackCreateError",
			"Can't create now",
			DBError,
		)
	}
)

type CreateFile struct {
	Filename string
	Size     int64
	Reader   io.Reader
}

type (
	CreateValidateFunc[CreateInputType, RepoCreateInputType any]              func(context.Context, CreateInputType) (RepoCreateInputType, error)
	CreateInTransactionValidateFunc[CreateInputType, RepoCreateInputType any] func(context.Context, CreateInputType, *ent.Client) (RepoCreateInputType, error)
	CreateWithFileValidateFunc[CreateInputType, UseCaseCreateInputType any]   func(context.Context, CreateInputType) (UseCaseCreateInputType, []*CreateFile, error)
)

type createModelUseCase[ModelType, CreateInputType, RepoCreateInputType any] struct {
	repository          repository.CreateModelRepository[ModelType, RepoCreateInputType]
	validateFunc        CreateValidateFunc[CreateInputType, RepoCreateInputType]
	wrapCreateErrorFunc func(error) error
}

type createModelInTransactionUseCase[ModelType, CreateInputType, RepoCreateInputType any] struct {
	repository            repository.CreateWithClientModelRepository[ModelType, RepoCreateInputType]
	transactionRepository repository.TransactionRepository
	validateFunc          CreateInTransactionValidateFunc[CreateInputType, RepoCreateInputType]
	wrapCreateErrorFunc   func(error) error
}

type createModelHavingFileUseCase[ModelType, CreateInputType, UseCaseCreateInputType any] struct {
	createUseCase       CreateModelUseCase[ModelType, UseCaseCreateInputType]
	existFileRepository repository.ExistFileRepository
	writeFileRepository repository.WriteFileRepository
	validateFunc        CreateWithFileValidateFunc[CreateInputType, UseCaseCreateInputType]
	writeFileTimeout    time.Duration
	l                   logger.Interface
}

func (u *createModelUseCase[ModelType, CreateInputType, _]) Create(
	ctx context.Context, input CreateInputType,
) (ModelType, error) {
	repoCreateInput, err := u.validateFunc(ctx, input)
	if err != nil {
		return generic.Zero[ModelType](), err
	}
	instance, err := u.repository.Create(ctx, repoCreateInput)
	if err != nil {
		return generic.Zero[ModelType](), u.wrapCreateErrorFunc(err)
	}
	return instance, nil
}

func (u *createModelInTransactionUseCase[ModelType, CreateInputType, _]) Create(
	ctx context.Context, input CreateInputType,
) (instance ModelType, err error) {
	client, commitFunc, rollbackFunc, err := u.transactionRepository.Start(ctx)
	if err != nil {
		err = _wrapStartCreateTransactionError(err)
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
				err = _wrapRollbackCreateError(rerr)
			}
			return
		}
		if cerr := commitFunc(); cerr != nil {
			err = _wrapCommitCreateError(cerr)
			instance = generic.Zero[ModelType]()
		}
	}()

	// validate input
	repoCreateInput, err := u.validateFunc(ctx, input, client)
	if err != nil {
		return
	}

	// create instance
	instance, err = u.repository.CreateWithClient(ctx, client, repoCreateInput)
	if err != nil {
		err = u.wrapCreateErrorFunc(err)
		return
	}
	return
}

func (u *createModelHavingFileUseCase[ModelType, CreateInputType, _]) Create(
	ctx context.Context, input CreateInputType,
) (ModelType, error) {
	createInput, files, err := u.validateFunc(ctx, input)
	if err != nil {
		return generic.Zero[ModelType](), err
	}
	instance, err := u.createUseCase.Create(ctx, createInput)
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
					"createModelHavingFileUseCase - Create - u.writeFileRepository.Write: Upload %d bytes",
					n,
				)
			}
		}()
	}
	return instance, nil
}
