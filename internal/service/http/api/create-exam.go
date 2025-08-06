package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateExamParams struct {
	DetraderId     *int64  `json:"detraderId"`
	DetraderName   *string `json:"detraderName"`
	RegionCertType *[]int  `json:"regionCertType"`
}
type CreateExamData struct {
	ExamId int64 `json:"id"`
}

type CreateExamResponse = Response[CreateExamData]

func (s *Service) CreateExam(ctx *gin.Context) {
	s.logger.Printf("running api")
	var params CreateExamParams
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewErrorResponse[CreateExamResponse](err.Error(), "invalid params"))
		return
	}

	if params.DetraderId == nil || params.RegionCertType == nil || params.DetraderName == nil {
		ctx.JSON(http.StatusBadRequest, NewErrorResponse[CreateExamResponse]("", "invalid params"))
		return
	}

	exam, err := s.RepoData.AddExam(ctx, *params.DetraderId, *params.DetraderName, *params.RegionCertType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewErrorResponse[CreateExamResponse](err.Error(), "create exam info error"))
		return
	}

	// 判断缓存是否有该modelId、MH-xx的考试状态

	// 如果有

	// 如果没有则将key = modelId+MATxx+Protesting加入缓存， 在生成完考试并通过后将该key值删掉。

	s.triggers.NewExamStart(exam)

	ctx.JSON(http.StatusAccepted, NewSuccessResponse(&exam.Id))
}
