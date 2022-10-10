package commands

import (
	"fmt"

	"github.com/dosovma/otus_arch/internal/app/entity"
)

type GetBuilderCmd struct {
	tl      entity.IThreadLocal
	cmdName string
	obj     entity.UObject
}

func NewGetBuilderCmd(tl entity.IThreadLocal, cmdName string, obj entity.UObject) GetBuilderCmd {
	return GetBuilderCmd{
		tl:      tl,
		cmdName: cmdName,
		obj:     obj,
	}
}

func (b GetBuilderCmd) Execute() error {
	scope := b.tl.GetCurrentScope()

	builder, ok := scope[b.cmdName]
	if !ok {
		return fmt.Errorf("not found b for command %s", b.cmdName)
	}

	builder(b.obj)

	return nil
}
