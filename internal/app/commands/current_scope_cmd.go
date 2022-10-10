package commands

import "github.com/dosovma/otus_arch/internal/app/entity"

type CurrentScopeCmd struct {
	tl        entity.IThreadLocal
	scopeName string
}

func NewCurrentScopeCmd(tl entity.IThreadLocal, scopeName string) CurrentScopeCmd {
	return CurrentScopeCmd{
		tl:        tl,
		scopeName: scopeName,
	}
}

func (s CurrentScopeCmd) Execute() error {
	scope, _ := s.tl.GetValue(s.scopeName)
	s.tl.SetCurrentScope(scope)

	return nil
}
