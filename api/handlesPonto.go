package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s Server) HandleGetPonto(c *gin.Context) {
	p := NewPonto()
	data := c.Param("data")
	if data == "" {
		c.Error(errors.New("parametro <data> é obrigatório"))
		return
	}
	p.GetPontoByData(data)
}

func (s Server) HandlePostPonto(c *gin.Context) {
	var p Ponto
	if err := c.BindJSON(&p); err != nil {
		c.Error(err)
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, p.PostPonto(&p))
}

func (s Server) HandlePutPonto(c *gin.Context) {
	var p Ponto
	if err := c.BindJSON(&p); err != nil {
		c.Error(err)
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, p.PutPonto(&p))
}

func (s Server) HandleDeletePonto(c *gin.Context) {
	var p Ponto
	if err := c.BindJSON(&p); err != nil {
		c.Error(err)
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, p.DeletePonto(&p))
}
