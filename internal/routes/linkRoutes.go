package routes

import (
	"short-url-api/internal/handler"

	"github.com/gin-gonic/gin"
)

type LinkRoutes struct {
	handler *handler.LinkHandler
}

func NewLinkRoutes(handler *handler.LinkHandler) *LinkRoutes {
	return &LinkRoutes{
		handler: handler,
	}
}

func (ur *LinkRoutes) Register(r *gin.RouterGroup) {
	links := r.Group("/links")
	{
		links.GET("/test", ur.handler.TestHandler)
	}
}