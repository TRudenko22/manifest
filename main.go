package main

import (
  "embed"
  "fmt"
  "log"
  "os"

  "github.com/urfave/cli/v2"
)

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
  return nil
}

func main() {
  app := &cli.App{
    Commands: []*cli.Command{
      {
        Name: "make",
        Aliases: []string{"m", "mk"},
        Usage: "Generates a resource template",
        Action: generate,
      },
      {
        Name: "list",
        Aliases: []string{"l"},
        Usage: "Lists available templates",
        Action: list,
      },
    },
  }

  if err := app.Run(os.Args); err != nil {
    log.Fatal(err)
  }
}
