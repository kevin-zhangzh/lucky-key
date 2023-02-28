package main

import (
	"github.com/kevin-zhangzh/lucky-key"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := &cli.App{
		Name: "Lucky",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "sqlite_dir", Value: "./data/sqlite", Usage: "sqlite db dir path", EnvVars: []string{"SQLITE_DIR"}},

			&cli.StringFlag{Name: "web3", Value: "https://mainnet.infura.io/v3/3ddf2116829f4cd8ace7e0070568fc4a", Usage: "eth rpc", EnvVars: []string{"WEB3"}},
		},
		Action: run,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	db := c.String("sqlite_dir")
	rpc := c.String("web3")
	s := lucky.NewLucky(db, rpc)
	s.Run()

	<-signals

	return nil
}
