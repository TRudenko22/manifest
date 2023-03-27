package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

//TODO: add jobs to templates
//TODO: create -o flag to generate a file
//TODO: add jobs to templates

var (
	//go:embed templates/*
	templates embed.FS
)

func generate(ctx *cli.Context) error {
	file := ctx.Args().Get(0)

	template, err := templates.ReadFile(fmt.Sprintf("templates/%s.yaml", file))
	if err != nil {
		return fmt.Errorf("Template not found")
	} else {
		fmt.Println(string(template))
		return nil
	}
}

func list(ctx *cli.Context) error {
	dir, _ := templates.ReadDir("templates")

	fmt.Println("The following templates are in this binary:")
	fmt.Println("-------------------------------------------")
	for _, file := range dir {
		resource := strings.Split(file.Name(), ".")[0]
		fmt.Println(resource)
	}
	fmt.Println("-------------------------------------------")

	return nil
}

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "make",
				Aliases: []string{"m", "mk"},
				Usage:   "Generates a resource template",
				Action:  generate,
			},
			{
				Name:    "list",
				Aliases: []string{"l", "ls"},
				Usage:   "Lists available templates",
				Action:  list,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
