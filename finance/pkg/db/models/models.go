package models

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type FinanceType string

const (
	FinanceTypeMomo FinanceType = "momo"
	FinanceTypeBank FinanceType = "bank"
)


type Finance struct {
	Id           uuid.UUID `gorm:"uniqueIndex;not null;type:uuid"`
	BusinessID    uuid.UUID `gorm:"index;not null"` 
	Type          FinanceType `gorm:"not null"`
	Amount        float64     `gorm:"not null;default:0"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}


type Transaction struct {
	Id            uuid.UUID `gorm:"uniqueIndex;not null;type:uuid"`
	BusinessID    uuid.UUID `gorm:"index;not null"` 
	FinanceType   FinanceType `gorm:"not null"`
	Amount        float64     `gorm:"not null"`
	Description   string
	Date          time.Time   `gorm:"index;not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

var (
	ErrInvalidFinanceType = errors.New("invalid finance type")
)

func (f FinanceType) IsValid() error {
	switch f {
	case FinanceTypeMomo, FinanceTypeBank:
		return nil
	default:
		return ErrInvalidFinanceType
	}
}