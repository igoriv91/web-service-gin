package api

import (
	"time"
)

type LogDiario struct {
	Duration  int64     `json:"duration"`
	Descricao string    `json:"description"`
	Data      time.Time `json:"data"`
}

func NewLogDiario() *LogDiario {
	return &LogDiario{
		Duration:  0,
		Descricao: "",
		Data:      time.Now(),
	}
}

func (l *LogDiario) GetLogDiarioByDate(data string) *LogDiario {

	return NewLogDiario()
}
