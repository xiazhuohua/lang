package main

import (
	"github.com/go-easygen/cli"
	clix "github.com/go-easygen/cli/ext"
)

type argT struct {
	cli.Helper
	PidFile clix.PidFile `cli:"pid" usage:"pid file" dft:"013-pidfile.pid"`
}

func main() {
	cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)

		if err := argv.PidFile.New(); err != nil {
			return err
		}
		defer argv.PidFile.Remove()

		return nil
	})
}
