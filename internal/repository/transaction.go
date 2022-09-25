package repository

import (
	"context"

	"github.com/TcMits/ent-clean-template/ent"
)

type transactionRepository struct {
	client *ent.Client
}

func NewTransactionRepository(client *ent.Client) TransactionRepository {
	if client == nil {
		panic("client is required")
	}
	return &transactionRepository{client: client}
}

func (r *transactionRepository) Start(
	ctx context.Context,
) (*ent.Client, func() error, func() error, error) {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, nil, nil, err
	}
	commitFunc := func() error {
		return tx.Commit()
	}
	rollbackFunc := func() error {
		return tx.Rollback()
	}
	return tx.Client(), commitFunc, rollbackFunc, nil
}
