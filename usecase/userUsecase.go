package userUseCase

import (
	"bankApp1/models"
	"bankApp1/txManager"
	"bankApp1/utils"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo interface {
	Create(ctx context.Context, u *models.User) (models.UserID, error)
	Get(ctx context.Context, filter models.UserFilter) (models.User, error)
	GetMany(ctx context.Context, filter models.UserFilter) (models.ManyUsers, error)
	Delete(ctx context.Context, id models.UserID) error
}

type UserUC struct {
	manager  *txManager.TxManager
	userRepo UserRepo
}

func NewUserUC(manager *txManager.TxManager, userRepo UserRepo) *UserUC {
	return &UserUC{
		manager:  manager,
		userRepo: userRepo,
	}
}

func (u *UserUC) Register(user *models.User) (uid models.UserID, err error) {
	ctx := context.Background()
	if err := u.manager.Do(ctx, func(ctx context.Context) error {
		hashedPswd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPswd)

		hashedPsrt, err := bcrypt.GenerateFromPassword([]byte(user.PassportNumber), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.PassportNumber = string(hashedPsrt)

		uid, err = u.userRepo.Create(ctx, user)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return -1, err
	}
	return uid, nil
}

func (u *UserUC) Login(filter models.UserFilter, pwd string) (uid models.UserID, err error) {
	ctx := context.Background()
	if err := u.manager.Do(ctx, func(ctx context.Context) error {
		user, err := u.userRepo.Get(ctx, filter)
		if err != nil {
			return errors.New("user not found")
		}

		if !utils.IsPasswordCorrect([]byte(pwd), []byte(user.Password)) {
			return errors.New("wrong password")
		}
		uid = user.UserID
		return nil
	}); err != nil {
		return -1, err
	}
	return uid, nil
}
