package grpc

import (
	"context"
	"strconv"
	"time"

	"github.com/likoscp/finalAddProgramming/comics/internal/service"
	"github.com/likoscp/finalAddProgramming/comics/models"
	comicpb "github.com/likoscp/finalAddProgramming/finalProto/gen/go/comics"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	// "github.com/nats-io/nats.go"
)

type ComicGRPCHandler struct {
	service *service.ComicsService
	comicpb.UnimplementedComicsServiceServer 
	// natsPublisher *nats.Publisher
}


func NewComicGRPCHandler(service *service.ComicsService /*, natsPublisher *nats.Publisher*/) *ComicGRPCHandler {
	return &ComicGRPCHandler{
		service: service,
		// natsPublisher: natsPublisher,
	}
}

func (h *ComicGRPCHandler) CreateComic(ctx context.Context, req *comicpb.CreateComicRequest) (*comicpb.CreateComicResponse, error) {
	translator_id, err := strconv.ParseUint(req.TranslatorId, 10, 64)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user id format: %v", err)
	}

	comic := models.Comic{
		TranslatorID: uint(translator_id),
		Status:       req.Status,
		Title:        req.Title,
		Description:  req.Description,
		CoverImage:   req.CoverImage,
		ComicReleaseDate: time.Time{}, 
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Views:        0,
		Rating:       0.0,
	}

	id, err := h.service.CreateComic(ctx, comic)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create comic: %v", err)
	}

	// if h.natsPublisher != nil {
	// 	natsData := struct {
	// 		ComicID string `json:"comicId"`
	// 		UserID  string `json:"userId"`
	// 	}{
	// 		ComicID: strconv.FormatUint(uint64(id), 10),
	// 		UserID:  req.UserId,
	// 	}
	// 	if err := h.natsPublisher.PublishComicUpdated(natsData); err != nil {
	// 		log.Printf("Failed to publish comic created event: %v", err)
	// 	}
	// } else {
	// 	log.Println("NATS publisher is not initialized")
	// }

	return &comicpb.CreateComicResponse{Id: strconv.FormatUint(uint64(id), 10)}, nil
}
func (h *ComicGRPCHandler) ListComics(ctx context.Context, _ *comicpb.Empty) (*comicpb.ComicList, error) {
	comicList, err := h.service.GetAllComics(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list comic: %v", err)
	}

	var res []*comicpb.Comic
	for _, comic := range comicList {
		var altTitles []*comicpb.AltTitle
		for _, alt := range comic.AltTitles {
			altTitles = append(altTitles, &comicpb.AltTitle{
				Id:    strconv.FormatUint(uint64(alt.ID), 10),
				Title: alt.Title,
			})
		}

		var genres []*comicpb.Genre
		for _, genre := range comic.Genres {
			genres = append(genres, &comicpb.Genre{
				Id:   strconv.FormatUint(uint64(genre.ID), 10),
				Name: genre.Name,
			})
		}

		var chapters []*comicpb.Chapter
		for _, ch := range comic.Chapters {
			chapters = append(chapters, &comicpb.Chapter{
				Id:       strconv.FormatUint(uint64(ch.ID), 10),
				Title:    ch.Title,
				Number:   ch.Number,
				Likes:    int32(ch.Likes),
				Dislikes: int32(ch.Dislikes),
			})
		}

		res = append(res, &comicpb.Comic{
			Id:               strconv.FormatUint(uint64(comic.ID), 10),
			AuthorId:         strconv.FormatUint(uint64(comic.AuthorID), 10),
			TranslatorId:     strconv.FormatUint(uint64(comic.TranslatorID), 10),
			ArtistId:         strconv.FormatUint(uint64(comic.ArtistID), 10),
			Title:            comic.Title,
			Description:      comic.Description,
			CoverImage:       comic.CoverImage,
			Status:           comic.Status,
			ComicReleaseDate: comic.ComicReleaseDate.Format(time.RFC3339),
			CreatedAt:        comic.CreatedAt.Format(time.RFC3339),
			UpdatedAt:        comic.UpdatedAt.Format(time.RFC3339),
			Views:            int32(comic.Views),
			Rating:           comic.Rating,
			AltTitles:        altTitles,
			Genres:           genres,
			Chapters:         chapters,
		})
	}

	return &comicpb.ComicList{Comics: res}, nil
}

