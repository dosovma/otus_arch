package commands //nolint:gochecknoglobals

import "github.com/dosovma/otus_arch/internal/app/entity"

const (
	IoCRegister   = "IoC.Register"
	ScopesCurrent = "Scopes.Current"
	ScopesNew     = "Scopes.New"
)

type IIoC interface {
	Resolve(key string, args ...any) entity.Executable
}

type IoC struct {
	threadLocal entity.IThreadLocal
	actions     map[string]func(args []any) entity.Executable
}

func NewIoc() *IoC {
	i := &IoC{
		threadLocal: entity.NewThreadLocal(),
	}
	i.initActions()

	return i
}

func (i *IoC) Resolve(key string, args ...any) entity.Executable {
	action, ok := i.actions[key]
	if !ok {
		return i.getBuilderAction([]any{key, args[0]})
	}

	return action(args)
}

func (i *IoC) initActions() {
	actions := make(map[string]func(args []any) entity.Executable)

	actions[IoCRegister] = i.registerAction
	actions[ScopesCurrent] = i.currentScopeAction
	actions[ScopesNew] = i.newScopeAction

	i.actions = actions
}
