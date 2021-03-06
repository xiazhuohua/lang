package main

import (
	"github.com/go-easygen/cli"
)

type argT struct {
	cli.Helper
	Id uint8 `cli:"*id" usage:"this is a required flag, note the *"`
}

func main() {
	cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		ctx.String("%d\n", argv.Id)
		return nil
	})
}
