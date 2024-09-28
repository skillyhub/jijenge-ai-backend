package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/jijengeai/jijengeai/systems/business/pb/gen"
	m "github.com/jijengeai/jijengeai/systems/business/pkg/db/models"
	repository "github.com/jijengeai/jijengeai/systems/business/pkg/db/repo"
	"github.com/sirupsen/logrus"

	h "github.com/jijengeai/jijengeai/systems/business/pkg/handler"
	srv "github.com/jijengeai/jijengeai/systems/business/pkg/service"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	dbConnectionString := "host=localhost user=mac password= dbname=postgres port=5432 sslmode=disable"

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

	dbConnectionString = "host=localhost user=mac password= dbname=jejengeai port=5432 sslmode=disable"

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
	pb.RegisterBusinessServiceServer(grpcServer, handler)

	log.Printf("Starting gRPC server on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatalf("Failed to serve: %v", err)
	}
}

func migrate(db *gorm.DB) error {
    // Check if phone_number column exists and remove it if it does
    if db.Migrator().HasColumn(&m.Business{}, "phone_number") {
        if err := db.Migrator().DropColumn(&m.Business{}, "phone_number"); err != nil {
            return fmt.Errorf("failed to drop phone_number column: %w", err)
        }
        log.Println("Dropped phone_number column from Business table")
    }

    // Auto migrate to create tables and add new columns
    if err := db.AutoMigrate(&m.Business{}); err != nil {
        return fmt.Errorf("failed to auto-migrate: %w", err)
    }

 
    log.Println("Migration completed successfully")
    return nil
}