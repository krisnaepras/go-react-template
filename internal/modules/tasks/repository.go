package tasks

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct{ db *pgxpool.Pool }

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository { return &PostgresRepository{db: db} }
func (r *PostgresRepository) List(ctx context.Context, limit int32) ([]Task, error) {
	rows, err := r.db.Query(ctx, `SELECT id, title, completed, created_at FROM tasks ORDER BY created_at DESC, id DESC LIMIT $1`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tasks := make([]Task, 0)
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Completed, &t.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, rows.Err()
}
func (r *PostgresRepository) Create(ctx context.Context, title string) (Task, error) {
	var t Task
	err := r.db.QueryRow(ctx, `INSERT INTO tasks (title) VALUES ($1) RETURNING id, title, completed, created_at`, title).Scan(&t.ID, &t.Title, &t.Completed, &t.CreatedAt)
	return t, err
}
func (r *PostgresRepository) SetCompleted(ctx context.Context, id int64, completed bool) (Task, error) {
	var t Task
	err := r.db.QueryRow(ctx, `UPDATE tasks SET completed = $2 WHERE id = $1 RETURNING id, title, completed, created_at`, id, completed).Scan(&t.ID, &t.Title, &t.Completed, &t.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return Task{}, ErrNotFound
	}
	return t, err
}
func (r *PostgresRepository) Delete(ctx context.Context, id int64) error {
	res, err := r.db.Exec(ctx, `DELETE FROM tasks WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}
