package depositRepo

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

type DepositRepo struct {
	manager *txManager.TxManager
}

func NewDepositRepo(manager *txManager.TxManager) *DepositRepo {
	return &DepositRepo{manager: manager}
}

func (r *DepositRepo) Create(ctx context.Context, d models2.Deposit) (models2.DepositID, error) {
	query, args, err := sq.Insert(sqlQueries2.DepositTable).
		Columns(sqlQueries2.InsertDepositColumns...).
		Values(
			d.UserID,
			d.Type,
			d.InterestRate,
			time.Now(),
		).
		Suffix(fmt.Sprintf("RETURNING %s", sqlQueries2.DepositIDColumnName)).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return -1, err
	}

	t, err := r.manager.ExtractTXOrDB(ctx)
	if err != nil {
		return -1, err
	}

	var id models2.DepositID
	if err := t.QueryRow(query, args...).Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (r *DepositRepo) Get(ctx context.Context, filter models2.DepositFilter) (models2.Deposit, error) {
	deposits, err := r.GetMany(ctx, filter)
	if err != nil {
		return models2.Deposit{}, err
	}

	if len(deposits) == 0 {
		return models2.Deposit{}, sql.ErrNoRows
	}
	return deposits[0], nil
}

func (r *DepositRepo) GetMany(ctx context.Context, filter models2.DepositFilter) (models2.ManyDeposits, error) {
	conds := r.getConds(filter)

	query, args, err := sq.Select(sqlQueries2.GetDepositColumns...).
		From(sqlQueries2.DepositTable).
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

	var manyDeposits models2.ManyDeposits
	if err := t.Select(&manyDeposits, query, args...); err != nil {
		return nil, err
	}
	return manyDeposits, nil
}

func (r *DepositRepo) getConds(filter models2.DepositFilter) sq.And {
	var sb sq.And

	if len(filter.IDs) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries2.DepositIDColumnName: filter.IDs,
		})
	}
	if len(filter.UserIDs) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries2.UserIDColumnName: filter.UserIDs,
		})
	}
	if len(filter.Types) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries2.DepositTypeColumnName: filter.Types,
		})
	}
	if len(filter.InterestRates) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries2.InterestRateColumnName: filter.InterestRates,
		})
	}
	sb = append(sb, sq.Eq{
		sqlQueries2.DeletedAtColumnName: nil,
	})
	return sb
}

func (r *DepositRepo) Delete(ctx context.Context, id models2.DepositID) error {
	query, args, err := sq.Update(sqlQueries2.DepositTable).
		Set(sqlQueries2.DeletedAtColumnName, time.Now()).
		Where(sq.Eq{
			sqlQueries2.DepositIDColumnName: id,
		}).
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
