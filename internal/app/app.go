package app

import (
	"github.com/dosovma/otus_arch/internal/app/commands"
	"github.com/dosovma/otus_arch/internal/app/errhandler"
	"github.com/dosovma/otus_arch/pkg"
)

// Run is an internal main func to run app.
func Run() {
	log := pkg.Log{}

	log.Error("app is running")

	queue := pkg.NewQueue()
	errHandler := errhandler.NewTwoRepeatStrategic(queue, log)

	baseCmd := commands.BaseCmd{}
	queue.Push(baseCmd)

	for {
		cmd, err := queue.Pull()
		if err != nil {
			log.Error(err.Error())

			return
		}

		if err = cmd.Execute(); err != nil {
			errHandler.Handle(cmd, err)
		}
	}
}
