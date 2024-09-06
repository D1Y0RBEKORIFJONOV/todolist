package redisrepository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
	"todolist/internal/config"
	entityuser "todolist/internal/entity/user"
)

type RedisUserRepository struct {
	redisClient *redis.Client
}

func NewRedis(cfg config.Config) *RedisUserRepository {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return &RedisUserRepository{
		redisClient: client,
	}
}

func (r *RedisUserRepository) SaveUserReq(ctx context.Context, user entityuser.UserRegisterReq, ttl time.Duration, key string) error {
	userJson, err := json.Marshal(user)
	if err != nil {
		return err
	}
	key += fmt.Sprintf(":%s", user.Email)
	err = r.redisClient.Set(ctx, key, string(userJson), ttl).Err()
	if err != nil {
		return err
	}
	return nil
}
func (r *RedisUserRepository) GetUserRegister(ctx context.Context, email, key string) (*entityuser.UserRegisterReq, error) {
	key += fmt.Sprintf(":%s", email)

	val, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var user entityuser.UserRegisterReq
	err = json.Unmarshal([]byte(val), &user)
	if err != nil {

		return nil, err
	}

	return &user, nil
}

func (r *RedisUserRepository) GetUser(ctx context.Context, email string, key string) (*entityuser.User, error) {
	key += fmt.Sprintf(":%s", email)
	val, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var user entityuser.User
	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
