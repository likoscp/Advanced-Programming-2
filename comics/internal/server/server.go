package server

import (
	"fmt"

	// "github.com/likoscp/finalAddProgramming/producer/pkg/nats"
	"github.com/likoscp/finalAddProgramming/comics/internal/config"
	grpcCustom "github.com/likoscp/finalAddProgramming/comics/internal/grpc"
	"github.com/likoscp/finalAddProgramming/comics/internal/repository"
	"github.com/likoscp/finalAddProgramming/comics/internal/service"

	comicsPb "github.com/likoscp/finalAddProgramming/finalProto/gen/go/comics"
	chaptersPb "github.com/likoscp/finalAddProgramming/finalProto/gen/go/chapters"

	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	cfg        *config.Config
	grpcServer *grpc.Server
}

func NewServer(cfg *config.Config) *Server {

	return &Server{
		cfg: cfg,
	}
}
func (s *Server) StartGRPC() error {
	lis, err := net.Listen("tcp", s.cfg.Addr)
	if err != nil {
		log.Printf("❌ Failed to listen: %v", err)
		return fmt.Errorf("failed to listen: %w", err)
	}

	s.grpcServer = grpc.NewServer()

	dsn := s.cfg.PostgresDSN
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("❌ Postgres connection failed: %v", err)
		return fmt.Errorf("postgres connection failed: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("❌ Failed to get DB from gorm: %v", err)
		return err
	}

	if err := sqlDB.Ping(); err != nil {
		log.Printf("❌ Postgres ping failed: %v", err)
		return err
	}

	log.Println("✅ Successfully connected to Postgres")

	comicRepo := repository.NewComicRepository(db)
	comicService := service.NewComicsService(comicRepo, s.cfg.Secret)

	// publisher, err := nats.NewPublisher("nats://nats:4222")
	// if err != nil {
	// 	log.Fatalf("failed to create NATS publisher: %v", err)
	// }

	comicGRPC := grpcCustom.NewComicGRPCHandler(comicService) 
	comicsPb.RegisterComicsServiceServer(s.grpcServer, comicGRPC)


	chapterGRPC := grpcCustom.NewChapterGRPCHandler(service.NewChaptersService(repository.NewChapterRepository(db)))
	chaptersPb.RegisterChaptersServiceServer(s.grpcServer, chapterGRPC)
	reflection.Register(s.grpcServer)
	log.Println("🚀 gRPC server started on port " + s.cfg.Addr)
	return s.grpcServer.Serve(lis)
}
