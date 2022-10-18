package app

import (
	"fmt"

	"github.com/dosovma/otus_arch/internal/app/commands"
)

// Run is an internal main func to run app.
func Run() {
	fmt.Println("app is running")

	// MacroCmd: движение по прямой с расходом топлива
	cmds := []commands.Executable{
		commands.CheckFuelCmd{},
		commands.Move{},
		commands.BurnFuelCmd{},
	}

	macroCmd := commands.NewMacroCmd(cmds...)

	_ = macroCmd.Execute()

	// MacroCmd: команда поворота с изменением вектора мгновенной скорости, если он есть
	cmds = []commands.Executable{
		commands.Rotate{},
		commands.ChangeVelocityCmd{},
	}

	macroCmd = commands.NewMacroCmd(cmds...)

	_ = macroCmd.Execute()
}
