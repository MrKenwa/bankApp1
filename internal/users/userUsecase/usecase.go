package userUsecase

import (
	"bankApp1/config"
	"bankApp1/internal/models"
	"bankApp1/pkg/utils"
	"context"
	"errors"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserUC struct {
	manager  *manager.Manager
	userRepo UserRepo
	config   *config.Config
}

func NewUserUC(manager *manager.Manager, userRepo UserRepo, cfg *config.Config) UserUC {
	return UserUC{
		manager:  manager,
		userRepo: userRepo,
		config:   cfg,
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

	token := jwt.New(jwt.SigningMethodES256)
	token.Claims = models.Claims{
		UserID:         user.UserID,
		Email:          user.Email,
		ExpiresAt:      time.Now().Add(time.Minute * 5),
		StandardClaims: jwt.StandardClaims{},
	}

	tokenStr, err := token.SignedString(u.config.PrivateKey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func (u *UserUC) GetUser(ctx context.Context, uid models.UserID) (models.User, error) {
	filter := models.UserFilter{
		IDs: []models.UserID{uid},
	}
	user, err := u.userRepo.Get(ctx, filter)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u *UserUC) RefreshToken(ctx context.Context, uid models.UserID) (string, error) {
	user, err := u.userRepo.Get(ctx, models.UserFilter{IDs: []models.UserID{uid}})
	if err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodES256)
	token.Claims = models.Claims{
		UserID:         user.UserID,
		Email:          user.Email,
		ExpiresAt:      time.Now().Add(time.Minute * 5),
		StandardClaims: jwt.StandardClaims{},
	}

	tokenStr, err := token.SignedString(u.config.PrivateKey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
