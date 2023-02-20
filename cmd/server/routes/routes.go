package routes

import (
	"github.com/go-co-op/gocron"
	"github.com/juanse1801/chatbot-naranja/cmd/server/handler"
	"github.com/juanse1801/chatbot-naranja/internal/interaction"
	"github.com/juanse1801/chatbot-naranja/internal/messaging"
	"github.com/juanse1801/chatbot-naranja/internal/scheduler"
	"github.com/juanse1801/chatbot-naranja/internal/state"
	"github.com/juanse1801/chatbot-naranja/internal/webhooks"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r   *gin.Engine
	rg  *gin.RouterGroup
	sch *gocron.Scheduler
	db  *mongo.Client
}

func NewRouter(r *gin.Engine, sch *gocron.Scheduler, db *mongo.Client) Router {
	return &router{
		r:   r,
		sch: sch,
		db:  db,
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
	itcRepository := interaction.NewRepository(r.db)
	mssgService := messaging.NewService()
	schService := scheduler.NewService(r.sch)
	stateService := state.NewService()
	itcService := interaction.NewService(itcRepository)
	service := webhooks.NewService(itcService, schService, stateService, mssgService)
	handler := handler.NewWebHook(service)
	r.rg.GET("/webhooks", handler.GetValidate())
	r.rg.POST("/webhooks", handler.PostReceiveMessage())

}
