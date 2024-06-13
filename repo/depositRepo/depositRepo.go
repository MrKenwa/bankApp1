package depositRepo

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

type DepositRepo struct {
	manager *txManager.TxManager
}

func NewDepositRepo(manager *txManager.TxManager) *DepositRepo {
	return &DepositRepo{manager: manager}
}

func (r *DepositRepo) Create(ctx context.Context, d models.Deposit) (models.DepositID, error) {
	query, args, err := sq.Insert(sqlQueries.DepositTable).
		Columns(sqlQueries.InsertDepositColumns...).
		Values(
			d.UserID,
			d.Type,
			d.InterestRate,
			time.Now(),
		).
		Suffix(fmt.Sprintf("RETURNING %s", sqlQueries.DepositIDColumnName)).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return -1, err
	}

	t, err := r.manager.ExtractTXOrDB(ctx)
	if err != nil {
		return -1, err
	}

	var id models.DepositID
	if err := t.QueryRow(query, args...).Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (r *DepositRepo) Get(ctx context.Context, filter models.DepositFilter) (models.Deposit, error) {
	deposits, err := r.GetMany(ctx, filter)
	if err != nil {
		return models.Deposit{}, err
	}

	if len(deposits) == 0 {
		return models.Deposit{}, sql.ErrNoRows
	}
	return deposits[0], nil
}

func (r *DepositRepo) GetMany(ctx context.Context, filter models.DepositFilter) (models.ManyDeposits, error) {
	conds := r.getConds(filter)

	query, args, err := sq.Select(sqlQueries.GetDepositColumns...).
		From(sqlQueries.DepositTable).
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

	var manyDeposits models.ManyDeposits
	if err := t.Select(&manyDeposits, query, args...); err != nil {
		return nil, err
	}
	return manyDeposits, nil
}

func (r *DepositRepo) getConds(filter models.DepositFilter) sq.And {
	var sb sq.And

	if len(filter.IDs) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries.DepositIDColumnName: filter.IDs,
		})
	}
	if len(filter.UserIDs) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries.UserIDColumnName: filter.UserIDs,
		})
	}
	if len(filter.Types) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries.DepositTypeColumnName: filter.Types,
		})
	}
	if len(filter.InterestRates) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries.InterestRateColumnName: filter.InterestRates,
		})
	}
	sb = append(sb, sq.Eq{
		sqlQueries.DeletedAtColumnName: nil,
	})
	return sb
}

func (r *DepositRepo) Delete(ctx context.Context, id models.DepositID) error {
	query, args, err := sq.Update(sqlQueries.DepositTable).
		Set(sqlQueries.DeletedAtColumnName, time.Now()).
		Where(sq.Eq{
			sqlQueries.DepositIDColumnName: id,
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
