package usecase

import (
	"context"
	"fmt"

	"github.com/TcMits/ent-clean-template/ent"
	"github.com/TcMits/ent-clean-template/internal/repository"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
	"github.com/TcMits/ent-clean-template/pkg/tool/generic"
)

var (
	_wrapStartCreateTransactionError = func(err error) error {
		return useCaseModel.NewUseCaseError(
			fmt.Errorf("createModelInTransactionUseCase - Create - u.transactionRepository.Start: %w", err),
			"internal.usecase.create.createModelInTransactionUseCase.Create.StartCreateTransactionError",
			"Can't create now",
			DBError,
		)
	}
	_wrapCommitCreateError = func(err error) error {
		return useCaseModel.NewUseCaseError(
			fmt.Errorf("createModelInTransactionUseCase - Create - u.transactionRepository.Commit: %w", err),
			"internal.usecase.create.createModelInTransactionUseCase.Create.CommitCreateError",
			"Can't create now",
			DBError,
		)
	}
	_wrapRollbackCreateError = func(err error) error {
		return useCaseModel.NewUseCaseError(
			fmt.Errorf("createModelInTransactionUseCase - Create - u.transactionRepository.Rollback: %w", err),
			"internal.usecase.create.createModelInTransactionUseCase.Create.RollbackCreateError",
			"Can't create now",
			DBError,
		)
	}
)

type CreateValidateFunc[CreateInput, RepoCreateInput any] func(CreateInput) (RepoCreateInput, error)
type CreateInTransactionValidateFunc[CreateInput, RepoCreateInput any] func(CreateInput, *ent.Client) (RepoCreateInput, error)

type createModelUseCase[ModelType, CreateInput, RepoCreateInput any] struct {
	repository          repository.CreateModelRepository[ModelType, RepoCreateInput]
	validateFunc        CreateValidateFunc[CreateInput, RepoCreateInput]
	wrapCreateErrorFunc func(error) error
}

type createModelInTransactionUseCase[ModelType, CreateInput, RepoCreateInput any] struct {
	repository            repository.CreateWithClientModelRepository[ModelType, RepoCreateInput]
	transactionRepository repository.TransactionRepository
	validateFunc          CreateInTransactionValidateFunc[CreateInput, RepoCreateInput]
	wrapCreateErrorFunc   func(error) error
}

func (u *createModelUseCase[ModelType, CreateInput, RepoCreateInput]) Create(
	ctx context.Context, input CreateInput,
) (ModelType, error) {
	repoCreateInput, err := u.validateFunc(input)
	if err != nil {
		return generic.Zero[ModelType](), err
	}
	instance, err := u.repository.Create(ctx, repoCreateInput)
	if err != nil {
		return generic.Zero[ModelType](), u.wrapCreateErrorFunc(err)
	}
	return instance, nil
}

func (u *createModelInTransactionUseCase[ModelType, CreateInput, RepoCreateInput]) Create(
	ctx context.Context, input CreateInput,
) (ModelType, error) {
	tx, err := u.transactionRepository.Start(ctx)
	if err != nil {
		return generic.Zero[ModelType](), _wrapStartCreateTransactionError(err)
	}
	client := tx.Client()

	// validate input
	repoCreateInput, err := u.validateFunc(input, client)
	if err != nil {
		if rerr := u.transactionRepository.Rollback(tx); rerr != nil {
			return generic.Zero[ModelType](), _wrapRollbackCreateError(rerr)
		}
		return generic.Zero[ModelType](), err
	}

	// create instance
	instance, err := u.repository.CreateWithClient(ctx, client, repoCreateInput)
	if err != nil {
		if rerr := u.transactionRepository.Rollback(tx); rerr != nil {
			return generic.Zero[ModelType](), _wrapRollbackCreateError(rerr)
		}
		return generic.Zero[ModelType](), u.wrapCreateErrorFunc(err)
	}

	if err = u.transactionRepository.Commit(tx); err != nil {
		return generic.Zero[ModelType](), _wrapCommitCreateError(err)
	}
	return instance, nil
}
