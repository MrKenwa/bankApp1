package userUsecase

import (
	"bankApp1/internal/models"
	"bankApp1/pkg/utils"
	"context"
	"errors"
	"fmt"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUC struct {
	manager       *manager.Manager
	userRepo      UserRepo
	userRedisRepo UserRedisRepo
}

func NewUserUC(manager *manager.Manager, userRepo UserRepo, userRedisRepo UserRedisRepo) UserUC {
	return UserUC{
		manager:       manager,
		userRepo:      userRepo,
		userRedisRepo: userRedisRepo,
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

func (u *UserUC) Login(ctx context.Context, logData LoginUser) (string, error) {
	filter := logData.toUserFilter()
	user, err := u.userRepo.Get(ctx, filter)
	if err != nil {
		return "", errors.New("user not found")
	}

	if !utils.IsPasswordCorrect([]byte(logData.Password), []byte(user.Password)) {
		return "", errors.New("wrong password")
	}

	sessionKey := fmt.Sprintf("%d:%s", user.UserID, uuid.NewString())

	if err := u.userRedisRepo.SetUserSession(ctx, sessionKey, models.Claims{
		UserID: user.UserID,
		Email:  user.Email,
	}); err != nil {
		return "", err
	}

	return sessionKey, nil
}

func (u *UserUC) GetUser(ctx context.Context, uid models.UserID) (models.User, error) {
	return u.userRepo.Get(ctx, models.UserFilter{
		IDs: []models.UserID{uid},
	})
}
