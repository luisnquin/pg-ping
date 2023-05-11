package main

import (
	"fmt"
	"os"

	"github.com/luisnquin/pg-ping/cmd"
	"github.com/urfave/cli"
)

var version string

func main() {
	app := cli.NewApp()

	if version == "" {
		version = "dev"
	}

	app.Version = version

	cli.HelpFlag = cli.BoolFlag{Name: "help"}

	if err := cmd.Execute(app); err != nil {
		fmt.Fprintf(os.Stderr, "pg-ping failed: %v\n", err)
		os.Exit(1)
	}
}
