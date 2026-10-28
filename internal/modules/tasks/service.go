package tasks

import (
	"context"
	"strings"
	"unicode/utf8"
)

type Repository interface {
	List(ctx context.Context, limit int32) ([]Task, error)
	Create(ctx context.Context, title string) (Task, error)
	SetCompleted(ctx context.Context, id int64, completed bool) (Task, error)
	Delete(ctx context.Context, id int64) error
}
type Service struct{ repo Repository }

func NewService(repo Repository) *Service                   { return &Service{repo: repo} }
func (s *Service) List(ctx context.Context) ([]Task, error) { return s.repo.List(ctx, 100) }
func (s *Service) Create(ctx context.Context, title string) (Task, error) {
	title = strings.TrimSpace(title)
	if title == "" || utf8.RuneCountInString(title) > 200 {
		return Task{}, ErrInvalidTitle
	}
	return s.repo.Create(ctx, title)
}
func (s *Service) SetCompleted(ctx context.Context, id int64, completed bool) (Task, error) {
	if id <= 0 {
		return Task{}, ErrNotFound
	}
	return s.repo.SetCompleted(ctx, id, completed)
}
func (s *Service) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return ErrNotFound
	}
	return s.repo.Delete(ctx, id)
}
