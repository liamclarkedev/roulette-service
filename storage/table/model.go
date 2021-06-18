package table

import (
	"time"

	"github.com/clarke94/roulette-service/internal/pkg/table"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Table is a storage model.
type Table struct {
	ID         uuid.UUID `gorm:"primaryKey"`
	Name       string
	MaximumBet int
	MinimumBet int
	Currency   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func domainToStorage(t table.Table) Table {
	return Table{
		ID:         uuid.New(),
		Name:       t.Name,
		MaximumBet: t.MaximumBet,
		MinimumBet: t.MinimumBet,
		Currency:   t.Currency,
	}
}