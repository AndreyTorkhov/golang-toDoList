package storage

import (
    "context"
    "database/sql"
    "fmt"
    "log"

    "todo-golang/internal/config"

    "github.com/jackc/pgx/v5/pgxpool"
)

type TaskRepository interface {
    GetAll() ([]model.Task, error)
    GetByID(id int) (model.Task, error) 
    Add(task model.Task) error
    Delete(id int) error
    MarkDone(id int) error
    GetFiltered(done *bool) ([]model.Task, error)
}

type PostgresTaskRepository struct {
    db *pgxpool.Pool
}

func NewPostgresTaskRepository(db *pgxpool.Pool) *PostgresTaskRepository {
    return &PostgresTaskRepository{db: db}
}

func NewPostgresDB(dsn string) (*pgxpool.Pool, error) {
    dbpool, err := pgxpool.New(context.Background(), dsn)
    if err != nil {
        return nil, fmt.Errorf("failed to create connection pool: %w", err)
    }

    if err := dbpool.Ping(context.Background()); err != nil {
        return nil, fmt.Errorf("failed to ping database: %w", err)
    }

    log.Println("Connected to PostgreSQL")

    if err := createTasksTable(context.Background(), dbpool); err != nil {
        return nil, fmt.Errorf("failed to create tasks table: %w", err)
    }

    return dbpool, nil
}

func createTasksTable(ctx context.Context, db *pgxpool.Pool) error {
    query := `
    CREATE TABLE IF NOT EXISTS tasks (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        done BOOLEAN NOT NULL DEFAULT FALSE
    );`

    _, err := db.Exec(ctx, query)
    if err != nil {
        return fmt.Errorf("failed to create table: %w", err)
    }

    log.Println("Table 'tasks' is ensured to exist")
    return nil
}

func (r *PostgresTaskRepository) GetAll() ([]model.Task, error) {
    var tasks []model.Task

    query := `SELECT id, title, done FROM tasks`
    rows, err := r.db.Query(context.Background(), query)
    if err != nil {
        return nil, fmt.Errorf("failed to get tasks: %w", err)
    }
    defer rows.Close()

    for rows.Next() {
        var task model.Task
        if err := rows.Scan(&task.ID, &task.Title, &task.Done); err != nil {
            return nil, fmt.Errorf("failed to scan task: %w", err)
        }
        tasks = append(tasks, task)
    }

    return tasks, nil
}

func (r *PostgresTaskRepository) GetByID(id int) (model.Task, error) {
    var task model.Task
    query := `SELECT id, title, done FROM tasks WHERE id = $1`

    row := r.db.QueryRow(context.Background(), query, id)
    err := row.Scan(&task.ID, &task.Title, &task.Done)

    if err != nil {
        if err == sql.ErrNoRows {
            return task, fmt.Errorf("task not found")
        }
        return task, fmt.Errorf("failed to get task: %w", err)
    }
    
    return task, nil
}

func (r *PostgresTaskRepository) Add(task model.Task) error {
    query := `INSERT INTO tasks (title, done) VALUES ($1, $2)`
    _, err := r.db.Exec(context.Background(), query, task.Title, task.Done)
    if err != nil {
        return fmt.Errorf("failed to add task: %w", err)
    }

    log.Println("Task added successfully")
    return nil
}

func (r *PostgresTaskRepository) Delete(id int) error {
    query := `DELETE FROM tasks WHERE id = $1`
    _, err := r.db.Exec(context.Background(), query, id)
    if err != nil {
        return fmt.Errorf("failed to delete task: %w", err)
    }

    log.Println("Task deleted successfully")
    return nil
}

func (r *PostgresTaskRepository) MarkDone(id int) error {
    query := `UPDATE tasks SET done = TRUE WHERE id = $1`
    _, err := r.db.Exec(context.Background(), query, id)
    if err != nil {
        return fmt.Errorf("failed to mark task as done: %w", err)
    }

    log.Println("Task marked as done")
    return nil
}

func (r *PostgresTaskRepository) GetFiltered(done *bool) ([]model.Task, error) {
    var tasks []model.Task
    query := "SELECT id, title, done FROM tasks"
    var args []interface{}

    if done != nil {
        query += " WHERE done = $1"
        args = append(args, *done)
    }

    rows, err := r.db.Query(context.Background(), query, args...)
    if err != nil {
        return nil, fmt.Errorf("failed to query tasks: %w", err)
    }
    defer rows.Close()

    for rows.Next() {
        var task model.Task
        if err := rows.Scan(&task.ID, &task.Title, &task.Done); err != nil {
            return nil, fmt.Errorf("failed to scan task: %w", err)
        }
        tasks = append(tasks, task)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("rows iteration error: %w", err)
    }

    return tasks, nil
}