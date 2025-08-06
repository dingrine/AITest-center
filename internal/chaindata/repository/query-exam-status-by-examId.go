package repository

import (
	"context"
	"examCenter/internal/chaindata/models"
)

func (r *Repository) QueryExamStatusByExamId(ctx context.Context, examId int64) (*models.ExamInfo, error) {
	var model models.ExamInfo
	db := r.db.Select("exam_count", "certification_status").Where("exam_id = ?", examId).Limit(1)
	if db.Error != nil {
		return nil, db.Error
	}

	if db.RowsAffected == 0 {
		return nil, nil
	}
	return &model, nil
}
