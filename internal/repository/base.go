package repository

import (
	"booking-identity-management/config"
	"booking-identity-management/internal/helper"
	"context"
	"database/sql"
)

type (
	BaseRepository interface {
		BuildTransactions(ctx context.Context, transactioID string) (TransactionHandler, error)
	}

	baseRepository struct {
		cfg      *config.Config
		dbHelper helper.DBHelper
	}

	transaction struct {
		tx    *sql.Tx
		ctx   context.Context
		steps []TransactionItem
	}

	// TransactionItem wrap function for single tx executing step
	TransactionItem func(tx *TransactionExt) error

	// TransactionHandler interface for transaction handler
	TransactionHandler interface {
		SetTransactionStep(item TransactionItem) error
		ExecThenCommit() error
		RollBack()
	}
	TransactionExt struct {
		tx *sql.Tx
	}
)

// NewBaseRepository new instance for base repository
func NewBaseRepository(cfg *config.Config, dbHelper helper.DBHelper) BaseRepository {
	return &baseRepository{
		cfg:      cfg,
		dbHelper: dbHelper,
	}
}

func (r *baseRepository) BuildTransactions(ctx context.Context, transactioID string) (TransactionHandler, error) {
	tx, err := r.dbHelper.Open().Begin()
	if err != nil {
		return nil, err
	}
	return &transaction{
		ctx:   ctx,
		tx:    tx,
		steps: make([]TransactionItem, 0, 10),
	}, nil
}

func (t *transaction) SetTransactionStep(item TransactionItem) error {
	t.steps = append(t.steps, item)
	return nil
}
func (t *transaction) exec() error {
	txExt := &TransactionExt{
		tx: t.tx,
	}
	for _, transactionItem := range t.steps {
		err := transactionItem(txExt)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *transaction) commit() error {
	return t.tx.Commit()
}
func (t *transaction) rollback() error {
	return t.tx.Rollback()
}
func (t *transaction) RollBack() {
	_ = t.rollback()
}

func (t *transaction) ExecThenCommit() error {
	err := t.exec()
	if err != nil {
		_ = t.rollback()
		return err
	}

	err = t.commit()
	if err != nil {
		_ = t.rollback()
		return err
	}
	return nil
}
