package main

import (
	"fmt"
	"os"

	"github.com/iancoleman/strcase"
	"github.com/urfave/cli"
)

const version = "0.2.4"

func main() {
	app := cli.NewApp()

	app.Usage = ""
	app.Version = version

	app.Commands = []cli.Command{
		{
			Name:    "snake",
			Aliases: []string{"s"},
			Usage:   "Convert input to snake_case",
			Action: func(c *cli.Context) error {
				fmt.Println(strcase.ToSnake(c.Args().First()))
				return nil
			},
		},
		{
			Name:    "screaming-snake",
			Aliases: []string{"ss"},
			Usage:   "Convert input to SCREAMING_SNAKE_CASE",
			Action: func(c *cli.Context) error {
				fmt.Println(strcase.ToScreamingSnake(c.Args().First()))
				return nil
			},
		},
		{
			Name:    "kebab",
			Aliases: []string{"k"},
			Usage:   "Convert input to kebab-case",
			Action: func(c *cli.Context) error {
				fmt.Println(strcase.ToKebab(c.Args().First()))
				return nil
			},
		},
		{
			Name:    "screaming-kebab",
			Aliases: []string{"sk"},
			Usage:   "Convert input to SCREAMING-KEBAB-CASE",
			Action: func(c *cli.Context) error {
				fmt.Println(strcase.ToScreamingKebab(c.Args().First()))
				return nil
			},
		},
		{
			Name:    "camel",
			Aliases: []string{"c"},
			Usage:   "Convert input to CamelCase",
			Action: func(c *cli.Context) error {
				fmt.Println(strcase.ToCamel(c.Args().First()))
				return nil
			},
		},
		{
			Name:    "lower-camel",
			Aliases: []string{"lc"},
			Usage:   "Convert input to lowerCamelCase",
			Action: func(c *cli.Context) error {
				fmt.Println(strcase.ToLowerCamel(c.Args().First()))
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
