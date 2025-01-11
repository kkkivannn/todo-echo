package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
	"todo_echo/internal/domain/models"
)

type TaskStorage struct {
	db *sql.DB
}

func New(storagePath string) (*TaskStorage, error) {
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return &TaskStorage{db: db}, nil
}

func (ts *TaskStorage) Stop() error {
	return ts.db.Close()
}

func (ts *TaskStorage) CreateTask(ctx context.Context, task models.CreateTask) (int, error) {
	const op = "storage.sqlite.task.create"

	req, err := ts.db.Prepare("INSERT INTO Tasks(title, body, created_at, task_status_id) VALUES (?, ?, ?, ?)")
	defer req.Close()
	if err != nil {
		log.Errorf("%s: %d", op, err)
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	res, err := req.ExecContext(ctx, task.Title, task.Body, task.CreatedAt, task.StatusID)
	if err != nil {
		log.Errorf("%s: %d", op, err)
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Errorf("%s: %d", op, err)
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return int(id), nil
}

func (ts *TaskStorage) ReadTaskByID(ctx context.Context, taskID int) (models.Task, error) {
	const op = "storage.sqlite.task.getByID"

	var task models.Task
	var status models.Status

	req, err := ts.db.Prepare("SELECT t.id, t.title, t.body, t.created_at, s.id, s.status FROM Tasks t INNER JOIN Statuses s ON t.task_status_id = s.id WHERE t.id = ?")
	if err != nil {
		log.Errorf("%s: %d", op, err)
		return models.Task{}, fmt.Errorf("%s: %w", op, err)
	}

	row := req.QueryRowContext(ctx, taskID)

	err = row.Scan(&task.ID, &task.Title, &task.Body, &task.CreatedAt, &status.ID, &status.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Task{}, fmt.Errorf("%s: %w", op, err)
		}
		return models.Task{}, fmt.Errorf("%s: %w", op, err)
	}

	task.Status = status

	return task, nil
}

func (ts *TaskStorage) ReadAllTasks(ctx context.Context) ([]models.Task, error) {
	const op = "storage.sqlite.task.getAll"

	var tasks []models.Task

	req, err := ts.db.Prepare("SELECT t.id, t.title, t.body, t.created_at, s.id, s.status FROM Tasks t INNER JOIN Statuses s ON t.task_status_id = s.id")
	if err != nil {
		log.Errorf("%s: %d", op, err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := req.QueryContext(ctx)
	if err != nil {
		log.Errorf("%s: %d", op, err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for rows.Next() {
		var task models.Task
		var status models.Status

		err = rows.Scan(&task.ID, &task.Title, &task.Body, &task.CreatedAt, &status.ID, &status.Status)
		if err != nil {
			log.Errorf("%s: %d", op, err)
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		task.Status = status

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		log.Errorf("%s: %d", op, err)
		return nil, fmt.Errorf("%s: %w", op, err)

	}

	return tasks, nil
}

func (ts *TaskStorage) DeleteTask(ctx context.Context, taskID int) error {
	const op = "storage.sqlite.task.remove"

	req, err := ts.db.Prepare("DELETE FROM Tasks WHERE id = ?")
	if err != nil {
		log.Errorf("%s: %d", op, err)
		return fmt.Errorf("%s: %w", op, err)
	}

	res, err := req.ExecContext(ctx, taskID)
	if err != nil {
		log.Errorf("%s: %d", op, err)
		return fmt.Errorf("%s: %w", op, err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Errorf("%s: %d", op, err)
		return fmt.Errorf("%s: %w", op, err)
	}

	if rows == 0 {
		log.Errorf("%s: %d", op, err)
		return fmt.Errorf("%s: no rows affected", op)
	}

	return nil
}

func (ts *TaskStorage) UpdateTask(ctx context.Context, taskID int, title string, body string, statusID int) (models.Task, error) {
	const op = "storage.sqlite.task.update"

	req, err := ts.db.Prepare("UPDATE Tasks SET title = ?, body = ?, task_status_id = ? WHERE id = ?")
	if err != nil {
		log.Errorf("%s: %d", op, err)
		return models.Task{}, fmt.Errorf("%s: %w", op, err)
	}

	_, err = req.ExecContext(ctx, title, body, statusID, taskID)
	if err != nil {
		log.Errorf("%s: %d", op, err)
		return models.Task{}, fmt.Errorf("%s: %w", op, err)
	}

	task, err := ts.ReadTaskByID(ctx, taskID)
	if err != nil {
		log.Errorf("%s: %d", op, err)
		return models.Task{}, fmt.Errorf("%s: %w", op, err)
	}
	return task, nil
}
