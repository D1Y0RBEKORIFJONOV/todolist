package taskusecase

import (
	"context"
	tasksentity "todolist/internal/entity/tasks"
)

type taskUseCase interface {
	CreateTask(ctx context.Context, req *tasksentity.CreateTaskReq) (*tasksentity.Task, error)
	UpdateTask(ctx context.Context, req *tasksentity.UpdateTaskReq) (*tasksentity.Task, error)
	GetTask(ctx context.Context, req *tasksentity.GetTaskReq) (*tasksentity.Task, error)
	GetALlTask(ctx context.Context, req *tasksentity.GetAllTaskReq) ([]*tasksentity.Task, error)
	DeleteTask(ctx context.Context, userID, taskID string) error
}

type TaskUseCaseIml struct {
	task taskUseCase
}

func NewTaskUseCaseIml(task taskUseCase) *TaskUseCaseIml {
	return &TaskUseCaseIml{
		task: task,
	}
}

func (t *TaskUseCaseIml) CreateTask(ctx context.Context, req *tasksentity.CreateTaskReq) (*tasksentity.Task, error) {
	return t.task.CreateTask(ctx, req)
}

func (t *TaskUseCaseIml) UpdateTask(ctx context.Context, req *tasksentity.UpdateTaskReq) (*tasksentity.Task, error) {
	return t.task.UpdateTask(ctx, req)
}

func (t *TaskUseCaseIml) GetTask(ctx context.Context, req *tasksentity.GetTaskReq) (*tasksentity.Task, error) {
	return t.task.GetTask(ctx, req)
}

func (t *TaskUseCaseIml) GetALlTask(ctx context.Context, req *tasksentity.GetAllTaskReq) ([]*tasksentity.Task, error) {
	return t.task.GetALlTask(ctx, req)
}

func (t *TaskUseCaseIml) DeleteTask(ctx context.Context, userID, taskID string) error {
	return t.task.DeleteTask(ctx, userID, taskID)
}
