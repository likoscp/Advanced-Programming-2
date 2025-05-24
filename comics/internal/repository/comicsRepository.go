package repository

import (
	"context"
	"errors"

	"github.com/likoscp/finalAddProgramming/comics/models"
	"gorm.io/gorm"
)

type ComicRepository struct {
	db *gorm.DB
}

func NewComicRepository(db *gorm.DB) *ComicRepository {
	return &ComicRepository{db: db}
}

func (r *ComicRepository) CreateComic(ctx context.Context, comic models.Comic) (uint, error) {
	if  comic.Title == ""  {
		return 0, errors.New("comic has invalid or empty fields")
	}

	if err := r.db.WithContext(ctx).Create(&comic).Error; err != nil {
		return 0, err
	}

	return comic.ID, nil
}

func (r *ComicRepository) GetByID(ctx context.Context, id uint) (*models.Comic, error) {
	var comic models.Comic
	err := r.db.WithContext(ctx).
		Preload("Chapters").
		First(&comic, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("comic not found")
		}
		return nil, err
	}
	return &comic, nil
}


func (r *ComicRepository) GetAllComics(ctx context.Context) ([]models.Comic, error) {
	var comics []models.Comic

	err := r.db.WithContext(ctx).
		Preload("Chapters").
		Find(&comics).Error
	if err != nil {
		return nil, err
	}

	return comics, nil
}


func (r *ComicRepository) UpdateComic(ctx context.Context, id uint, updated models.Comic) error {
	var comic models.Comic
	if err := r.db.WithContext(ctx).First(&comic, id).Error; err != nil {
		return err
	}
	updated.ID = id
	return r.db.WithContext(ctx).Save(&updated).Error
}

func (r *ComicRepository) DeleteComic(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&models.Comic{}, id)
	if result.RowsAffected == 0 {
		return errors.New("comic not found")
	}
	return result.Error
}

func (r *ComicRepository) GetComicsByUserID(ctx context.Context, userID uint) ([]models.Comic, error) {
	var comics []models.Comic
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&comics).Error; err != nil {
		return nil, err
	}
	return comics, nil
}
