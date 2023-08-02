package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (s Server) Start() {
	router := gin.Default()

	router.Use(handle)
	groupLogDiario := router.Group("logdiario")
	groupLogDiario.GET("/:data", s.HandleGetLogDiarioByDate)

	groupPonto := router.Group("ponto")
	groupPonto.GET("/:data", s.HandleGetPonto)
	groupPonto.POST("/", s.HandlePostPonto)
	groupPonto.PUT("/", s.HandlePutPonto)
	groupPonto.DELETE("/:uuid", s.HandleDeletePonto)

	router.Run("localhost:8080")
}

func handle(c *gin.Context) {
	c.Next()
	if len(c.Errors) == 0 {
		return
	}

}
