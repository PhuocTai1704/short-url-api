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
	r.GET("/:code", ur.handler.Redirect)

	links := r.Group("/links")
	{
		links.GET("", ur.handler.GetAllLinks)
		links.GET("/:id", ur.handler.GetById)
		links.GET("/link", ur.handler.GetByLink)
		links.POST("", ur.handler.CreateLink)
	}
}
