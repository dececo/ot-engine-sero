package main

import (
	"path/filepath"
	"os"
	"fmt"
	"gopkg.in/urfave/cli.v2"
	"github.com/dececo/ot-engine-sero/cmd/utils"
)

var app *cli.App

func init() {
	app = &cli.App{
		Name:      filepath.Base(os.Args[0]),
		Usage:     "The oes(OpenTask Engine on SERO) command line interface",
		Version:   "0.1.0",
		Action:    oes,
	}
	Info(app)
	app.Commands = []*cli.Command{
		serveCommand,
		deCommand,
		listenCommand,
	}
	app.Flags = []cli.Flag{
		utils.ConfigFlag,
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func oes(ctx *cli.Context) error {
	return nil
}