package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookingTransaction struct {
	gorm.Model
	BookingTrxID                string  `json:"booking_trx_id" gorm:"type:char(36);"`
	Name                        string  `json:"name"`
	Phone                       string  `json:"phone"`
	Email                       string  `json:"email"`
	Proof                       string  `json:"proof"`
	TotalAmount                 float64 `json:"total_amount"`
	Price                       float64 `json:"price"`
	TotalTaxAmount              float64 `json:"total_tax_amount"`
	IsPaid                      bool    `json:"is_paid"`
	StartedAt                   string  `json:"started_at"`
	WeddingPackageTransactionID uint    `json:"wedding_package_id"`
}

func (b *BookingTransaction) GenerateUUID() {
	b.BookingTrxID = uuid.New().String() // Menghasilkan UUID baru dan mengonversinya ke string
}
