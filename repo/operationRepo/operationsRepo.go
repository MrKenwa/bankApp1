package operationRepo

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

type OperationRepo struct {
	manager *txManager.TxManager
}

func NewOperationRepo(manager *txManager.TxManager) *OperationRepo {
	return &OperationRepo{manager: manager}
}

// почему бы не принимать структуру operation через ссылку?
func (r *OperationRepo) Create(ctx context.Context, op models2.Operation) (models2.OperationID, error) {
	query, args, err := sq.Insert(sqlQueries2.OperationTable).
		Columns(sqlQueries2.InsertOperationColumns...).
		Values(
			op.SenderBalanceID,
			op.ReceiverBalanceID,
			op.Amount,
			op.OperationType,
			time.Now(),
		).
		Suffix(fmt.Sprintf("RETURNING %s", sqlQueries2.OperationIDColumnName)).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return -1, err
	}

	t, err := r.manager.ExtractTXOrDB(ctx)
	if err != nil {
		return -1, err
	}

	var id models2.OperationID
	if err := t.QueryRow(query, args...).Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (r *OperationRepo) Get(ctx context.Context, filter models2.OperationFilter) (models2.Operation, error) {
	operations, err := r.GetMany(ctx, filter)
	if err != nil {
		return models2.Operation{}, err
	}

	if len(operations) == 0 {
		return models2.Operation{}, sql.ErrNoRows
	}
	return operations[0], nil
}

func (r *OperationRepo) GetMany(ctx context.Context, filter models2.OperationFilter) (models2.ManyOperations, error) {
	conds := r.getConds(filter)

	query, args, err := sq.Select(sqlQueries2.GetOperationColumns...).
		From(sqlQueries2.OperationTable).
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

	var manyOperations models2.ManyOperations
	if err := t.Select(&manyOperations, query, args...); err != nil {
		return nil, err
	}
	return manyOperations, nil
}

func (r *OperationRepo) getConds(filter models2.OperationFilter) sq.And {
	var sb sq.And

	if len(filter.IDs) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries2.OperationIDColumnName: filter.IDs,
		})
	}
	if len(filter.SenderBalanceID) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries2.SenderBalanceIDColumnName: filter.SenderBalanceID,
		})
	}
	if len(filter.ReceiverBalanceID) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries2.ReceiverBalanceIDColumnName: filter.ReceiverBalanceID,
		})
	}
	if len(filter.OperationType) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries2.OperationTypeColumnName: filter.OperationType,
		})
	}
	sb = append(sb, sq.Eq{
		sqlQueries2.DeletedAtColumnName: nil,
	})
	return sb
}

func (r *OperationRepo) Delete(ctx context.Context, id models2.OperationID) error {
	query, args, err := sq.Update(sqlQueries2.OperationTable).
		Set(sqlQueries2.DeletedAtColumnName, time.Now()).
		Where(sq.Eq{
			sqlQueries2.OperationIDColumnName: id,
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
