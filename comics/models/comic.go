package models

import (
	"time"
)

type Comic struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	AuthorID        uint           `json:"author_id"`
	TranslatorID    uint           `json:"translator_id"`
	ArtistID        uint           `json:"artist_id"`
	Title           string         `json:"title"`
	AltTitles       []AltTitle     `gorm:"foreignKey:ComicID" json:"alt_titles"`
	Description     string         `json:"description"`
	CoverImage      string         `json:"cover_image"`
	Genres          []Genre        `gorm:"many2many:comic_genres" json:"genres"`
	Status          string         `json:"status"`
	ComicReleaseDate time.Time     `json:"comic_date"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	Views           int            `json:"views"`
	Rating          float64        `json:"rating"`
	Chapters        []Chapter      `gorm:"foreignKey:ComicID" json:"chapters"`
}

type AltTitle struct {
	ID      uint   `gorm:"primaryKey"`
	ComicID uint
	Title   string
}

type Genre struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}

type Chapter struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ComicID   uint      `json:"comic_id"`
	Title     string    `json:"title"`
	Number    float64   `json:"number"`
	Pages     []Page    `gorm:"foreignKey:ChapterID" json:"pages"`
	CreatedAt time.Time `json:"created_at"`
	Likes     int       `json:"likes"`
	Dislikes  int       `json:"dislikes"`
}

type Page struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	ChapterID uint   `json:"chapter_id"`
	ImageURL  string `json:"image_url"`
	PageNum   int    `json:"page_num"`
}

type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	PageID    uint      `json:"page_id"`
	UserID    uint      `json:"user_id"`
	UserName  string    `json:"user_name"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	Replies   []Reply   `gorm:"foreignKey:CommentID" json:"replies"`
	Likes     int       `json:"likes"`
	Dislikes  int       `json:"dislikes"`
}

type Reply struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CommentID uint      `json:"comment_id"`
	UserID    uint      `json:"user_id"`
	UserName  string    `json:"user_name"`
	Reply     string    `json:"reply"`
	CreatedAt time.Time `json:"created_at"`
	Likes     int       `json:"likes"`
	Dislikes  int       `json:"dislikes"`
}
