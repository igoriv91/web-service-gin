package api

import (
	"example/web-service-gin/types"

	"github.com/gin-gonic/gin"
)

func (s Server) HandleGetLogDiarioByDate(c *gin.Context) {
	l := types.NewLogDiario()
	data := c.Param("data")

	l.GetLogDiarioByDate(data)
}
