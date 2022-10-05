package commands

import (
	"reflect"

	"github.com/dosovma/otus_arch/pkg"
)

// LogCmd команда, которая записывает информацию о выброшенном исключении в лог.
type LogCmd struct {
	ex  pkg.Executable
	err error
	log pkg.Logger
}

func NewLogCmd(ex pkg.Executable, err error, log pkg.Logger) LogCmd {
	return LogCmd{
		ex:  ex,
		err: err,
		log: log,
	}
}

func (l LogCmd) Execute() error {
	l.log.Error(
		"Command %s throw the error %s",
		reflect.TypeOf(l.ex).Name(),
		l.err.Error(),
	)

	return nil
}
