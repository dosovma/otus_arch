package commands

import "fmt"

type MacroCmd struct {
	Commands []Executable
}

func NewMacroCmd(commands ...Executable) MacroCmd {
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
