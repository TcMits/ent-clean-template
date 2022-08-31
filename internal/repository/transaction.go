package repository

import (
	"context"

	"github.com/TcMits/ent-clean-template/ent"
)

var _ TransactionRepository = &transactionRepository{}

type transactionRepository struct {
	client *ent.Client
}

func NewTransactionRepository(client *ent.Client) TransactionRepository {
	return &transactionRepository{client: client}
}

func (r *transactionRepository) Start(ctx context.Context) (*ent.Tx, error) {
	return r.client.Tx(ctx)
}

func (r *transactionRepository) Commit(tx *ent.Tx) error {
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (r *transactionRepository) Rollback(tx *ent.Tx) error {
	if err := tx.Rollback(); err != nil {
		return err
	}
	return nil
}
