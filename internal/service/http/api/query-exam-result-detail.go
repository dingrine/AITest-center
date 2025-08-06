package api

import (
	"github.com/gin-gonic/gin"
)

type QueryExamResultDetailData struct {
}

type QueryExamResultDetailResponse = Response[QueryExamResultDetailData]

// QueryExamResultDetail godoc
//
//	@Summary		query nft by token id
//	@Description	query nft by NFT's id
//	@Tags			Query NFT by token id
//	@Accept			json
//	@Produce		json
//	@Param			tokenId	path		string							true	"nft's id"
//	@Success		200		{object}	QueryExamResultDetailResponse	"新导出NFT成功。NFT的导出流程会在后台自动进行"
//	@Success		202		{object}	QueryExamResultDetailResponse	"后台服务正在创建NFT"
//	@Failure		404		{object}	QueryExamResultDetailResponse	"传入的 token id 并没有被请求导出过"
//	@Failure		400		{object}	QueryExamResultDetailResponse	"参数错误"
//	@Failure		500		{object}	QueryExamResultDetailResponse
//	@Router			/v1/api/nft/{tokenId} [get]
func (s *Service) QueryExamResultDetail(ctx *gin.Context) {
}
