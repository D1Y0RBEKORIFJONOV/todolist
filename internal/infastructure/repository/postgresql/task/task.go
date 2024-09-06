package taskpostgres

import (
	"context"
	"log/slog"
	tasksentity "todolist/internal/entity/tasks"
	"todolist/internal/postgres"
)

type TaskRepository struct {
	db        *postgres.PostgresDB
	tableName string
	log       *slog.Logger
}

func NewTaskRepository(db *postgres.PostgresDB, log *slog.Logger) *TaskRepository {
	return &TaskRepository{
		db:        db,
		tableName: "tasks",
		log:       log,
	}
}

func selectQuery() string {
	return `
	id,
	user_id,
	title,
	created_at
`
}

func (t *TaskRepository) CreateTask(ctx context.Context, task *tasksentity.TaskPostgres) error {
	const op = "RepositorySQL_CreateTask"
	log := t.log.With(slog.String("method", op))
	log.Info("Start Create task ")
	defer log.Info("End Create task ")
	data := map[string]interface{}{
		"id":         task.Id,
		"user_id":    task.UserId,
		"title":      task.Title,
		"created_at": task.CreateAt,
	}
	query, args, err := t.db.Sq.Builder.Insert(t.tableName).
		SetMap(data).ToSql()
	if err != nil {
		log.Error(op, err)
	}
	_, err = t.db.Exec(ctx, query, args...)
	if err != nil {
		log.Error(op, err)
		return err
	}
	return nil
}

func (t *TaskRepository) UpdateTask(ctx context.Context, userId, title string) error {
	const op = "RepositorySQL_UpdateTask"
	log := t.log.With(slog.String("method", op))
	log.Info("Start Update task ")
	defer log.Info("End Update task ")
	data := map[string]interface{}{
		"title": title,
	}

	query, args, err := t.db.Sq.Builder.Update(t.tableName).SetMap(data).
		Where(t.db.Sq.Equal("user_id", userId)).ToSql()
	if err != nil {
		log.Error(op, err)
		return err
	}
	_, err = t.db.Exec(ctx, query, args...)
	if err != nil {
		log.Error(op, err)
		return err
	}
	return nil
}

func (t *TaskRepository) DeleteTask(ctx context.Context, userId, taskID string) error {
	const op = "RepositorySQL_DeleteTask"
	log := t.log.With(slog.String("method", op))
	log.Info("Start Delete task ")
	defer log.Info("End Delete task ")

	query, args, err := t.db.Sq.Builder.Delete(t.tableName).Where(t.db.Sq.And(
		t.db.Sq.Equal("id", taskID), t.db.Sq.Equal("user_id", userId))).
		ToSql()
	if err != nil {
		log.Error(op, err)
		return err
	}
	_, err = t.db.Exec(ctx, query, args...)
	if err != nil {
		log.Error(op, err)
		return err
	}
	return nil
}

func (t *TaskRepository) GetAllTasks(ctx context.Context, req *tasksentity.GetAllTaskReq) ([]*tasksentity.TaskPostgres, error) {
	const op = "RepositorySQL_GetAllTasks"
	log := t.log.With(slog.String("method", op))
	log.Info("Start GetAllTasks")
	defer log.Info("End GetAllTasks")

	toSql := t.db.Sq.Builder.Select(selectQuery()).From(t.tableName)
	if req.Field != "" && req.Value != "" {
		toSql = toSql.Where(t.db.Sq.And(
			t.db.Sq.Equal("user_id", req.UserId),
			t.db.Sq.Equal(req.Field, req.Value)))
	} else {
		toSql = toSql.Where(t.db.Sq.And(
			t.db.Sq.Equal("user_id", req.UserId)))
	}
	if req.Offset != 0 {
		toSql = toSql.Offset(uint64(req.Offset))
	}
	if req.Limit != 0 {
		toSql = toSql.Limit(uint64(req.Limit))
	}

	query, args, err := toSql.ToSql()
	if err != nil {
		log.Error(op, err)
		return nil, err
	}
	rows, err := t.db.Query(ctx, query, args...)
	if err != nil {
		log.Error(op, err)
		return nil, err
	}
	defer rows.Close()
	tasks := make([]*tasksentity.TaskPostgres, 0)
	for rows.Next() {
		var task tasksentity.TaskPostgres
		err = rows.Scan(&task.Id, &task.UserId, &task.Title, &task.CreateAt)
		if err != nil {
			log.Error(op, err)
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	if err := rows.Err(); err != nil {
		log.Error(op, err)
		return nil, err
	}
	return tasks, nil
}

func (t *TaskRepository) GetTask(ctx context.Context, userId, field, value string) (*tasksentity.TaskPostgres, error) {
	const op = "RepositorySQL_GetTask"
	log := t.log.With(slog.String("method", op))
	log.Info("Start Get task ")
	defer log.Info("End Get task ")

	query, args, err := t.db.Sq.Builder.Select(selectQuery()).From(t.tableName).
		Where(t.db.Sq.And(
			t.db.Sq.Equal("user_id", userId), t.db.Sq.Equal(field, value))).ToSql()

	if err != nil {
		log.Error(op, err)
		return nil, err
	}
	var task tasksentity.TaskPostgres
	err = t.db.QueryRow(ctx, query, args...).Scan(&task.Id, &task.UserId, &task.Title,
		&task.CreateAt)
	if err != nil {
		log.Error(op, err)
		return nil, err
	}
	return &task, nil
}
