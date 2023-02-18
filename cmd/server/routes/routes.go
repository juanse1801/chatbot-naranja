package routes

import (
	"github.com/go-co-op/gocron"
	"github.com/juanse1801/chatbot-naranja/cmd/server/handler"
	"github.com/juanse1801/chatbot-naranja/internal/webhooks"

	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r   *gin.Engine
	rg  *gin.RouterGroup
	sch *gocron.Scheduler
}

func NewRouter(r *gin.Engine, sch *gocron.Scheduler) Router {
	return &router{
		r:   r,
		sch: sch,
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
