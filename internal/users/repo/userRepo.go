package repo

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

type UserRepo struct {
	manager *txManager.TxManager
}

func NewUserRepo(mng *txManager.TxManager) *UserRepo {
	return &UserRepo{manager: mng}
}

func (r *UserRepo) Create(ctx context.Context, u *models2.User) (models2.UserID, error) {
	query, args, err := sq.Insert(sqlQueries2.UserTable).
		Columns(sqlQueries2.InsertUserColumns...).
		Values(
			u.Name,
			u.Lastname,
			u.Patronymic,
			u.Email,
			u.Password,
			u.PassportNumber,
			time.Now(),
		).
		Suffix(fmt.Sprintf("RETURNING %s", sqlQueries2.IDColumnName)).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return -1, err
	}

	t, err := r.manager.ExtractTXOrDB(ctx)
	if err != nil {
		return -1, err
	}

	var id models2.UserID
	if err := t.QueryRow(query, args...).Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (r *UserRepo) Get(ctx context.Context, filter models2.UserFilter) (models2.User, error) {
	users, err := r.GetMany(ctx, filter)
	if err != nil {
		return models2.User{}, err
	}

	if len(users) == 0 {
		return models2.User{}, sql.ErrNoRows
	}
	return users[0], nil
}

func (r *UserRepo) GetMany(ctx context.Context, filter models2.UserFilter) (models2.ManyUsers, error) {
	conds := r.getConds(filter)

	query, args, err := sq.Select(sqlQueries2.GetUserColumns...).
		From(sqlQueries2.UserTable).
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

	var manyUsers models2.ManyUsers
	if err := t.Select(&manyUsers, query, args...); err != nil {
		return nil, err
	}
	return manyUsers, nil

}

func (r *UserRepo) Delete(ctx context.Context, id models2.UserID) error {
	query, args, err := sq.Update(sqlQueries2.UserTable).
		Set(sqlQueries2.DeletedAtColumnName, time.Now()).
		Where(sq.Eq{
			sqlQueries2.UserIDColumnName: id,
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

func (r *UserRepo) getConds(filter models2.UserFilter) sq.And {
	var sb sq.And

	if len(filter.IDs) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries2.IDColumnName: filter.IDs,
		})
	}
	if len(filter.Names) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries2.NameColumnName: filter.Names,
		})
	}
	if len(filter.Lastnames) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries2.PasswordColumnName: filter.Lastnames,
		})
	}
	if len(filter.Patronymics) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries2.PasswordColumnName: filter.Patronymics,
		})
	}
	if len(filter.Emails) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries2.EmailColumnName: filter.Emails,
		})
	}
	sb = append(sb, sq.Eq{
		sqlQueries2.DeletedAtColumnName: nil,
	})
	return sb
}
