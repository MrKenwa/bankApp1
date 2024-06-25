package usecase

import (
	"bankApp1/internal/models"
	"bankApp1/pkg/utils"
	"bankApp1/txManager"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

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

func (u *UserUC) Register(regData *RegisterUser) (uid models.UserID, err error) {
	ctx := context.Background()
	if err := u.manager.Do(ctx, func(ctx context.Context) error {
		hashedPswd, err := bcrypt.GenerateFromPassword([]byte(regData.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		regData.Password = string(hashedPswd)

		hashedPsrt, err := bcrypt.GenerateFromPassword([]byte(regData.PassportNumber), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		regData.PassportNumber = string(hashedPsrt)

		user := regData.toUser()

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
			return errors.New("users not found")
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
