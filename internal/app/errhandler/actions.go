package errhandler

import (
	"github.com/dosovma/otus_arch/internal/app/commands"
	"github.com/dosovma/otus_arch/pkg"
)

// logAction обработчик исключения, который ставит Команду, пишущую в лог в очередь Команд.
func (h ErrorHandler) logAction(e pkg.Executable, err error) {
	h.queue.Push(
		commands.NewLogCmd(e, err, h.log),
	)
}

// firstRepeatAction обработчик исключения, который ставит в очередь Команду - повторитель команды, выбросившей исключение.
func (h ErrorHandler) firstRepeatAction(e pkg.Executable, err error) {
	h.queue.Push(
		commands.NewFirstRepeatCmd(e, err),
	)
}

// firstRepeatAction обработчик, который повторяет команду два раза.
func (h ErrorHandler) secondRepeatAction(e pkg.Executable, err error) {
	h.queue.Push(
		commands.NewSecondCmd(e, err),
	)
}
