package userusecase

import (
	"context"
	entityuser "todolist/internal/entity/user"
)

type userUseCase interface {
	RegisterUser(ctx context.Context, user entityuser.CreateUserReq) (*entityuser.StatusMessage, error)
	VerifyUser(ctx context.Context, user entityuser.VerifyUserReq) (*entityuser.User, error)
	Login(ctx context.Context, user entityuser.LoginReq) (*entityuser.LoginRes, error)
}

type UserUseCaseImpl struct {
	user userUseCase
}

func NewUseCaseIml(user userUseCase) *UserUseCaseImpl {
	return &UserUseCaseImpl{user: user}
}

func (u *UserUseCaseImpl) RegisterUser(ctx context.Context, user entityuser.CreateUserReq) (*entityuser.StatusMessage, error) {
	return u.user.RegisterUser(ctx, user)
}
func (u *UserUseCaseImpl) Login(ctx context.Context, user entityuser.LoginReq) (*entityuser.LoginRes, error) {
	return u.user.Login(ctx, user)
}

func (u *UserUseCaseImpl) VerifyUser(ctx context.Context, user entityuser.VerifyUserReq) (*entityuser.User, error) {
	return u.user.VerifyUser(ctx, user)
}
