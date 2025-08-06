package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type QueryExamListByExamIdData struct {
	ExamId              int64  `json:"examId"`
	ExamName            string `json:"examName"`
	RegionCertType      int64  `json:"regionCertType"`
	Score               int64  `json:"score"`
	CertificationStatus int64  `json:"certificationStatus"`
	CertifiedAt         string `json:"CertifiedAt"`
}

type QueryExamListByExamIdResponse = Response[QueryExamListByExamIdData]

// QueryExamListByExamId godoc
//
//	@Summary		query nft by token id
//	@Description	query nft by NFT's id
//	@Tags			Query NFT by token id
//	@Accept			json
//	@Produce		json
//	@Param			tokenId	path		string					true	"nft's id"
//	@Success		200		{object}	QueryExamListByExamIdResponse	"新导出NFT成功。NFT的导出流程会在后台自动进行"
//	@Success		202		{object}	QueryExamListByExamIdResponse	"后台服务正在创建NFT"
//	@Failure		404		{object}	QueryExamListByExamIdResponse	"传入的 token id 并没有被请求导出过"
//	@Failure		400		{object}	QueryExamListByExamIdResponse	"参数错误"
//	@Failure		500		{object}	QueryExamListByExamIdResponse
//	@Router			/v1/api/exam/history/list/query/{tokenId} [get]
func (s *Service) QueryExamListByExamId(ctx *gin.Context) {
	detraderStr := strings.ToLower(ctx.Param("detraderId"))
	detraderId, err := strconv.ParseInt(detraderStr, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewErrorResponse[QueryExamListByExamIdData](err.Error(), "invalid params"))
		return
	}
	s.logger.Printf("after return")

	exam, err := s.RepoData.QueryExamHistoryById(ctx, detraderId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewErrorResponse[QueryExamListByExamIdResponse](err.Error(), "query exam history failed"))
		return
	}

	// exam为空，返回暂无考试记录
	// if exam == nil {
	// 	ctx.JSON(http.StatusAccepted, NewErrorResponse[QueryExamListByExamIdResponse]("", "exam is null"))
	// 	return
	// }

	s.logger.Printf("exam is ", exam)

	// &QueryExamListByExamIdData{ExamId: exam.Id}
	ctx.JSON(http.StatusOK, NewSuccessResponse(exam))
}
