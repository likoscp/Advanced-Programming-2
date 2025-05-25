package service

import (
	"context"
	"errors"
	"github.com/likoscp/finalAddProgramming/comics/producer"
	"log"
	"strconv"

	"github.com/likoscp/finalAddProgramming/comics/internal/repository"
	"github.com/likoscp/finalAddProgramming/comics/models"
)

type ChaptersService struct {
	repo          *repository.ChapterRepository
	natsPublisher *producer.Publisher
}

func NewChaptersService(repo *repository.ChapterRepository, natsPublisher *producer.Publisher) *ChaptersService {
	return &ChaptersService{
		repo:          repo,
		natsPublisher: natsPublisher,
	}
}

func (s *ChaptersService) CreateChapter(ctx context.Context, req models.Chapter) (uint, error) {
	chapterID, err := s.repo.CreateChapter(ctx, req)
	if err != nil {
		return 0, err
	}
	return chapterID, nil
}

func (s *ChaptersService) UpdateChapter(ctx context.Context, id uint, req models.Chapter) error {
	chapter, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return errors.New("chapter not found")
	}
	err = s.repo.UpdateChapter(ctx, id, req)
	if err != nil {
		return err
	}
	// Publish chapter updated event
	event := producer.ChapterUpdatedEvent{
		ChapterID: strconv.FormatUint(uint64(id), 10),
		ComicID:   strconv.FormatUint(uint64(chapter.ComicID), 10),
		Title:     req.Title,
		Number:    req.Number,
	}
	if err := s.natsPublisher.PublishChapterUpdated(event); err != nil {
		// Log the error but don't fail the update
		// This ensures the chapter update succeeds even if NATS publishing fails
		log.Printf("Failed to publish chapter.updated event: %v", err)
	}
	return nil
}

func (s *ChaptersService) GetByID(ctx context.Context, id uint) (*models.Chapter, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ChaptersService) GetAllChapters(ctx context.Context) ([]models.Chapter, error) {
	return s.repo.GetAllChapters(ctx)
}

func (s *ChaptersService) DeleteChapter(ctx context.Context, id uint) error {
	if _, err := s.repo.GetByID(ctx, id); err != nil {
		return errors.New("chapter not found")
	}
	return s.repo.DeleteChapter(ctx, id)
}

func (s *ChaptersService) GetChaptersByUserID(ctx context.Context, comicID uint) ([]models.Chapter, error) {
	return s.repo.GetChaptersByUserID(ctx, comicID)
}

func (s *ChaptersService) AddPageToChapter(ctx context.Context, chapterID uint, page models.Page) (uint, error) {
	chapter, err := s.repo.GetByID(ctx, chapterID)
	if err != nil {
		return 0, errors.New("chapter not found")
	}
	if chapter.ID == 0 {
		return 0, errors.New("invalid chapter ID")
	}
	page.ChapterID = chapterID
	return s.repo.AddPageToChapter(ctx, chapterID, page.ImageURL, int32(page.PageNum))
}
