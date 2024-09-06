package taskservice

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
	"time"
	tasksentity "todolist/internal/entity/tasks"
	taskusecase "todolist/internal/usecase/task"
)

type Task struct {
	task   *taskusecase.TaskRepoIml
	logger *slog.Logger
}

func NewTask(task *taskusecase.TaskRepoIml, logger *slog.Logger) *Task {
	return &Task{
		task:   task,
		logger: logger,
	}
}

func (t *Task) CreateTask(ctx context.Context, req *tasksentity.CreateTaskReq) (*tasksentity.Task, error) {
	const operation = "Service.task.CreateTask"
	log := t.logger.With(
		"operation", operation)
	log.Info("Start to create task")
	defer log.Info("End to create task")
	task := tasksentity.TaskPostgres{
		UserId:   req.UserId,
		Title:    req.Title,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		Id:       uuid.NewString(),
	}
	status := tasksentity.MongoTaskDetails{
		TaskId:      task.Id,
		Condition:   req.Status.Condition,
		Important:   req.Status.Important,
		Description: req.Status.Description,
	}
	err := t.task.CreateTask(ctx, &task)
	if err != nil {
		log.Error(operation, err)
		return nil, err
	}
	err = t.task.SaveDetails(ctx, &status)
	if err != nil {
		log.Error(operation, err)
		return nil, err
	}

	return &tasksentity.Task{
		Id:       task.Id,
		UserId:   task.UserId,
		Title:    task.Title,
		CreateAt: task.CreateAt,
		Status: tasksentity.Status{
			Condition:   status.Condition,
			Important:   status.Important,
			Description: status.Description,
		},
	}, nil
}

func (t *Task) UpdateTask(ctx context.Context, req *tasksentity.UpdateTaskReq) (*tasksentity.Task, error) {
	const operation = "Service.task.UpdateTask"
	log := t.logger.With(
		"operation", operation)
	log.Info("Start to update task")
	defer log.Info("End to update task")
	if req.Title != "" {
		err := t.task.UpdateTask(ctx, req.UserId, req.Title)
		if err != nil {
			log.Error(operation, err)
			return nil, err
		}
	}
	err := t.task.UpdateDetails(ctx, &tasksentity.MongoTaskDetails{
		TaskId:      req.Id,
		Condition:   req.Status.Condition,
		Important:   req.Status.Important,
		Description: req.Status.Description,
	})
	if err != nil {
		log.Error(operation, err)
		return nil, err
	}
	task, err := t.task.GetTask(ctx, req.UserId, "id", req.Id)
	if err != nil {
		log.Error(operation, err)
		return nil, err
	}
	return &tasksentity.Task{
		Id:       task.Id,
		UserId:   task.UserId,
		Title:    task.Title,
		CreateAt: task.CreateAt,
		Status: tasksentity.Status{
			Condition:   req.Status.Condition,
			Important:   req.Status.Important,
			Description: req.Status.Description,
		},
	}, nil
}

func (t *Task) GetTask(ctx context.Context, req *tasksentity.GetTaskReq) (*tasksentity.Task, error) {
	const operation = "Service.task.GetTask"
	log := t.logger.With(
		"operation", operation)
	log.Info("Start to get task")
	defer log.Info("End to get task")

	task, err := t.task.GetTask(ctx, req.UserId, req.Field, req.Value)
	if err != nil {
		log.Error(operation, err)
		return nil, err
	}
	status, err := t.task.GetDetails(ctx, task.Id)
	if err != nil {
		log.Error(operation, err)
		return nil, err
	}
	return &tasksentity.Task{
		Id:       task.Id,
		UserId:   task.UserId,
		Title:    task.Title,
		CreateAt: task.CreateAt,
		Status: tasksentity.Status{
			Condition:   status.Condition,
			Important:   status.Important,
			Description: status.Description,
		},
	}, nil
}

func (t *Task) GetALlTask(ctx context.Context, req *tasksentity.GetAllTaskReq) ([]*tasksentity.Task, error) {
	const operation = "Service.task.GetAllTask"
	log := t.logger.With(
		"operation", operation)

	log.Info("Start to get all task")
	defer log.Info("End to get all task")
	tasks, err := t.task.GetAllTasks(ctx, req)
	if err != nil {
		log.Error(operation, err)
		return nil, err
	}

	result := []*tasksentity.Task{}
	for _, task := range tasks {
		status, err := t.task.GetDetails(ctx, task.Id)
		if err != nil {
			log.Error(operation, err)
			return nil, err
		}
		result = append(result, &tasksentity.Task{
			Id:       task.Id,
			UserId:   task.UserId,
			Title:    task.Title,
			CreateAt: task.CreateAt,
			Status: tasksentity.Status{
				Condition:   status.Condition,
				Important:   status.Important,
				Description: status.Description,
			},
		})
	}
	return result, nil
}

func (t *Task) DeleteTask(ctx context.Context, userID, taskID string) error {
	const operation = "Service.task.DeleteTask"
	log := t.logger.With(
		"operation", operation)
	log.Info("Start to delete task")
	defer log.Info("End to delete task")
	err := t.task.DeleteTask(ctx, userID, taskID)
	if err != nil {
		log.Error(operation, err)
		return err
	}
	err = t.task.DeleteDetails(ctx, taskID)
	if err != nil {
		log.Error(operation, err)
	}
	return nil
}
