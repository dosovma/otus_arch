package app

import (
	"fmt"

	"github.com/dosovma/otus_arch/internal/app/adapters"
	"github.com/dosovma/otus_arch/internal/app/commands"
	"github.com/dosovma/otus_arch/internal/app/entity"
)

// Run is an internal main func to run app.
func Run() {
	fmt.Println("app is running")

	moveBuilder := func(obj entity.UObject) entity.Executable {
		fmt.Println("It works")

		return commands.Move{
			Move: adapters.MovableAdapter{
				Obj: obj,
			},
		}
	}

	rotateBuilder := func(obj entity.UObject) entity.Executable {
		return commands.Rotate{
			Rotate: adapters.RotatableAdapter{
				Obj: obj,
			},
		}
	}

	i := commands.NewIoc()
	_ = i.Resolve("IoC.Register", "Commands.Move", moveBuilder).Execute()
	_ = i.Resolve("IoC.Register", "Commands.Rotate", rotateBuilder).Execute()
	_ = i.Resolve("Scopes.New", "new scope").Execute()
	_ = i.Resolve("Scopes.Current", "new scope").Execute()
	_ = i.Resolve("IoC.Register", "Commands.Move", moveBuilder).Execute()
	_ = i.Resolve("IoC.Register", "Commands.Rotate", rotateBuilder).Execute()
	_ = i.Resolve("Scopes.Current", "default").Execute()
	_ = i.Resolve("Commands.Move", entity.Object{}).Execute()
}
