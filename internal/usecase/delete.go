package usecase

import (
	"context"
	"fmt"

	"github.com/TcMits/ent-clean-template/internal/repository"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
)

var (
	_wrapStartDeleteTransactionError = func(err error) error {
		return useCaseModel.NewUseCaseError(
			fmt.Errorf(
				"deleteModelInTransactionUseCase - Delete - u.transactionRepository.Start: %w",
				err,
			),
			"internal.usecase.delete.deleteModelInTransactionUseCase.Delete.StartDeleteTransactionError",
			"Can't delete now",
			DBError,
		)
	}
	_wrapCommitDeleteError = func(err error) error {
		return useCaseModel.NewUseCaseError(
			fmt.Errorf(
				"deleteModelInTransactionUseCase - Delete - u.transactionRepository.Commit: %w",
				err,
			),
			"internal.usecase.delete.deleteModelInTransactionUseCase.Delete.CommitDeleteError",
			"Can't delete now",
			DBError,
		)
	}
	_wrapRollbackDeleteError = func(err error) error {
		return useCaseModel.NewUseCaseError(
			fmt.Errorf(
				"deleteModelInTransactionUseCase - Delete - u.transactionRepository.Rollback: %w",
				err,
			),
			"internal.usecase.delete.deleteModelInTransactionUseCase.Delete.RollbackDeleteError",
			"Can't delete now",
			DBError,
		)
	}
)

type deleteModelUseCase[ModelType, WhereInputType, RepoWhereInputType any] struct {
	getUseCase          GetModelUseCase[ModelType, WhereInputType]
	repository          repository.DeleteModelRepository[ModelType]
	wrapDeleteErrorFunc func(error) error
}

type deleteModelInTransactionUseCase[ModelType, WhereInputType, RepoWhereInputType any] struct {
	repository            repository.DeleteWithClientModelRepository[ModelType]
	getRepository         repository.GetWithClientModelRepository[ModelType, RepoWhereInputType]
	transactionRepository repository.TransactionRepository
	toRepoWhereInputFunc  ConverFunc[WhereInputType, RepoWhereInputType]
	selectForUpdate       bool
	wrapGetErrorFunc      func(error) error
	wrapDeleteErrorFunc   func(error) error
}

func (u *deleteModelUseCase[ModelType, WhereInputType, _]) GetAndDelete(
	ctx context.Context, input WhereInputType,
) error {
	instance, err := u.getUseCase.Get(ctx, input)
	if err != nil {
		return err
	}
	err = u.repository.Delete(ctx, instance)
	if err != nil {
		return u.wrapDeleteErrorFunc(err)
	}
	return nil
}

func (u *deleteModelInTransactionUseCase[ModelType, WhereInputType, _]) GetAndDelete(
	ctx context.Context, input WhereInputType,
) (err error) {
	client, commitFunc, rollbackFunc, err := u.transactionRepository.Start(ctx)
	if err != nil {
		err = _wrapStartDeleteTransactionError(err)
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
				err = _wrapRollbackDeleteError(rerr)
			}
			return
		}
		if cerr := commitFunc(); cerr != nil {
			err = _wrapCommitDeleteError(cerr)
		}
	}()

	// get whereInput
	repoWhereInput, err := u.toRepoWhereInputFunc(ctx, input)
	if err != nil {
		return
	}

	// get instance
	instance, err := u.getRepository.GetWithClient(ctx, client, repoWhereInput, u.selectForUpdate)
	if err != nil {
		err = u.wrapGetErrorFunc(err)
		return
	}

	// delete instance
	err = u.repository.DeleteWithClient(ctx, client, instance)
	if err != nil {
		err = u.wrapDeleteErrorFunc(err)
		return
	}
	return
}
