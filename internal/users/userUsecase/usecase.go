package userUsecase

import (
	"bankApp1/internal/models"
	"bankApp1/pkg/utils"
	"context"
	"errors"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"golang.org/x/crypto/bcrypt"
)

type UserUC struct {
	manager  *manager.Manager
	userRepo UserRepo
}

func NewUserUC(manager *manager.Manager, userRepo UserRepo) UserUC {
	return UserUC{
		manager:  manager,
		userRepo: userRepo,
	}
}

func (u *UserUC) Register(ctx context.Context, regData RegisterUser) (uid models.UserID, err error) {

	hashedPswd, err := bcrypt.GenerateFromPassword([]byte(regData.Password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}
	regData.Password = string(hashedPswd)

	hashedPsrt, err := bcrypt.GenerateFromPassword([]byte(regData.PassportNumber), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}
	regData.PassportNumber = string(hashedPsrt)

	user := regData.toUser()

	uid, err = u.userRepo.Create(ctx, user)
	if err != nil {
		return -1, err
	}

	return uid, nil
}

func (u *UserUC) Login(ctx context.Context, logData LoginUser) (uid models.UserID, err error) {
	filter := logData.toUserFilter()
	user, err := u.userRepo.Get(ctx, filter)
	if err != nil {
		return -1, errors.New("user not found")
	}

	if !utils.IsPasswordCorrect([]byte(logData.Password), []byte(user.Password)) {
		return -1, errors.New("wrong password")
	}
	uid = user.UserID

	return uid, nil
}
