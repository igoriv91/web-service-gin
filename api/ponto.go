package api

import (
	"time"

	"github.com/google/uuid"
)

type Ponto struct {
	Id   string    `json:"id"`
	Data time.Time `json:"data"`
}

func NewPonto() *Ponto {
	return &Ponto{
		Id:   uuid.New().String(),
		Data: time.Now(),
	}
}

func (p *Ponto) GetPontoByData(data string) *Ponto {
	return NewPonto()
}

func (p *Ponto) PostPonto(newPonto *Ponto) *Ponto {
	return newPonto
}

func (p *Ponto) PutPonto(currentPonto *Ponto) *Ponto {
	return currentPonto
}

func (p *Ponto) DeletePonto(currentPonto *Ponto) *Ponto {
	return currentPonto
}
