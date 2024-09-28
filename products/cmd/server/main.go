package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/jijengeai/jijengeai/systems/products/pb/gen"
	m "github.com/jijengeai/jijengeai/systems/products/pkg/db/models"
	repository "github.com/jijengeai/jijengeai/systems/products/pkg/db/repo"
	"github.com/sirupsen/logrus"

	h "github.com/jijengeai/jijengeai/systems/products/pkg/handler"
	srv "github.com/jijengeai/jijengeai/systems/products/pkg/service"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	dbConnectionString := os.Getenv("DATABASE_URL")
	if dbConnectionString == "" {
		logger.Fatalf("DATABASE_URL environment variable is not set")
	}

	db, err := gorm.Open(postgres.Open(dbConnectionString), &gorm.Config{})
	if err != nil {
		logger.Fatalf("Failed to connect to database: %v", err)
	}
	logger.Info("Connected to default database successfully...")

	err = db.Exec("CREATE DATABASE jejengeai").Error
	if err != nil {
		// If the database already exists, this will throw an error, which we can ignore
		logger.Info("Database 'jejengeai' may already exist: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatalf("Failed to get database instance: %v", err)
	}
	sqlDB.Close()

	db, err = gorm.Open(postgres.Open(dbConnectionString), &gorm.Config{})
	if err != nil {
		logger.Fatalf("Failed to connect to jejengeai database: %v", err)
	}
	logger.Info("Connected to jejengeai database successfully...")

	if err := migrate(db); err != nil {
		logger.Fatalf("Failed to migrate database: %v", err)
	}

	repo := repository.NewRepository(db)
	svc := srv.NewService(repo)
	handler := h.NewHandler(svc, logger)

	
	port := os.Getenv("GRPC_PORT")
    lis, err := net.Listen("tcp", ":"+port)
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }
	
	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, handler)

	
	log.Printf("Starting gRPC server on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatalf("Failed to serve: %v", err)
	}
	
}

func migrate(db *gorm.DB) error {
    if db.Migrator().HasColumn(&m.Criteria{}, "phone_number") {
        if err := db.Migrator().DropColumn(&m.Criteria{}, "phone_number"); err != nil {
            return fmt.Errorf("failed to drop phone_number column: %w", err)
        }
        log.Println("Dropped phone_number column from Business table")
    }

    if err := db.AutoMigrate(&m.FinancialInstitution{}); err != nil {
        return fmt.Errorf("failed to auto-migrate: %w", err)
    }

 
    log.Println("Migration completed successfully")
    return nil
}