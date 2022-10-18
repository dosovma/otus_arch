package pkg

import (
	"github.com/rs/zerolog/log"
)

//go:generate mockgen -destination=./mocks/logger.go -package=mocks -source=logger.go
type Logger interface {
	Error(format string, v ...any)
}

type Log struct{}

func (l Log) Error(format string, v ...any) {
	log.Error().Msgf(format, v...)
}
