package commands

import "github.com/dosovma/otus_arch/pkg"

// FirstRepeatCmd команда, которая повторяет Команду, выбросившую исключение.
type FirstRepeatCmd struct {
	ex  pkg.Executable
	err error
}

func NewFirstRepeatCmd(ex pkg.Executable, err error) FirstRepeatCmd {
	return FirstRepeatCmd{
		ex:  ex,
		err: err,
	}
}

func (f FirstRepeatCmd) Execute() error {
	return f.ex.Execute()
}
