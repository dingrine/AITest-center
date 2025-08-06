package repository

import (
	"context"
	"examCenter/internal/chaindata/models"
)

func (r *Repository) AddQARecord(ctx context.Context, examId int64, qCode, q, detraderAns, analysis string) (*models.ExamQa, error) {
	var model = models.ExamQa{
		ExamId:       examId,
		QuestionCode: qCode,
		Question:     q,
		Answer:       detraderAns,
		Analysis:     &analysis,
	}
	db := r.db.Create(&model)
	if db.Error != nil {
		return nil, db.Error
	}
	return &model, nil
}
