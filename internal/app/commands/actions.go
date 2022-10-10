package commands

import (
	"fmt"

	"github.com/dosovma/otus_arch/internal/app/entity"
)

func (i *IoC) registerAction(args []any) entity.Executable {
	cmdName, ok := args[0].(string)
	if !ok {
		return NewErrorCmd(
			fmt.Sprintf("invalid command name %v for action %s", args[0], IoCRegister),
		)
	}

	cmdBuilder, ok := args[1].(func(obj entity.UObject) entity.Executable)
	if !ok {
		return NewErrorCmd(
			fmt.Sprintf("invalid command b %v for action %s", args[1], IoCRegister),
		)
	}

	return NewRegisterCmd(i.threadLocal, cmdName, cmdBuilder)
}

func (i *IoC) newScopeAction(args []any) entity.Executable {
	scopeName, ok := args[0].(string)
	if !ok {
		return NewErrorCmd(
			fmt.Sprintf("invalid scope name %v for action %s", args[0], ScopesNew),
		)
	}

	return NewNewScopeCmd(i.threadLocal, scopeName)
}

func (i *IoC) currentScopeAction(args []any) entity.Executable {
	scopeName, ok := args[0].(string)
	if !ok {
		return NewErrorCmd(
			fmt.Sprintf("invalid scope name %v for action %s", args[0], ScopesCurrent),
		)
	}

	return NewCurrentScopeCmd(i.threadLocal, scopeName)
}

func (i *IoC) getBuilderAction(args []any) entity.Executable {
	cmdName, ok := args[0].(string)
	if !ok {
		return NewErrorCmd(
			fmt.Sprintf("invalid command name %v for action get command b by command name", args[0]),
		)
	}

	obj, ok := args[1].(entity.UObject)
	if !ok {
		return NewErrorCmd(
			fmt.Sprintf("invalid obj %v for action get command b by command name", args[1]),
		)
	}

	return NewGetBuilderCmd(i.threadLocal, cmdName, obj)
}
