package main

import (
	"net/http"
	"gopkg.in/urfave/cli.v2"
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/sero-cash/go-sero/seroclient"
	"log"
	"github.com/sero-cash/go-sero/common"
	"github.com/sero-cash/go-sero"
	"context"
	"github.com/sero-cash/go-sero/core/types"
	"github.com/dececo/ot-engine-sero/process"
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

func de(ctx *cli.Context) (err error) {
	stack := makeConfigNode(ctx)
	fmt.Printf("server: %s, contract: %s\n", stack.Config.Server, stack.Config.Contract)

	client, err := seroclient.Dial(stack.Config.Server)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("we have a connection now.")
	}

	address := common.Base58ToAddress(stack.Config.Contract)
	query := sero.FilterQuery{
		Addresses: []common.Address{address},
	}
	if stack.Config.FromFlag {
		query.FromBlock = stack.Config.FromBlock
	}
	if stack.Config.ToFlag {
		query.ToBlock = stack.Config.ToBlock
	}
	logs, err := client.FilterLogs(context.Background(), query)

	if err != nil {
		log.Fatal(err)
	}

	for i, vLog := range logs {
		fmt.Printf("TxHash[%d]: %s\n", i, vLog.TxHash.Hex())
		err := process.ParseOTLog(vLog)
		if err != nil {
			continue
		}
	}
	return err
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

func listen(ctx *cli.Context) (err error) {
	stack := makeConfigNode(ctx)
	fmt.Printf("server: %s, contract: %s\n", stack.Config.Server, stack.Config.Contract)

	client, err := seroclient.Dial(stack.Config.Server)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("we have a connection now.")
	}

	address := common.Base58ToAddress(stack.Config.Contract)
	query := sero.FilterQuery{
		Addresses: []common.Address{address},
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)

	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			fmt.Println(err)
		case vLog := <-logs:
			err := process.ParseOTLog(vLog)
			if err != nil {
				continue
			}
		}
	}
	return
}
