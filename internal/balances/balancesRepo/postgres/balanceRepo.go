package postgres

import (
	"bankApp1/internal/models"
	"bankApp1/pkg/dbConnector/postgres"
	"bankApp1/pkg/sqlQueries"
	"context"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"time"
)

type BalanceRepo struct {
	getter *trmsqlx.CtxGetter
	db     *postgres.PostgresDB
}

func NewBalanceRepo(getter *trmsqlx.CtxGetter, db *postgres.PostgresDB) BalanceRepo {
	return BalanceRepo{getter: getter, db: db}
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

	t := r.getter.DefaultTrOrDB(ctx, r.db.DB)

	var id models.BalanceID
	if err := t.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
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

	t := r.getter.DefaultTrOrDB(ctx, r.db.DB)

	var manyBalances models.ManyBalances
	if err := t.SelectContext(ctx, &manyBalances, query, args...); err != nil {
		return nil, err
	}
	return manyBalances, nil
}

func (r *BalanceRepo) Increase(ctx context.Context, filter models.BalanceFilter, amount int64) error {
	conds := r.getConds(filter)
	t := r.getter.DefaultTrOrDB(ctx, r.db.DB)

	//проверка на существование баланса
	_, err := r.Get(ctx, filter)
	if err != nil {
		return err
	}

	query, args, err := sq.Update(sqlQueries.BalanceTable).
		Set(sqlQueries.AmountColumnName, sq.Expr(fmt.Sprintf("%s + %d", sqlQueries.AmountColumnName, amount))).
		Where(conds).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	if _, err := t.ExecContext(ctx, query, args...); err != nil {
		return err
	}
	return nil
}

func (r *BalanceRepo) Decrease(ctx context.Context, filter models.BalanceFilter, amount int64) error {
	conds := r.getConds(filter)

	//проверка на существование баланса
	_, err := r.Get(ctx, filter)
	if err != nil {
		return err
	}

	query, args, err := sq.Update(sqlQueries.BalanceTable).
		Set(sqlQueries.AmountColumnName, sq.Expr(fmt.Sprintf("%s - %d", sqlQueries.AmountColumnName, amount))).
		Where(conds).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	t := r.getter.DefaultTrOrDB(ctx, r.db.DB)

	if _, err := t.ExecContext(ctx, query, args...); err != nil {
		return err
	} //не возвращает ошибку, если такого баланса нет
	return nil
}

func (r *BalanceRepo) Delete(ctx context.Context, filter models.BalanceFilter) error {
	conds := r.getConds(filter)

	//проверка, существует ли баланс
	_, _, err := sq.Select(sqlQueries.GetBalanceColumns...).
		From(sqlQueries.BalanceTable).
		Where(conds).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	query, args, err := sq.Update(sqlQueries.BalanceTable).
		Set(sqlQueries.DeletedAtColumnName, time.Now()).
		Where(conds).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	t := r.getter.DefaultTrOrDB(ctx, r.db.DB)

	if _, err := t.ExecContext(ctx, query, args...); err != nil {
		return err
	}
	return nil
}

func (r *BalanceRepo) getConds(filter models.BalanceFilter) sq.And {
	var sb sq.And

	if len(filter.IDs) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries.BalanceIDColumnName: filter.IDs,
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
