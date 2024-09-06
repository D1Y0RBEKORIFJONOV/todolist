package userservices

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"time"
	"todolist/internal/email"
	entityuser "todolist/internal/entity/user"
	"todolist/internal/tokens"
	userusecase "todolist/internal/usecase/user"
)

type User struct {
	user   userusecase.UserUseCaseImplRepo
	logger *slog.Logger
}

func NewUserService(user userusecase.UserUseCaseImplRepo, logger *slog.Logger) *User {
	return &User{
		user:   user,
		logger: logger,
	}
}

func (u *User) RegisterUser(ctx context.Context, user entityuser.CreateUserReq) (*entityuser.StatusMessage, error) {
	const op = "Service.RegisterUser"
	log := slog.With(
		slog.String("method", op))
	log.Info("Start Register User")
	defer log.Info("End Register User")
	if user.Password != user.ConfirmPassword {
		log.Error("err", "password and confirm password did not match")
		return nil, errors.New("password and confirm password did not match")
	}

	secretCode, err := email.SenSecretCode([]string{user.Email})
	if err != nil {
		log.Error("err", err)
		return nil, errors.New("error sending email")
	}

	err = u.user.SaveUserReq(ctx, entityuser.UserRegisterReq{
		UserName:     user.UserName,
		Email:        user.Email,
		PasswordHash: user.Password,
		SecretKey:    secretCode,
	}, time.Minute*10, "user:register")

	if err != nil {
		log.Error("err", err)
		return nil, errors.New("error sending email")
	}

	return &entityuser.StatusMessage{
		Message: "Check your email",
	}, nil
}

func (u *User) VerifyUser(ctx context.Context, req entityuser.VerifyUserReq) (*entityuser.User, error) {
	const op = "Service.VerifyUser"
	log := slog.With(
		slog.String("method", op))
	log.Info("Start Verify User")
	defer log.Info("End Verify User")

	user, err := u.user.GetUserRegister(ctx, req.Email, "user:register")
	if err != nil {
		log.Error("err", err)
		return nil, errors.New("error sending email")
	}
	if req.SecretCode != user.SecretKey {
		log.Error("err", "secretCode and userSecretKey do not match")
		return nil, errors.New("secretCode and userSecretKey do not match")
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		log.Error("Failed to hash password", err.Error())
	}

	userSave := entityuser.User{
		Id:           uuid.NewString(),
		UserName:     user.UserName,
		Email:        user.Email,
		PasswordHash: string(passHash),
	}
	err = u.user.SaveUser(ctx, &userSave)
	if err != nil {
		log.Error("err", err)
		return nil, errors.New("error sending email")
	}

	return &userSave, nil
}

func (u *User) Login(ctx context.Context, req entityuser.LoginReq) (*entityuser.LoginRes, error) {
	const op = "Service.Login"
	log := u.logger.With(
		slog.String("method", op))
	log.Info("Start login req")
	defer log.Info("End login req")

	user, err := u.user.GetUser(ctx, "email", req.Email)
	if err != nil {
		log.Error("err", err)
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		log.Error("err", err)
		return nil, errors.New("email or password incorrect")
	}
	log.Info("Login Success")
	var token entityuser.Token
	token.AccessToken, token.RefreshToken, err = tokens.GenerateTokens(user)
	return &entityuser.LoginRes{
		Token: token,
	}, nil
}
