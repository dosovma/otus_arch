package commands

import "github.com/dosovma/otus_arch/internal/app/entity"

type RegisterCmd struct {
	tl      entity.IThreadLocal
	cmdName string
	b       func(obj entity.UObject) entity.Executable
}

func NewRegisterCmd(tl entity.IThreadLocal, cmdName string, builder func(obj entity.UObject) entity.Executable) RegisterCmd {
	return RegisterCmd{
		tl:      tl,
		cmdName: cmdName,
		b:       builder,
	}
}

func (r RegisterCmd) Execute() error {
	scope := r.tl.GetCurrentScope()
	scope[r.cmdName] = r.b

	return nil
}
