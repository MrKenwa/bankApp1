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

type UserRepo struct {
	getter *trmsqlx.CtxGetter
	db     *postgres.PostgresDB
}

func NewUserRepo(getter *trmsqlx.CtxGetter, db *postgres.PostgresDB) UserRepo {
	return UserRepo{getter: getter, db: db}
}

func (r *UserRepo) Create(ctx context.Context, u models.User) (models.UserID, error) {
	query, args, err := sq.Insert(sqlQueries.UserTable).
		Columns(sqlQueries.InsertUserColumns...).
		Values(
			u.Name,
			u.Lastname,
			u.Patronymic,
			u.Email,
			u.Password,
			u.PassportNumber,
			time.Now(),
		).
		Suffix(fmt.Sprintf("RETURNING %s", sqlQueries.IDColumnName)).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return -1, err
	}

	t := r.getter.DefaultTrOrDB(ctx, r.db.DB)

	var id models.UserID
	if err := t.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (r *UserRepo) Get(ctx context.Context, filter models.UserFilter) (models.User, error) {
	users, err := r.GetMany(ctx, filter)
	if err != nil {
		return models.User{}, err
	}

	if len(users) == 0 {
		return models.User{}, sql.ErrNoRows
	}
	return users[0], nil
}

func (r *UserRepo) GetMany(ctx context.Context, filter models.UserFilter) (models.ManyUsers, error) {
	conds := r.getConds(filter)

	query, args, err := sq.Select(sqlQueries.GetUserColumns...).
		From(sqlQueries.UserTable).
		Where(conds).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	t := r.getter.DefaultTrOrDB(ctx, r.db.DB)

	var manyUsers models.ManyUsers
	if err := t.SelectContext(ctx, &manyUsers, query, args...); err != nil {
		return nil, err
	}
	return manyUsers, nil

}

func (r *UserRepo) Delete(ctx context.Context, id models.UserID) error {
	_, _, err := sq.Select(sqlQueries.GetUserColumns...).
		From(sqlQueries.UserTable).
		Where(sq.Eq{
			sqlQueries.UserIDColumnName: id,
		}).
		ToSql()
	if err != nil {
		return err
	}

	query, args, err := sq.Update(sqlQueries.UserTable).
		Set(sqlQueries.DeletedAtColumnName, time.Now()).
		Where(sq.Eq{
			sqlQueries.UserIDColumnName: id,
		}).
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

func (r *UserRepo) getConds(filter models.UserFilter) sq.And {
	var sb sq.And

	if len(filter.IDs) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries.IDColumnName: filter.IDs,
		})
	}
	if len(filter.Names) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries.NameColumnName: filter.Names,
		})
	}
	if len(filter.Lastnames) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries.PasswordColumnName: filter.Lastnames,
		})
	}
	if len(filter.Patronymics) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries.PasswordColumnName: filter.Patronymics,
		})
	}
	if len(filter.Emails) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries.EmailColumnName: filter.Emails,
		})
	}
	sb = append(sb, sq.Eq{
		sqlQueries.DeletedAtColumnName: nil,
	})
	return sb
}
