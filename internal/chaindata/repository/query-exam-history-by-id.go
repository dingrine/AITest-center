package repository

import (
	"context"
	"examCenter/internal/chaindata/models"
)

func (r *Repository) QueryExamHistoryById(ctx context.Context, detraderId int64) (*[]models.ExamInfo, error) {
	var model []models.ExamInfo
	db := r.db.Where("detrader_id = ?", detraderId).Order("latest_certified_at ASC").Find(&model)
	if db.Error != nil {
		return nil, db.Error
	}

	if db.RowsAffected == 0 {
		return nil, nil
	}
	return &model, nil
}
