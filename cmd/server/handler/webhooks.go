package handler

import (
	"fmt"
	"strconv"

	"github.com/juanse1801/chatbot-naranja/internal/domain"
	"github.com/juanse1801/chatbot-naranja/internal/webhooks"
	"github.com/juanse1801/chatbot-naranja/pkg/web"

	"github.com/gin-gonic/gin"
)

type WebHooks struct {
	webhookService webhooks.Service
}

func NewWebHook(w webhooks.Service) *WebHooks {
	return &WebHooks{
		webhookService: w,
	}
}

func (w *WebHooks) GetValidate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		challengue, err := strconv.Atoi(ctx.Query("hub.challenge"))
		token := ctx.Query("hub.verify_token")

		if err != nil {
			web.Error(ctx, 400, "Error: id format")
			return
		}

		fmt.Println(token)

		isValid := w.webhookService.GetValidateEntry(ctx, token)

		if !isValid {
			web.Error(ctx, 401, "Error: Not authorized")
			return
		}

		ctx.JSON(200, challengue)
	}
}

func (w *WebHooks) PostReceiveMessage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newMessage domain.NewMessageReceived

		if err := ctx.ShouldBindJSON(&newMessage); err != nil {
			web.Error(ctx, 422, "wrong data")
			return
		}

		w.webhookService.PostReceiveMessage(ctx, newMessage)

		ctx.JSON(200, "")

	}
}
