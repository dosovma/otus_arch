package errhandler

import (
	"reflect"

	"github.com/dosovma/otus_arch/internal/app/commands"
	"github.com/dosovma/otus_arch/pkg"
)

type ErrorHandler struct {
	errToAction map[string]map[error]func(e pkg.Executable, err error)
	queue       pkg.IQueue
	log         pkg.Logger
}

// NewOneRepeatStrategic хендлер реализующий стратегию:
// при первом выбросе исключения повторить команду, при повторном выбросе исключения записать информацию в лог.
func NewOneRepeatStrategic(queue pkg.IQueue, log pkg.Logger) ErrorHandler {
	errHandler := ErrorHandler{}
	errHandler.queue = queue
	errHandler.log = log

	errToAction := make(map[string]map[error]func(e pkg.Executable, err error))

	errToAction["BaseCmd"] = make(map[error]func(e pkg.Executable, err error))
	errToAction["BaseCmd"][commands.ErrConnectionTimeout] = errHandler.firstRepeatAction

	errToAction["FirstRepeatCmd"] = make(map[error]func(e pkg.Executable, err error))
	errToAction["FirstRepeatCmd"][commands.ErrConnectionTimeout] = errHandler.logAction

	errHandler.errToAction = errToAction

	return errHandler
}

// NewTwoRepeatStrategic хендлер реализующий стратегию:
// повторить два раза, потом записать в лог.
func NewTwoRepeatStrategic(queue pkg.IQueue, log pkg.Logger) ErrorHandler {
	errHandler := ErrorHandler{}
	errHandler.queue = queue
	errHandler.log = log

	errToAction := make(map[string]map[error]func(e pkg.Executable, err error))

	errToAction["BaseCmd"] = make(map[error]func(e pkg.Executable, err error))
	errToAction["BaseCmd"][commands.ErrConnectionTimeout] = errHandler.firstRepeatAction

	errToAction["FirstRepeatCmd"] = make(map[error]func(e pkg.Executable, err error))
	errToAction["FirstRepeatCmd"][commands.ErrConnectionTimeout] = errHandler.secondRepeatAction

	errToAction["SecondRepeatCmd"] = make(map[error]func(e pkg.Executable, err error))
	errToAction["SecondRepeatCmd"][commands.ErrConnectionTimeout] = errHandler.logAction

	errHandler.errToAction = errToAction

	return errHandler
}

func (h ErrorHandler) Handle(cmd pkg.Executable, err error) {
	cmdType := reflect.TypeOf(cmd).Name()

	action, ok := h.errToAction[cmdType][err]
	if !ok {
		h.log.Error("There is no action or error for command %s", cmdType)

		return
	}

	action(cmd, err)
}
