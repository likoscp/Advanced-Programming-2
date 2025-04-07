package service

import (
	"context"
	"fmt"

	"github.com/likoscp/Advanced-Programming-2/forum/internal/repository"
	"github.com/likoscp/Advanced-Programming-2/forum/models"

)

type ForumService struct {
	repo *repository.ForumRepository
}

func NewForumService(r *repository.ForumRepository) *ForumService {
	return &ForumService{repo: r}
}

func (s *ForumService) CreateThread(ctx context.Context, thread models.Thread) error {
	return s.repo.CreateThread(ctx, thread)
}

func (s *ForumService) GetThread(ctx context.Context, id string) (models.Thread, error) {
	return s.repo.GetThread(ctx, id)
}

func (s *ForumService) UpdateThread(ctx context.Context, id string, content string) error {

    thread, err := s.repo.GetThread(ctx, id)
    if err != nil {
        return fmt.Errorf("thread not found: %w", err)
    }
    thread.Content = content

    err = s.repo.UpdateThread(ctx, id, &thread)
    if err != nil {
        return fmt.Errorf("failed to update thread: %w", err)
    }

    return nil

}

func (s *ForumService) DeleteThread(ctx context.Context, id string) error {
	return s.repo.DeleteThread(ctx, id)
}
