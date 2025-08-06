package http

import "github.com/gin-gonic/gin"

type Service interface {
	RegisterRouter(*gin.RouterGroup)
}
