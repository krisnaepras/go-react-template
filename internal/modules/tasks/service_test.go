package tasks

import (
	"context"
	"errors"
	"testing"
)

type fakeRepo struct{ lastTitle string }

func (*fakeRepo) List(context.Context, int32) ([]Task, error) { return []Task{}, nil }
func (f *fakeRepo) Create(_ context.Context, title string) (Task, error) {
	f.lastTitle = title
	return Task{Title: title}, nil
}
func (*fakeRepo) SetCompleted(context.Context, int64, bool) (Task, error) { return Task{}, nil }
func (*fakeRepo) Delete(context.Context, int64) error                     { return nil }
func TestCreateTrimsTitle(t *testing.T) {
	f := &fakeRepo{}
	got, err := NewService(f).Create(context.Background(), "  Example  ")
	if err != nil || got.Title != "Example" || f.lastTitle != "Example" {
		t.Fatalf("unexpected result: %+v, %v", got, err)
	}
}
func TestCreateRejectsInvalidTitle(t *testing.T) {
	_, err := NewService(&fakeRepo{}).Create(context.Background(), "   ")
	if !errors.Is(err, ErrInvalidTitle) {
		t.Fatalf("expected invalid title, got %v", err)
	}
}
