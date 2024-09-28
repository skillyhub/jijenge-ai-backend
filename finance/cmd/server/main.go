package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/jijengeai/jijengeai/systems/finance/pb/gen"
	m "github.com/jijengeai/jijengeai/systems/finance/pkg/db/models"
	repository "github.com/jijengeai/jijengeai/systems/finance/pkg/db/repo"
	"github.com/sirupsen/logrus"

	h "github.com/jijengeai/jijengeai/systems/finance/pkg/handler"
	srv "github.com/jijengeai/jijengeai/systems/finance/pkg/service"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	// First, connect to the default 'postgres' database
	dbConnectionString := "host=localhost user=mac password= dbname=postgres port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dbConnectionString), &gorm.Config{})
	if err != nil {
		logger.Fatalf("Failed to connect to database: %v", err)
	}
	logger.Info("Connected to default database successfully...")

	// Create the 'jejengeai' database if it doesn't exist
	err = db.Exec("CREATE DATABASE jejengeai").Error
	if err != nil {
		// If the database already exists, this will throw an error, which we can ignore
		logger.Info("Database 'jejengeai' may already exist: ", err)
	}

	// Close the connection to the 'postgres' database
	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatalf("Failed to get database instance: %v", err)
	}
	sqlDB.Close()

	// Now connect to the 'jejengeai' database
	dbConnectionString = "host=localhost user=mac password= dbname=jejengeai port=5432 sslmode=disable"

	db, err = gorm.Open(postgres.Open(dbConnectionString), &gorm.Config{})
	if err != nil {
		logger.Fatalf("Failed to connect to jejengeai database: %v", err)
	}
	logger.Info("Connected to jejengeai database successfully...")

	// Perform migration
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
	pb.RegisterFinanceServiceServer(grpcServer, handler)

	log.Printf("Starting gRPC server on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatalf("Failed to serve: %v", err)
	}
	
}

func migrate(db *gorm.DB) error {
    // Check if phone_number column exists and remove it if it does
    if db.Migrator().HasColumn(&m.Finance{}, "phone_number") {
        if err := db.Migrator().DropColumn(&m.Finance{}, "phone_number"); err != nil {
            return fmt.Errorf("failed to drop phone_number column: %w", err)
        }
        log.Println("Dropped phone_number column from finances table")
    }
	if db.Migrator().HasColumn(&m.Transaction{}, "phone_number") {
        if err := db.Migrator().DropColumn(&m.Transaction{}, "phone_number"); err != nil {
            return fmt.Errorf("failed to drop phone_number column: %w", err)
        }
        log.Println("Dropped phone_number column from finances table")
    }

    // Auto migrate to create tables and add new columns
    if err := db.AutoMigrate(&m.Finance{}, &m.Transaction{}); err != nil {
        return fmt.Errorf("failed to auto-migrate: %w", err)
    }

    // Add business_id column if it doesn't exist
    if err := db.Exec("ALTER TABLE finances ADD COLUMN IF NOT EXISTS business_id UUID").Error; err != nil {
        return fmt.Errorf("failed to add business_id column to finances: %w", err)
    }
    if err := db.Exec("ALTER TABLE transactions ADD COLUMN IF NOT EXISTS business_id UUID").Error; err != nil {
        return fmt.Errorf("failed to add business_id column to transactions: %w", err)
    }

    // Set a default business_id for existing rows
    defaultBusinessID := uuid.Must(uuid.NewV4())
    if err := db.Exec("UPDATE finances SET business_id = ? WHERE business_id IS NULL", defaultBusinessID).Error; err != nil {
        return fmt.Errorf("failed to update existing rows with default business_id in finances: %w", err)
    }
    if err := db.Exec("UPDATE transactions SET business_id = ? WHERE business_id IS NULL", defaultBusinessID).Error; err != nil {
        return fmt.Errorf("failed to update existing rows with default business_id in transactions: %w", err)
    }

    // Set business_id as NOT NULL
    if err := db.Exec("ALTER TABLE finances ALTER COLUMN business_id SET NOT NULL").Error; err != nil {
        return fmt.Errorf("failed to set business_id as non-nullable in finances: %w", err)
    }
    if err := db.Exec("ALTER TABLE transactions ALTER COLUMN business_id SET NOT NULL").Error; err != nil {
        return fmt.Errorf("failed to set business_id as non-nullable in transactions: %w", err)
    }

    // Create indexes
    if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_finances_business_id ON finances(business_id)").Error; err != nil {
        return fmt.Errorf("failed to create index on business_id in finances: %w", err)
    }
    if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_transactions_business_id ON transactions(business_id)").Error; err != nil {
        return fmt.Errorf("failed to create index on business_id in transactions: %w", err)
    }

    log.Println("Migration completed successfully")
    return nil
}