package services

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"time"
	"todo_echo/internal/domain/models"
	storage "todo_echo/internal/storage/sqlite"
)

type Tasks struct {
	saverTask    SaverTask
	providerTask ProviderTask
	removerTask  RemoverTask
}

func New(s *storage.TaskStorage) *Tasks {
	return &Tasks{
		saverTask:    s,
		providerTask: s,
		removerTask:  s,
	}
}

type SaverTask interface {
	CreateTask(ctx context.Context, task models.CreateTask) (int, error)
}

type ProviderTask interface {
	ReadTaskByID(ctx context.Context, taskID int) (models.Task, error)
	ReadAllTasks(ctx context.Context) ([]models.Task, error)
	UpdateTask(ctx context.Context, taskID int, title string, body string, statusID int) (models.Task, error)
}

type RemoverTask interface {
	DeleteTask(ctx context.Context, taskID int) error
}

func (t *Tasks) AddTask(ctx context.Context, title, body string) (int, error) {
	const op = "service.tasks.add"

	task := models.CreateTask{
		Title:     title,
		Body:      body,
		CreatedAt: time.Now().UTC(),
		StatusID:  1,
	}

	id, err := t.saverTask.CreateTask(ctx, task)
	if err != nil {
		log.Errorf("%s:%d", op, err)
		return 0, err
	}

	return id, err
}

func (t *Tasks) GetTask(ctx context.Context, taskID int) (models.Task, error) {
	const op = "service.task.getByID"

	task, err := t.providerTask.ReadTaskByID(ctx, taskID)
	if err != nil {
		log.Errorf("%s:%d", op, err)
		return models.Task{}, err
	}

	return task, nil
}

func (t *Tasks) GetTasks(ctx context.Context) ([]models.Task, error) {
	const op = "service.task.getAll"

	tasks, err := t.providerTask.ReadAllTasks(ctx)
	if err != nil {
		log.Errorf("%s: %d", op, err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return tasks, nil
}

func (t *Tasks) RemoveTask(ctx context.Context, taskID int) error {
	const op = "service.task.remove"

	if err := t.removerTask.DeleteTask(ctx, taskID); err != nil {
		log.Errorf("%s: %d", op, err)
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (t *Tasks) EditTask(ctx context.Context, taskID int, title string, body string, statusID int) (models.Task, error) {
	const op = "service.task.edit"

	task, err := t.providerTask.UpdateTask(ctx, taskID, title, body, statusID)
	if err != nil {
		log.Errorf("%s: %d", op, err)
		return models.Task{}, fmt.Errorf("%s: %w", op, err)
	}
	
	return task, nil
}
