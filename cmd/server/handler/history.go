package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/juanse1801/chatbot-naranja/internal/history"
	"github.com/juanse1801/chatbot-naranja/pkg/web"
)

type History struct {
	historyService history.Service
}

func NewHistory(h history.Service) *History {
	return &History{
		historyService: h,
	}
}

func (h *History) GetHistory() gin.HandlerFunc {
	return func(c *gin.Context) {
		date, _ := c.GetQuery("date")

		excel, err := h.historyService.GetHistory(c, date)

		if err != nil {
			web.Error(c, 400, "Error: Bad request")
		}

		c.Writer.Header().Set("Content-Type", "application/octet-stream")
		c.Writer.Header().Set("Content-Disposition", "attachment; filename="+date+".xlsx")
		c.Writer.Header().Set("Content-Transfer-Encoding", "binary")
		c.Writer.Header().Set("Expires", "0")

		excel.Write(c.Writer)

	}
}
