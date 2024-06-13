package balanceRepo

import (
	"bankApp1/models"
	"bankApp1/sqlQueries"
	"bankApp1/txManager"
	"context"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"time"
)

type BalanceRepo struct {
	manager *txManager.TxManager
}

func NewBalanceRepo(manager *txManager.TxManager) *BalanceRepo {
	return &BalanceRepo{manager: manager}
}

func (r *BalanceRepo) Create(ctx context.Context, b models.Balance) (models.BalanceID, error) {
	query, args, err := sq.Insert(sqlQueries.BalanceTable).
		Columns(sqlQueries.InsertBalanceColumns...).
		Values(
			b.CardID,
			b.DepositID,
			b.Amount,
			time.Now(),
		).
		Suffix(fmt.Sprintf("RETURNING %s", sqlQueries.BalanceIDColumnName)).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return -1, err
	}

	t, err := r.manager.ExtractTXOrDB(ctx)
	if err != nil {
		return -1, err
	}

	var id models.BalanceID
	if err := t.QueryRow(query, args...).Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (r *BalanceRepo) Get(ctx context.Context, filter models.BalanceFilter) (models.Balance, error) {
	balances, err := r.GetMany(ctx, filter)
	if err != nil {
		return models.Balance{}, err
	}

	if len(balances) == 0 {
		return models.Balance{}, sql.ErrNoRows
	}
	return balances[0], nil
}

func (r *BalanceRepo) GetMany(ctx context.Context, filter models.BalanceFilter) (models.ManyBalances, error) {
	conds := r.getConds(filter)

	query, args, err := sq.Select(sqlQueries.GetBalanceColumns...).
		From(sqlQueries.BalanceTable).
		Where(conds).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	t, err := r.manager.ExtractTXOrDB(ctx)
	if err != nil {
		return nil, err
	}

	var manyBalances models.ManyBalances
	if err := t.Select(&manyBalances, query, args...); err != nil {
		return nil, err
	}
	return manyBalances, nil
}

func (r *BalanceRepo) getConds(filter models.BalanceFilter) sq.And {
	var sb sq.And

	if len(filter.IDs) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries.OperationIDColumnName: filter.IDs,
		})
	}
	if len(filter.CardIDs) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries.CardIDColumnName: filter.CardIDs,
		})
	}
	if len(filter.DepositIDs) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries.DepositIDColumnName: filter.DepositIDs,
		})
	}
	sb = append(sb, sq.Eq{
		sqlQueries.DeletedAtColumnName: nil,
	})
	return sb
}

func (r *BalanceRepo) Delete(ctx context.Context, id models.BalanceID) error {
	query, args, err := sq.Update(sqlQueries.BalanceTable).
		Set(sqlQueries.DeletedAtColumnName, time.Now()).
		Where(sq.Eq{
			sqlQueries.BalanceIDColumnName: id,
		}).
		ToSql()
	if err != nil {
		return err
	}

	t, err := r.manager.ExtractTXOrDB(ctx)
	if err != nil {
		return err
	}

	if _, err := t.Exec(query, args...); err != nil {
		return err
	}
	return nil
}
