// Package main is the main package for the cli/calculator
package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "calculator",
		Usage:   "Bring some sparkles in your life",
		Version: "v0.0.1",
		Authors: []*cli.Author{
			{
				Name:  "Kevin LARQUEMIN",
				Email: "kevin.larquemin@gmail.com",
			},
		},
		Action: func(c *cli.Context) error {
			if c.Bool("interactive") {
				interactive()
			} else {
				inline()
			}
			return nil
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "interactive",
				Aliases: []string{"i"},
				Usage:   "start an interactive calculator",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
