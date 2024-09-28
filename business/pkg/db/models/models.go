package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)


type BusinessType string

const (
    TypeWholesaler      BusinessType = "WHOLESALER"
    TypeRetailer        BusinessType = "RETAILER"
    TypeManufacturer    BusinessType = "MANUFACTURER"
    TypeServiceProvider BusinessType = "SERVICE_PROVIDER"
    TypeDistributor     BusinessType = "DISTRIBUTOR"
    TypeOnlineStore     BusinessType = "ONLINE_STORE"
    TypeFranchise       BusinessType = "FRANCHISE"
)

type Business struct {
    Id       uuid.UUID    `gorm:"type:uuid;"`
    Name     string       `gorm:"not null"`
    RegID    string       `gorm:"unique;not null"`
    Email    string       `gorm:"unique;not null"`
    Phone    string       `gorm:"not null"`
    Type     BusinessType `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
