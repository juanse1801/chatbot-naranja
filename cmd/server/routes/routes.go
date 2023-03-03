package routes

import (
	"github.com/go-co-op/gocron"
	"github.com/juanse1801/chatbot-naranja/cmd/server/handler"
	"github.com/juanse1801/chatbot-naranja/internal/history"
	"github.com/juanse1801/chatbot-naranja/internal/interaction"
	"github.com/juanse1801/chatbot-naranja/internal/mailing"
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
	r.buildHistoryRoutes()
}

func (r *router) buildHistoryRoutes() {
	historyRepository := history.NewRepository(r.db)
	historyService := history.NewService(historyRepository)
	handler := handler.NewHistory(historyService)
	r.rg.GET("/history", handler.GetHistory())
}

func (r *router) buildWebHookRoutes() {
	historyRepository := history.NewRepository(r.db)
	historyService := history.NewService(historyRepository)
	itcRepository := interaction.NewRepository(r.db)
	mssgService := messaging.NewService()
	stateService := state.NewService()
	itcService := interaction.NewService(itcRepository)
	schService := scheduler.NewService(r.sch, mssgService, itcService)
	mailService := mailing.NewService()
	service := webhooks.NewService(itcService, schService, stateService, mssgService, mailService, historyService)
	handler := handler.NewWebHook(service)
	r.rg.GET("/webhooks", handler.GetValidate())
	r.rg.POST("/webhooks", handler.PostReceiveMessage())

}
