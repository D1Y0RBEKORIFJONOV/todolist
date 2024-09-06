package userpostgres

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"log/slog"
	entityuser "todolist/internal/entity/user"
	"todolist/internal/postgres"
)

type UserRepository struct {
	db        *postgres.PostgresDB
	tableName string
	log       *slog.Logger
}

func NewSolderRepository(db *postgres.PostgresDB, log *slog.Logger) *UserRepository {
	return &UserRepository{
		db:        db,
		tableName: "users",
		log:       log,
	}
}

func selectQuery() string {
	return `
	id,
	username,
	email,
	pass_hash
`
}

func (u *UserRepository) SaveUser(ctx context.Context, user *entityuser.User) error {
	const op = "Repository.SaveUser"
	log := u.log.With(
		slog.String("method", op))
	log.Info("Saver User start")
	defer log.Info("Saver User end")
	data := map[string]interface{}{}

	data["id"] = user.Id
	data["username"] = user.UserName
	data["email"] = user.Email
	data["pass_hash"] = user.PasswordHash

	query, args, err := u.db.Sq.Builder.Insert(u.tableName).SetMap(data).ToSql()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	_, err = u.db.Exec(ctx, query, args...)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (u *UserRepository) GetUser(ctx context.Context, field, value string) (*entityuser.User, error) {
	const op = "Repository.GetUser"
	log := u.log.With(
		slog.String("method", op))
	log.Info("Get User start")
	defer log.Info("Get User end")

	query, args, err := u.db.Sq.Builder.Select(selectQuery()).From(u.tableName).
		Where(u.db.Sq.Equal(field, value)).ToSql()
	if err != nil {
		log.Error("err", err)
		return nil, err
	}
	var user entityuser.User
	err = u.db.QueryRow(ctx, query, args...).Scan(
		&user.Id,
		&user.UserName,
		&user.Email,
		&user.PasswordHash)
	if err != nil {
		log.Error("err", err)
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) DeleteUser(ctx context.Context, id string) error {
	const op = "Repository.DeleteUser"
	log := u.log.With(slog.String("method", op))
	log.Info("Delete User start")
	defer log.Info("Delete User end")
	query, args, err := u.db.Sq.Builder.Delete(u.tableName).Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		log.Error("err", err)
		return err
	}
	_, err = u.db.Exec(ctx, query, args...)
	if err != nil {
		log.Error("err", err)
		return err
	}
	return nil
}
