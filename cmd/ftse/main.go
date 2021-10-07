package main

import (
	"github.com/hachi-n/full-text-search-engine/cmd/ftse/ngram"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "ftse",
		Usage: "full-text-search-engine",
		Commands: []*cli.Command{
			ngram.Cmd(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Fatal Error: %v", err)
	}
}
