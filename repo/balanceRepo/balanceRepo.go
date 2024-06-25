package balanceRepo

import (
	models2 "bankApp1/internal/models"
	sqlQueries2 "bankApp1/pkg/sqlQueries"
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

func (r *BalanceRepo) Create(ctx context.Context, b models2.Balance) (models2.BalanceID, error) {
	query, args, err := sq.Insert(sqlQueries2.BalanceTable).
		Columns(sqlQueries2.InsertBalanceColumns...).
		Values(
			b.CardID,
			b.DepositID,
			b.Amount,
			time.Now(),
		).
		Suffix(fmt.Sprintf("RETURNING %s", sqlQueries2.BalanceIDColumnName)).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return -1, err
	}

	t, err := r.manager.ExtractTXOrDB(ctx)
	if err != nil {
		return -1, err
	}

	var id models2.BalanceID
	if err := t.QueryRow(query, args...).Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (r *BalanceRepo) Get(ctx context.Context, filter models2.BalanceFilter) (models2.Balance, error) {
	balances, err := r.GetMany(ctx, filter)
	if err != nil {
		return models2.Balance{}, err
	}

	if len(balances) == 0 {
		return models2.Balance{}, sql.ErrNoRows
	}
	return balances[0], nil
}

func (r *BalanceRepo) GetMany(ctx context.Context, filter models2.BalanceFilter) (models2.ManyBalances, error) {
	conds := r.getConds(filter)

	query, args, err := sq.Select(sqlQueries2.GetBalanceColumns...).
		From(sqlQueries2.BalanceTable).
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

	var manyBalances models2.ManyBalances
	if err := t.Select(&manyBalances, query, args...); err != nil {
		return nil, err
	}
	return manyBalances, nil
}

func (r *BalanceRepo) getConds(filter models2.BalanceFilter) sq.And {
	var sb sq.And

	if len(filter.IDs) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries2.BalanceIDColumnName: filter.IDs,
		})
	}
	if len(filter.CardIDs) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries2.CardIDColumnName: filter.CardIDs,
		})
	}
	if len(filter.DepositIDs) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries2.DepositIDColumnName: filter.DepositIDs,
		})
	}
	sb = append(sb, sq.Eq{
		sqlQueries2.DeletedAtColumnName: nil,
	})
	return sb
}

func (r *BalanceRepo) Increase(ctx context.Context, filter models2.BalanceFilter, amount int64) error {
	conds := r.getConds(filter)

	query, args, err := sq.Update(sqlQueries2.BalanceTable).
		Set(sqlQueries2.AmountColumnName, sq.Expr(fmt.Sprintf("%s + %d", sqlQueries2.AmountColumnName, amount))).
		Where(conds).
		PlaceholderFormat(sq.Dollar).
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

func (r *BalanceRepo) Decrease(ctx context.Context, filter models2.BalanceFilter, amount int64) error {
	conds := r.getConds(filter)

	query, args, err := sq.Update(sqlQueries2.BalanceTable).
		Set(sqlQueries2.AmountColumnName, sq.Expr(fmt.Sprintf("%s - %d", sqlQueries2.AmountColumnName, amount))).
		Where(conds).
		PlaceholderFormat(sq.Dollar).
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

func (r *BalanceRepo) Delete(ctx context.Context, filter models2.BalanceFilter) error {
	conds := r.getConds(filter)
	query, args, err := sq.Update(sqlQueries2.BalanceTable).
		Set(sqlQueries2.DeletedAtColumnName, time.Now()).
		Where(conds).
		PlaceholderFormat(sq.Dollar).
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
