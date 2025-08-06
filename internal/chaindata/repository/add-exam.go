package repository

import (
	"context"
	"examCenter/internal/chaindata/models"
	"fmt"
)

func (r *Repository) AddExam(ctx context.Context, detraderId int64, detraderName string, regionCertType []int) (*models.ExamInfo, error) {
	// 目前版本regionCertType、EnvSystem只有一个，所以exam_id只有一个
	var model = &models.ExamInfo{DetraderId: detraderId,
		DetraderName:        detraderName,
		CertificationStatus: 1,
		RegionCertType:      int64(0),
		EvaluationSystem:    int64(0),
		LatestExamName:      fmt.Sprintf("%s_%s_%s", EvaSystemMap[0], RegionCertTypeMap[0], detraderName)}
	db := r.db.Create(&model)
	if db.Error != nil {
		return nil, db.Error
	}
	return model, nil
}
