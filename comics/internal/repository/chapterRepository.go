package repository

import (
	"context"
	"errors"

	"github.com/likoscp/finalAddProgramming/comics/models"
	"gorm.io/gorm"
)

type ChapterRepository struct {
	db *gorm.DB
}

func NewChapterRepository(db *gorm.DB) *ChapterRepository {
	return &ChapterRepository{db: db}
}

func (r *ChapterRepository) CreateChapter(ctx context.Context, chapter models.Chapter) (uint, error) {
	if chapter.ComicID == 0 {
		return 0, errors.New("chapter has invalid or empty fields")
	}

	if err := r.db.WithContext(ctx).Create(&chapter).Error; err != nil {
		return 0, err
	}
	return chapter.ID, nil
}

func (r *ChapterRepository) GetByID(ctx context.Context, id uint) (*models.Chapter, error) {
	var chapter models.Chapter
	if err := r.db.WithContext(ctx).Preload("Pages").First(&chapter, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("chapter not found")
		}
		return nil, err
	}
	return &chapter, nil
}

func (r *ChapterRepository) GetAllChapters(ctx context.Context) ([]models.Chapter, error) {
	var chapters []models.Chapter
	if err := r.db.WithContext(ctx).Preload("Pages").Find(&chapters).Error; err != nil {
		return nil, err
	}
	return chapters, nil
}

func (r *ChapterRepository) UpdateChapter(ctx context.Context, id uint, updated models.Chapter) error {
	var chapter models.Chapter
	if err := r.db.WithContext(ctx).First(&chapter, id).Error; err != nil {
		return err
	}
	updated.ID = id
	return r.db.WithContext(ctx).Save(&updated).Error
}

func (r *ChapterRepository) DeleteChapter(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&models.Chapter{}, id)
	if result.RowsAffected == 0 {
		return errors.New("chapter not found")
	}
	return result.Error
}

func (r *ChapterRepository) GetChaptersByUserID(ctx context.Context, userID uint) ([]models.Chapter, error) {
	var chapters []models.Chapter
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Preload("Pages").Find(&chapters).Error; err != nil {
		return nil, err
	}
	return chapters, nil
}

func (r *ChapterRepository) AddPageToChapter(ctx context.Context, chapterID uint, imageURL string, pageNum int32) (uint, error) {
	page := models.Page{
		ChapterID: chapterID,
		ImageURL:  imageURL,
		PageNum:   int(pageNum),
	}
	err := r.db.WithContext(ctx).Create(&page).Error
	if err != nil {
		return 0, err
	}
	return page.ID, nil
}