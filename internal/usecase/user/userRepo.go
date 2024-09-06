package userusecase

import (
	"context"
	"time"
	entityuser "todolist/internal/entity/user"
)

type (
	userUseCaseRepo interface {
		SaveUserReq(ctx context.Context, user entityuser.UserRegisterReq, ttl time.Duration, key string) error
		GetUserRegister(ctx context.Context, email, key string) (*entityuser.UserRegisterReq, error)
	}

	userUseCaseRepoPostgres interface {
		SaveUser(ctx context.Context, user *entityuser.User) error
		GetUser(ctx context.Context, field, value string) (*entityuser.User, error)
	}
)

type UserUseCaseImplRepo struct {
	userRedis    userUseCaseRepo
	userPostgres userUseCaseRepoPostgres
}

func NewUserUseCaseRepo(userRedis userUseCaseRepo, userPostgres userUseCaseRepoPostgres) *UserUseCaseImplRepo {
	return &UserUseCaseImplRepo{userRedis: userRedis, userPostgres: userPostgres}
}

func (u *UserUseCaseImplRepo) SaveUserReq(ctx context.Context, user entityuser.UserRegisterReq, ttl time.Duration, key string) error {
	return u.userRedis.SaveUserReq(ctx, user, ttl, key)
}

func (u *UserUseCaseImplRepo) GetUserRegister(ctx context.Context, email, key string) (*entityuser.UserRegisterReq, error) {
	return u.userRedis.GetUserRegister(ctx, email, key)
}

func (u *UserUseCaseImplRepo) SaveUser(ctx context.Context, user *entityuser.User) error {
	return u.userPostgres.SaveUser(ctx, user)
}

func (u *UserUseCaseImplRepo) GetUser(ctx context.Context, field, value string) (*entityuser.User, error) {
	return u.userPostgres.GetUser(ctx, field, value)
}
