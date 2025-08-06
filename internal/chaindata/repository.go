package chaindata

import (
	"examCenter/internal/chaindata/repository"

	"gorm.io/gorm"
)

func NewRepo(db *gorm.DB) *repository.Repository {
	return repository.New(db)
}
