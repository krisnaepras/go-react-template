-- name: ListTasks :many
SELECT id, title, completed, created_at FROM tasks ORDER BY created_at DESC, id DESC LIMIT $1;

-- name: GetTask :one
SELECT id, title, completed, created_at FROM tasks WHERE id = $1;

-- name: CreateTask :one
INSERT INTO tasks (title) VALUES ($1) RETURNING id, title, completed, created_at;

-- name: SetTaskCompleted :one
UPDATE tasks SET completed = $2 WHERE id = $1 RETURNING id, title, completed, created_at;

-- name: DeleteTask :execrows
DELETE FROM tasks WHERE id = $1;
