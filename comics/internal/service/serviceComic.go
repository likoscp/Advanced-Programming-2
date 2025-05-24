package service

import (
	"context"
	"errors"

	"github.com/likoscp/finalAddProgramming/comics/internal/repository"
	"github.com/likoscp/finalAddProgramming/comics/models"
)

type ComicsService struct {
	repo   *repository.ComicRepository
	secret string
}

func NewComicsService(repo *repository.ComicRepository, secret string) *ComicsService {
	return &ComicsService{
		repo:   repo,
		secret: secret,
	}
}

func (s *ComicsService) CreateComic(ctx context.Context, req models.Comic) (uint, error) {
	comicID, err := s.repo.CreateComic(ctx, req)
	if err != nil {
		return 0, err
	}
	return comicID, nil
}

func (s *ComicsService) UpdateComic(ctx context.Context, id uint, req models.Comic) error {
	if _, err := s.repo.GetByID(ctx, id); err != nil {
		return errors.New("comic not found")
	}
	return s.repo.UpdateComic(ctx, id, req)
}

func (s *ComicsService) GetByID(ctx context.Context, id uint) (*models.Comic, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ComicsService) GetAllComics(ctx context.Context) ([]models.Comic, error) {
	return s.repo.GetAllComics(ctx)
}

func (s *ComicsService) DeleteComic(ctx context.Context, id uint) error {
	if _, err := s.repo.GetByID(ctx, id); err != nil {
		return errors.New("comic not found")
	}
	return s.repo.DeleteComic(ctx, id)
}

func (s *ComicsService) GetComicsByUserID(ctx context.Context, userID uint) ([]models.Comic, error) {
	return s.repo.GetComicsByUserID(ctx, userID)
}