func (h *ComicGRPCHandler) GetComicByID(ctx context.Context, req *comicpb.GetComicByIDRequest) (*comicpb.Comic, error) {
	comicID, err := strconv.ParseUint(req.Id, 10, 64)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid comic id format: %v", err)
	}
	comic, err := h.service.GetByID(ctx, uint(comicID))
	if err != nil {
		if err.Error() == "comic not found" {
			return nil, status.Errorf(codes.NotFound, "comic not found: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to get comic: %v", err)
	}
	var altTitles []*comicpb.AltTitle
	for _, alt := range comic.AltTitles {
		altTitles = append(altTitles, &comicpb.AltTitle{
			Id:    strconv.FormatUint(uint64(alt.ID), 10),
			Title: alt.Title,
		})
	}
	var genres []*comicpb.Genre
	for _, genre := range comic.Genres {
		genres = append(genres, &comicpb.Genre{
			Id:   strconv.FormatUint(uint64(genre.ID), 10),
			Name: genre.Name,
		})
	}
	var chapters []*comicpb.Chapter
	for _, ch := range comic.Chapters {
		chapters = append(chapters, &comicpb.Chapter{
			Id:       strconv.FormatUint(uint64(ch.ID), 10),
			Title:    ch.Title,
			Number:   ch.Number,
			Likes:    int32(ch.Likes),
			Dislikes: int32(ch.Dislikes),
		})
	}

	return &comicpb.Comic{
		Id:               strconv.FormatUint(uint64(comic.ID), 10),
		AuthorId:         strconv.FormatUint(uint64(comic.AuthorID), 10),
		TranslatorId:     strconv.FormatUint(uint64(comic.TranslatorID), 10),
		ArtistId:         strconv.FormatUint(uint64(comic.ArtistID), 10),
		Title:            comic.Title,
		Description:      comic.Description,
		CoverImage:       comic.CoverImage,
		Status:           comic.Status,
		ComicReleaseDate: comic.ComicReleaseDate.Format(time.RFC3339),
		CreatedAt:        comic.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        comic.UpdatedAt.Format(time.RFC3339),
		Views:            int32(comic.Views),
		Rating:           comic.Rating,
		AltTitles:        altTitles,
		Genres:           genres,
		Chapters:         chapters,
	}, nil
}


func (h *ComicGRPCHandler) UpdateComic(ctx context.Context, req *comicpb.UpdateComicRequest) (*comicpb.Empty, error) {
	comicID, err := strconv.ParseUint(req.Id, 10, 64)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid comic id format: %v", err)
	}
	translatorID, err := strconv.ParseUint(req.TranslatorId, 10, 64)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid translator id format: %v", err)
	}
	authorID, err := strconv.ParseUint(req.AuthorId, 10, 64)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid author id format: %v", err)
	}
	artistID, err := strconv.ParseUint(req.ArtistId, 10, 64)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid artist id format: %v", err)
	}
	err = h.service.UpdateComic(ctx, uint(comicID), models.Comic{
		Title:            req.Title,
		ComicReleaseDate: time.Time{},
		Status:           req.Status,
		Description:      req.Description,
		CoverImage:       req.CoverImage,
		UpdatedAt:        time.Now(),
		Views:            int(req.Views),
		Rating:           float64(req.Rating),
		TranslatorID:     uint(translatorID),
		AuthorID:         uint(authorID),
		ArtistID:         uint(artistID),
		// AltTitles:        req.AltTitles, 
		// Genres:           req.Genres,

	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update: %v", err)
	}
	return &comicpb.Empty{}, nil
}

func (h *ComicGRPCHandler) DeleteComic(ctx context.Context, req *comicpb.DeleteComicRequest) (*comicpb.Empty, error) {
	comicID, err := strconv.ParseUint(req.Id, 10, 64)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid comic id format: %v", err)
	}
	err = h.service.DeleteComic(ctx, uint(comicID))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete: %v", err)
	}
	return &comicpb.Empty{}, nil
}

