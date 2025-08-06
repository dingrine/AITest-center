package repository

import (
	"context"
	"examCenter/internal/chaindata/models"
)

type ExamDataRepository interface {
	AddExam(context.Context, int64, string, []int) (*models.ExamInfo, error)
	AddQARecord(context.Context, int64, string, string, string, string) (*models.ExamQa, error)
	QueryExamHistoryById(context.Context, int64) (*[]models.ExamInfo, error)
	QueryExamStatusByExamId(context.Context, int64) (*models.ExamInfo, error)
}
