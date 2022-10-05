package commands

import (
	"github.com/dosovma/otus_arch/pkg"
)

// SecondRepeatCmd команда, которая второй раз повторяет Команду, выбросившую исключение.
type SecondRepeatCmd struct {
	ex  pkg.Executable
	err error
}

func NewSecondCmd(ex pkg.Executable, err error) SecondRepeatCmd {
	return SecondRepeatCmd{
		ex:  ex,
		err: err,
	}
}

func (s SecondRepeatCmd) Execute() error {
	return s.ex.Execute()
}
