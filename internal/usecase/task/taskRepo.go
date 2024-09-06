package taskusecase

import (
	"context"
	tasksentity "todolist/internal/entity/tasks"
)

type (
	task interface {
		CreateTask(ctx context.Context, task *tasksentity.TaskPostgres) error
		UpdateTask(ctx context.Context, userId, title string) error
		DeleteTask(ctx context.Context, userId, taskID string) error
		GetTask(ctx context.Context, userId, field, value string) (*tasksentity.TaskPostgres, error)
		GetAllTasks(ctx context.Context, req *tasksentity.GetAllTaskReq) ([]*tasksentity.TaskPostgres, error)
	}
	details interface {
		SaveDetails(ctx context.Context, req *tasksentity.MongoTaskDetails) error
		GetDetails(ctx context.Context, taskID string) (*tasksentity.MongoTaskDetails, error)
		UpdateDetails(ctx context.Context, req *tasksentity.MongoTaskDetails) error
		DeleteDetails(ctx context.Context, taskID string) error
	}
)

type TaskRepoIml struct {
	task    task
	details details
}

func NewTaskRepoIml(task task, details details) *TaskRepoIml {
	return &TaskRepoIml{
		task:    task,
		details: details,
	}
}
func (t *TaskRepoIml) CreateTask(ctx context.Context, task *tasksentity.TaskPostgres) error {
	return t.task.CreateTask(ctx, task)
}
func (t *TaskRepoIml) UpdateTask(ctx context.Context, userId, title string) error {
	return t.task.UpdateTask(ctx, userId, title)
}
func (t *TaskRepoIml) DeleteTask(ctx context.Context, userId, taskID string) error {
	return t.task.DeleteTask(ctx, userId, taskID)
}

func (t *TaskRepoIml) GetTask(ctx context.Context, userId, field, value string) (*tasksentity.TaskPostgres, error) {
	return t.task.GetTask(ctx, userId, field, value)
}

func (t *TaskRepoIml) GetAllTasks(ctx context.Context, req *tasksentity.GetAllTaskReq) ([]*tasksentity.TaskPostgres, error) {
	return t.task.GetAllTasks(ctx, req)
}

func (t *TaskRepoIml) SaveDetails(ctx context.Context, req *tasksentity.MongoTaskDetails) error {
	return t.details.SaveDetails(ctx, req)
}

func (t *TaskRepoIml) GetDetails(ctx context.Context, taskID string) (*tasksentity.MongoTaskDetails, error) {
	return t.details.GetDetails(ctx, taskID)
}

func (t *TaskRepoIml) UpdateDetails(ctx context.Context, req *tasksentity.MongoTaskDetails) error {
	return t.details.UpdateDetails(ctx, req)
}

func (t *TaskRepoIml) DeleteDetails(ctx context.Context, taskID string) error {
	return t.details.DeleteDetails(ctx, taskID)
}
