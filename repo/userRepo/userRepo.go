package userRepo

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

type UserRepo struct {
	manager *txManager.TxManager
}

func NewUserRepo(mng *txManager.TxManager) *UserRepo {
	return &UserRepo{manager: mng}
}

func (r *UserRepo) Create(ctx context.Context, u *models.User) (models.UserID, error) {
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
		Suffix(fmt.Sprintf("RETURNING %s", sqlQueries.UserIDColumnName)).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return -1, err
	}

	t, err := r.manager.ExtractTXOrDB(ctx)
	if err != nil {
		return -1, err
	}

	var id models.UserID
	if err := t.QueryRow(query, args...).Scan(&id); err != nil {
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

	t, err := r.manager.ExtractTXOrDB(ctx)
	if err != nil {
		return nil, err
	}

	var manyUsers models.ManyUsers
	if err := t.Select(&manyUsers, query, args...); err != nil {
		return nil, err
	}
	return manyUsers, nil

}

func (r *UserRepo) Delete(ctx context.Context, id models.UserID) error {
	query, args, err := sq.Update(sqlQueries.UserTable).
		Set(sqlQueries.DeletedAtColumnName, time.Now()).
		Where(sq.Eq{
			sqlQueries.UserIDColumnName: id,
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
