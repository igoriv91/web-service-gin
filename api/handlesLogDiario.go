package api

import "github.com/gin-gonic/gin"

func (s Server) HandleGetLogDiarioByDate(c *gin.Context) {
	l := NewLogDiario()
	data := c.Param("data")

	l.GetLogDiarioByDate(data)
}
