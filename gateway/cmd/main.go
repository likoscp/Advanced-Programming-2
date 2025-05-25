package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	chapterpb "github.com/likoscp/finalAddProgramming/finalProto/gen/go/chapters"
	comicpb "github.com/likoscp/finalAddProgramming/finalProto/gen/go/comics"
)

func main() {
	log.Println("Connecting to gRPC server...")
	conn, err := grpc.NewClient("comics-service:8088", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	comicClient := comicpb.NewComicsServiceClient(conn)
	chapterClient := chapterpb.NewChaptersServiceClient(conn)

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		log.Println("CORS middleware triggered")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.POST("/comics", func(c *gin.Context) {
		var req struct {
			Title        string `json:"title"`
			Description  string `json:"description"`
			Status       string `json:"status"`
			TranslatorId string `json:"translator_id"`
			CoverImage   string `json:"cover_image"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := comicClient.CreateComic(context.Background(), &comicpb.CreateComicRequest{
			Title:        req.Title,
			Description:  req.Description,
			Status:       req.Status,
			TranslatorId: req.TranslatorId,
			CoverImage:   req.CoverImage,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"id": resp.Id})
	})

	r.GET("/comics", func(c *gin.Context) {
		resp, err := comicClient.ListComics(context.Background(), &comicpb.Empty{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp.Comics)
	})

	r.GET("/comics/:id", func(c *gin.Context) {
		id := c.Param("id")
		resp, err := comicClient.GetComicByID(context.Background(), &comicpb.GetComicByIDRequest{Id: id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	})

	r.PUT("/comics/:id", func(c *gin.Context) {
		id := c.Param("id")
		var req struct {
			Title        string  `json:"title"`
			Description  string  `json:"description"`
			Status       string  `json:"status"`
			CoverImage   string  `json:"cover_image"`
			TranslatorId string  `json:"translator_id"`
			AuthorId     string  `json:"author_id"`
			ArtistId     string  `json:"artist_id"`
			Views        int32   `json:"views"`
			Rating       float64 `json:"rating"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err := comicClient.UpdateComic(context.Background(), &comicpb.UpdateComicRequest{
			Id:           id,
			Title:        req.Title,
			Description:  req.Description,
			Status:       req.Status,
			CoverImage:   req.CoverImage,
			TranslatorId: req.TranslatorId,
			AuthorId:     req.AuthorId,
			ArtistId:     req.ArtistId,
			Views:        req.Views,
			Rating:       int32(req.Rating),
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "updated"})
	})

	r.DELETE("/comics/:id", func(c *gin.Context) {
		id := c.Param("id")
		_, err := comicClient.DeleteComic(context.Background(), &comicpb.DeleteComicRequest{Id: id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "deleted"})
	})

	r.POST("/chapters", func(c *gin.Context) {
		var req struct {
			Title    string `json:"title"`
			Number   int32  `json:"number"`
			Likes    int32  `json:"likes"`
			Dislikes int32  `json:"dislikes"`
			ComicId  uint64 `json:"comic_id"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := chapterClient.CreateChapter(context.Background(), &chapterpb.CreateChapterRequest{
			Title:    req.Title,
			Number:   req.Number,
			Likes:    req.Likes,
			Dislikes: req.Dislikes,
			ComicId:  uint32(req.ComicId),
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"id": resp.Id})
	})

	r.GET("/chapters", func(c *gin.Context) {
		resp, err := chapterClient.ListChapters(context.Background(), &chapterpb.Empty{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp.Chapters)
	})
	
	r.GET("/chapters/:id", func(c *gin.Context) {
		id := c.Param("id")
		resp, err := chapterClient.GetChapterByID(context.Background(), &chapterpb.GetChapterByIDRequest{Id: id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	})

	r.POST("/chapters/:id/pages", func(c *gin.Context) {
		chapterID := c.Param("id")

		var req struct {
			ImageURL string `json:"image_url"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := chapterClient.AddPage(context.Background(), &chapterpb.AddPageRequest{
			ChapterId: chapterID,
			ImageUrl:  req.ImageURL,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"id": resp.Id})
	})
	log.Println("Starting server on port 8089...")

	if err := r.Run(":8089"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
