package commands

import "github.com/dosovma/otus_arch/internal/app/entity"

type NewScopeCmd struct {
	tl        entity.IThreadLocal
	scopeName string
}

func NewNewScopeCmd(tl entity.IThreadLocal, scopeName string) NewScopeCmd {
	return NewScopeCmd{
		tl:        tl,
		scopeName: scopeName,
	}
}

func (s NewScopeCmd) Execute() error {
	newScope := make(map[string]func(obj entity.UObject) entity.Executable)

	s.tl.SetValue(s.scopeName, newScope)

	return nil
}
