package commands

import (
	"fmt"

	"github.com/dosovma/otus_arch/internal/app/entity"
)

type MacroCmd struct {
	Commands []entity.Executable
}

func NewMacroCmd(commands ...entity.Executable) MacroCmd {
	return MacroCmd{Commands: commands}
}

func (m MacroCmd) Execute() error {
	for _, cmd := range m.Commands {
		if err := cmd.Execute(); err != nil {
			return fmt.Errorf("%s: %w", err.Error(), ErrMacroCmd)
		}
	}

	return nil
}
