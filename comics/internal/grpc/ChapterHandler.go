package grpc

import (
	"context"
	"strconv"
	"time"

	"github.com/likoscp/finalAddProgramming/comics/internal/service"
	"github.com/likoscp/finalAddProgramming/comics/models"
	chapterpb "github.com/likoscp/finalAddProgramming/finalProto/gen/go/chapters"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ChapterGRPCHandler struct {
	service *service.ChaptersService
	chapterpb.UnimplementedChaptersServiceServer
}

func NewChapterGRPCHandler(service *service.ChaptersService) *ChapterGRPCHandler {
	return &ChapterGRPCHandler{service: service}
}

func (h *ChapterGRPCHandler) CreateChapter(ctx context.Context, req *chapterpb.CreateChapterRequest) (*chapterpb.CreateChapterResponse, error) {
	chapter := models.Chapter{
		Title:     req.Title,
		CreatedAt: time.Now(),
		Number:    float64(req.Number),
		Likes:     int(req.Likes),
		Dislikes:  int(req.Dislikes),
		ComicID:   uint(req.ComicId),
	}

	id, err := h.service.CreateChapter(ctx, chapter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create chapter: %v", err)
	}

	return &chapterpb.CreateChapterResponse{Id: strconv.FormatUint(uint64(id), 10)}, nil
}

func (h *ChapterGRPCHandler) ListChapters(ctx context.Context, _ *chapterpb.Empty) (*chapterpb.ChapterList, error) {
	chapters, err := h.service.GetAllChapters(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list chapters: %v", err)
	}

	var res []*chapterpb.Chapter
	for _, ch := range chapters {
		var pages []*chapterpb.Page
		for _, p := range ch.Pages {
			pages = append(pages, &chapterpb.Page{
				Id:       strconv.FormatUint(uint64(p.ID), 10),
				ImageUrl: p.ImageURL,
				PageNum:  int32(p.PageNum),
			})
		}
		res = append(res, &chapterpb.Chapter{
			Id:        strconv.FormatUint(uint64(ch.ID), 10),
			Title:     ch.Title,
			Number:    int32(ch.Number),
			Likes:     int32(ch.Likes),
			Dislikes:  int32(ch.Dislikes),
			CreatedAt: ch.CreatedAt.Format(time.RFC3339),
			Pages:     pages,
			ComicId:   strconv.FormatUint(uint64(ch.ComicID), 10),
		})
	}

	return &chapterpb.ChapterList{Chapters: res}, nil
}

func (h *ChapterGRPCHandler) GetChapterByID(ctx context.Context, req *chapterpb.GetChapterByIDRequest) (*chapterpb.Chapter, error) {
	id, err := strconv.ParseUint(req.Id, 10, 64)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid chapter id format: %v", err)
	}

	chapter, err := h.service.GetByID(ctx, uint(id))
	if err != nil {
		if err.Error() == "chapter not found" {
			return nil, status.Errorf(codes.NotFound, "chapter not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get chapter: %v", err)
	}

	var pages []*chapterpb.Page
	for _, p := range chapter.Pages {
		pages = append(pages, &chapterpb.Page{
			Id:       strconv.FormatUint(uint64(p.ID), 10),
			ImageUrl: p.ImageURL,
			PageNum:  int32(p.PageNum),
		})
	}

	return &chapterpb.Chapter{
		Id:        strconv.FormatUint(uint64(chapter.ID), 10),
		Title:     chapter.Title,
		Number:    int32(chapter.Number),
		Likes:     int32(chapter.Likes),
		Dislikes:  int32(chapter.Dislikes),
		CreatedAt: chapter.CreatedAt.Format(time.RFC3339),
		Pages:     pages,
		ComicId:   strconv.FormatUint(uint64(chapter.ComicID), 10),
	}, nil
}

func (h *ChapterGRPCHandler) UpdateChapter(ctx context.Context, req *chapterpb.UpdateChapterRequest) (*chapterpb.Empty, error) {
	id, err := strconv.ParseUint(req.Id, 10, 64)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid chapter id format: %v", err)
	}

	updatedChapter := models.Chapter{
		Title:     req.Title,
		Number:    float64(req.Number),
		Likes:     int(req.Likes),
		Dislikes:  int(req.Dislikes),
		ComicID:   uint(req.ComicId),
	}

	err = h.service.UpdateChapter(ctx, uint(id), updatedChapter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update chapter: %v", err)
	}

	return &chapterpb.Empty{}, nil
}

func (h *ChapterGRPCHandler) DeleteChapter(ctx context.Context, req *chapterpb.DeleteChapterRequest) (*chapterpb.Empty, error) {
	id, err := strconv.ParseUint(req.Id, 10, 64)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid chapter id format: %v", err)
	}

	err = h.service.DeleteChapter(ctx, uint(id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete chapter: %v", err)
	}

	return &chapterpb.Empty{}, nil
}

func (h *ChapterGRPCHandler) AddPage(ctx context.Context, req *chapterpb.AddPageRequest) (*chapterpb.AddPageResponse, error) {
	chapterID, err := strconv.ParseUint(req.ChapterId, 10, 64)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid chapter id: %v", err)
	}

	page := models.Page{
		ImageURL: req.ImageUrl,
	}
	id, err := h.service.AddPageToChapter(ctx, uint(chapterID), page)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add page: %v", err)
	}

	return &chapterpb.AddPageResponse{Id: strconv.FormatUint(uint64(id), 10)}, nil
}
