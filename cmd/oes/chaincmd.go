package main

import (
	"net/http"
	"gopkg.in/urfave/cli.v2"
	"github.com/gin-gonic/gin"
)

var (
	deCommand = &cli.Command{
		Action:  de,
		Name:    "de",
		Aliases: []string{"d"},
		Usage:   "Download event log from blockchain",
		Flags: []cli.Flag{
		},
	}
	serveCommand = &cli.Command{
		Action:  serve,
		Name:    "serve",
		Aliases: []string{"s"},
		Usage:   "serve as a HTTP server",
		Flags: []cli.Flag{
		},
	}
	listenCommand = &cli.Command{
		Action:  listen,
		Name:    "listen",
		Aliases: []string{"l"},
		Usage:   "Listen to blockchain and process event log",
		Flags: []cli.Flag{
		},
	}
)

func de(ctx *cli.Context) error {
	return nil
}

func serve(ctx *cli.Context) error {
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.Run(":8080")
	return nil
}

func listen(ctx *cli.Context) error {
	return nil
}
