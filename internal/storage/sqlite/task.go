package storage

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

type TaskStorage struct {
    db *sql.DB
}

func New(connStr string) *TaskStorage {
    db, err := sql.Open("sqlite3", connStr)
    if err != nil {
        return nil
    }

    if err := db.Ping(); err != nil {
        return nil
    }

    return &TaskStorage{
        db: db,
    }
}

func (d *TaskStorage) HealthCheck() error {
    if err := d.db.Ping(); err != nil {
        return err
    }
    return nil
}
