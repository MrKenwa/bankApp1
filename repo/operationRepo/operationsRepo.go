package operationRepo

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

type OperationRepo struct {
	manager *txManager.TxManager
}

func NewOperationRepo(manager *txManager.TxManager) *OperationRepo {
	return &OperationRepo{manager: manager}
}

func (r *OperationRepo) Create(ctx context.Context, op models.Operation) (models.OperationID, error) {
	query, args, err := sq.Insert(sqlQueries.OperationTable).
		Columns(sqlQueries.InsertOperationColumns...).
		Values(
			op.SenderBalanceID,
			op.ReceiverBalanceID,
			op.Amount,
			op.OperationType,
			time.Now(),
		).
		Suffix(fmt.Sprintf("RETURNING %s", sqlQueries.OperationIDColumnName)).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return -1, err
	}

	t, err := r.manager.ExtractTXOrDB(ctx)
	if err != nil {
		return -1, err
	}

	var id models.OperationID
	if err := t.QueryRow(query, args...).Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (r *OperationRepo) Get(ctx context.Context, filter models.OperationFilter) (models.Operation, error) {
	operations, err := r.GetMany(ctx, filter)
	if err != nil {
		return models.Operation{}, err
	}

	if len(operations) == 0 {
		return models.Operation{}, sql.ErrNoRows
	}
	return operations[0], nil
}

func (r *OperationRepo) GetMany(ctx context.Context, filter models.OperationFilter) (models.ManyOperations, error) {
	conds := r.getConds(filter)

	query, args, err := sq.Select(sqlQueries.GetOperationColumns...).
		From(sqlQueries.OperationTable).
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

	var manyOperations models.ManyOperations
	if err := t.Select(&manyOperations, query, args...); err != nil {
		return nil, err
	}
	return manyOperations, nil
}

func (r *OperationRepo) getConds(filter models.OperationFilter) sq.And {
	var sb sq.And

	if len(filter.IDs) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries.OperationIDColumnName: filter.IDs,
		})
	}
	if len(filter.SenderBalanceID) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries.SenderBalanceIDColumnName: filter.SenderBalanceID,
		})
	}
	if len(filter.ReceiverBalanceID) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries.ReceiverBalanceIDColumnName: filter.ReceiverBalanceID,
		})
	}
	if len(filter.OperationType) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries.OperationTypeColumnName: filter.OperationType,
		})
	}
	sb = append(sb, sq.Eq{
		sqlQueries.DeletedAtColumnName: nil,
	})
	return sb
}

func (r *OperationRepo) Delete(ctx context.Context, id models.OperationID) error {
	query, args, err := sq.Update(sqlQueries.OperationTable).
		Set(sqlQueries.DeletedAtColumnName, time.Now()).
		Where(sq.Eq{
			sqlQueries.OperationIDColumnName: id,
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
