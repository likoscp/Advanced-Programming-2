package handler

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/likoscp/Advanced-Programming-2/backend/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type S3Handler struct {
	client     *minio.Client
	Bucketname string
}

func NewS3Handler(config *config.ConfigS3) (*S3Handler, error) {
	client, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.SecretKey, ""),
		Secure: config.UseSSL,
	})
	if err != nil {
		return nil, err
	}

	return &S3Handler{
		client:     client,
		Bucketname: config.Bucket,
	}, nil
}

func (h *S3Handler) UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Max memory for form parsing
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Get form files
	files := r.MultipartForm.File["files"]
	if len(files) == 0 {
		http.Error(w, "No files provided", http.StatusBadRequest)
		return
	}

	prefix := r.FormValue("comic_id") // optional: used to organize files by comic
	if prefix != "" && !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}

	for _, fileHeader := range files {
		err := h.uploadFileToS3(r.Context(), fileHeader, prefix)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to upload %s: %v", fileHeader.Filename, err), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Upload successful"))
}

func (h *S3Handler) uploadFileToS3(ctx context.Context, fileHeader *multipart.FileHeader, prefix string) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	objectName := prefix + filepath.Base(fileHeader.Filename)

	_, err = h.client.PutObject(ctx, h.Bucketname, objectName, file, fileHeader.Size, minio.PutObjectOptions{
		ContentType: fileHeader.Header.Get("Content-Type"),
	})
	return err
}
