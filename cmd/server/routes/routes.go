package routes

import (
	"github.com/juanse1801/chatbot-naranja/cmd/server/handler"
	"github.com/juanse1801/chatbot-naranja/internal/webhooks"

	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
}

func NewRouter(r *gin.Engine) Router {
	return &router{
		r: r,
	}
}

func (r *router) setGroup() {
	r.rg = r.r.Group("/api/v1")
}

func (r *router) MapRoutes() {
	r.setGroup()
	r.buildWebHookRoutes()
}

func (r *router) buildWebHookRoutes() {
	service := webhooks.NewService()
	handler := handler.NewWebHook(service)
	r.rg.GET("/webhooks", handler.GetValidate())
	r.rg.POST("/webhooks", handler.PostReceiveMessage())

}
