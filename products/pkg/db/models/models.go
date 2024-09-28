package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)




type FinancialInstitution struct {
    gorm.Model
    Name     string    `gorm:"not null"` 
    Criteria []Criteria `gorm:"foreignKey:InstitutionID"` 
}


type Criteria struct {
    gorm.Model
    InstitutionID      uuid.UUID                   `gorm:"not null"` 
    NumberOfTransactions int                   `gorm:"not null"` 
    TotalAmount        float64                `gorm:"not null"` 
    TaxPaid            float64                `gorm:"not null"` 
    FrequentAmounts    []float64             
    MinAmount          float64                `gorm:"not null"` 
    MaxAmount          float64                `gorm:"not null"` 
    FinancialInstitution FinancialInstitution `gorm:"constraint:OnDelete:CASCADE;"`
}