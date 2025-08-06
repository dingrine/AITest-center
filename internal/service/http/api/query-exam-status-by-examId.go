package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type QuerExamStatusByExamIdData struct {
	ExamStatus int64 `json:"examStatus"`
	Score      int64 `json:"score"`
}

type QueryExamStatusByExamIdResponse = Response[QuerExamStatusByExamIdData]

// QueryExamStatusByExamId godoc
//
//	@Summary		query nft by token id
//	@Description	query nft by NFT's id
//	@Tags			Query NFT by token id
//	@Accept			json
//	@Produce		json
//	@Param			ExamId	path		int						true	"exam's id"
//	@Success		200		{object}	QueryExamStatusByExamIdResponse	"查询考试状态"
//	@Success		202		{object}	QueryExamStatusByExamIdResponse	"后台服务正在查询考试状态"
//	@Failure		404		{object}	QueryExamStatusByExamIdResponse	"传入的 exam id 并没有进行考试"
//	@Failure		400		{object}	QueryExamStatusByExamIdResponse	"参数错误"
//	@Failure		500		{object}	QueryExamStatusByExamIdResponse
//	@Router			/v1/api/exam/status/query/{examId} [get]
func (s *Service) QueryExamStatusByExamId(ctx *gin.Context) {
	detraderStr := strings.ToLower(ctx.Param("examId"))
	examId, err := strconv.ParseInt(detraderStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewErrorResponse[QuerExamStatusByExamIdData](err.Error(), "invalid params"))
		return
	}

	exam, err := s.RepoData.QueryExamStatusByExamId(ctx, examId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewErrorResponse[QuerExamStatusByExamIdData](err.Error(), "query exam status failed"))
		return
	}
	fmt.Println("exam", exam)
	ctx.JSON(http.StatusOK, NewSuccessResponse(&QuerExamStatusByExamIdData{}))
	// ctx.JSON(http.StatusAccepted, NewSuccessResponse(&exam))
}
