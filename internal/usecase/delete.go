package usecase

import (
	"context"
	"fmt"

	"github.com/TcMits/ent-clean-template/internal/repository"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var (
	// i18n messages
	_startDeleteTransactionErrorMsg = &i18n.Message{
		ID:    "internal.usecase.delete.deleteModelInTransactionUseCase.Delete.StartDeleteTransactionError",
		Other: "Can't delete now",
	}
	_commitDeleteErrorMsg = &i18n.Message{
		ID:    "internal.usecase.delete.deleteModelInTransactionUseCase.Delete.CommitDeleteError",
		Other: "Can't delete now",
	}
	_rollbackDeleteErrorMsg = &i18n.Message{
		ID:    "internal.usecase.delete.deleteModelInTransactionUseCase.Delete.RollbackDeleteError",
		Other: "Can't delete now",
	}

	// wrap error
	_wrapStartDeleteTransactionError = func(err error) error {
		return model.NewTranslatableError(
			fmt.Errorf(
				"deleteModelInTransactionUseCase - Delete - u.transactionRepository.Start: %w",
				err,
			),
			_startDeleteTransactionErrorMsg,
			DBError,
			nil,
		)
	}
	_wrapCommitDeleteError = func(err error) error {
		return model.NewTranslatableError(
			fmt.Errorf(
				"deleteModelInTransactionUseCase - Delete - u.transactionRepository.Commit: %w",
				err,
			),
			_commitDeleteErrorMsg,
			DBError,
			nil,
		)
	}
	_wrapRollbackDeleteError = func(err error) error {
		return model.NewTranslatableError(
			fmt.Errorf(
				"deleteModelInTransactionUseCase - Delete - u.transactionRepository.Rollback: %w",
				err,
			),
			_rollbackDeleteErrorMsg,
			DBError,
			nil,
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
