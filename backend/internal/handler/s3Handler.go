package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"time"

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

	ctx := context.Background()

	// Make bucket if not exists
	err = client.MakeBucket(ctx, config.Bucket, minio.MakeBucketOptions{Region: ""})
	if err != nil {
		exists, errBucketExists := client.BucketExists(ctx, config.Bucket)
		if errBucketExists == nil && exists {
			fmt.Printf("Bucket %s already exists\n", config.Bucket)
		} else {
			log.Fatalf("Failed to create bucket: %v", err)
		}
	} else {
		fmt.Printf("Successfully created bucket: %s\n", config.Bucket)
	}
	handler := &S3Handler{
		client:     client,
		Bucketname: config.Bucket,
	}
	if err := handler.SetBucketPublic(ctx); err != nil {
		log.Fatalf("Failed to set public policy: %v", err)
	}

	return handler, nil
}

func (h *S3Handler) UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["files"]
	if len(files) == 0 {
		http.Error(w, "No files provided", http.StatusBadRequest)
		return
	}

	prefix := r.FormValue("comic_id")
	if prefix != "" && !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}

	var urls []string

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, "Failed to open file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		objectName := prefix + filepath.Base(fileHeader.Filename)

		_, err = h.client.PutObject(r.Context(), h.Bucketname, objectName, file, fileHeader.Size, minio.PutObjectOptions{
			ContentType: fileHeader.Header.Get("Content-Type"),
		})
		if err != nil {
			http.Error(w, "Failed to upload file to S3", http.StatusInternalServerError)
			return
		}
		reqParams := make(url.Values)
		reqParams.Set("response-content-disposition", "inline")
		presignedURL, err := h.client.PresignedGetObject(context.Background(), h.Bucketname, objectName, time.Second*24*60*60, nil)
		if err != nil {
			slog.Error("eror to generate url", "error", err)
			http.Error(w, "Failed to generate URL", http.StatusInternalServerError)
			return
		}
		slog.Info("url here", "ur", presignedURL.String())

		urls = append(urls, presignedURL.String())
	}

	// Return URLs as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Upload successful",
		"urls":    urls,
	})
}

func (h *S3Handler) SetBucketPublic(ctx context.Context) error {
	policy := `{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Effect": "Allow",
				"Principal": {"AWS": "*"},
				"Action": ["s3:GetObject"],
				"Resource": ["arn:aws:s3:::` + h.Bucketname + `/*"]
			}
		]
	}`

	return h.client.SetBucketPolicy(ctx, h.Bucketname, policy)
}
