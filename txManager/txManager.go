package txManager

import (
	"bankApp1/pkg/dbConnector"
	"context"
	"github.com/jmoiron/sqlx"
	"log"
)

type txKey struct{}

type TxManager struct {
	db dbConnector.DBOps
}

func NewTxManager(db dbConnector.DBOps) *TxManager {
	return &TxManager{db: db}
}

// Deprecated!
func (m *TxManager) Do(ctx context.Context, fn func(ctx context.Context) error) error {
	t, err := m.db.Beginx()
	if err != nil {
		return err
	}

	err = fn(m.injectTx(ctx, t))
	if err != nil {
		if errRollBack := t.Rollback(); errRollBack != nil {
			log.Printf("Rollback failed: %v", errRollBack)
		}
		return err
	}

	if err := t.Commit(); err != nil {
		return err
	}
	return nil
}

func (m *TxManager) injectTx(ctx context.Context, tx *sqlx.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

func (m *TxManager) ExtractTXOrDB(ctx context.Context) (dbConnector.DBOps, error) {
	tx, ok := ctx.Value(txKey{}).(dbConnector.DBOps)
	if !ok {
		return m.db, nil
	}
	return tx, nil
}
