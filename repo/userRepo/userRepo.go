package userRepo

import (
	"bankApp1/config"
	"bankApp1/dbConnector"
	"bankApp1/models"
	"time"
)

type UserRepo struct {
	db  dbConnector.DBOps
	cfg *config.Config
}

func NewUserRepo(db dbConnector.DBOps, cfg *config.Config) *UserRepo {
	return &UserRepo{db, cfg}
}

func (r *UserRepo) Create(u *models.User) (models.UserID, error) {
	q := `INSERT INTO users (name, lastname, email, password, passport_number, created_at)
	VALUES ($1, $2, $3, $4, $5, $6) 
	RETURNING id`

	var id models.UserID

	args := []interface{}{u.Name, u.Lastname, u.Email, u.Password, u.PassportNumber, time.Now()}
	if err := r.db.Get(&id, q, args...); err != nil {
		return -1, err
	}
	return id, nil
}

func (r *UserRepo) Get(id models.UserID) (models.User, error) {
	q := `SELECT * FROM users WHERE id = $1`
	var u models.User
	if err := r.db.Get(&u, q, id); err != nil {
		return models.User{}, err
	}
	return u, nil
}

func (r *UserRepo) GetByEmail(email string) (models.User, error) {
	q := `SELECT * FROM users WHERE email = $1`
	var u models.User
	if err := r.db.Get(&u, q, email); err != nil {
		return models.User{}, err
	}
	return u, nil
}

func (r *UserRepo) Delete(id models.UserID) error {
	q := `UPDATE users SET deleted_at = $1 WHERE id = $2`
	if _, err := r.db.Exec(q, time.Now(), id); err != nil {
		return err
	}
	return nil
}
