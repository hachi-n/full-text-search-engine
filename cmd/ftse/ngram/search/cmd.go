package search

import (
	"fmt"
	"github.com/hachi-n/full-text-search-engine/cmd/ftse/internal/flag"
	"github.com/hachi-n/full-text-search-engine/internal/entrypoint/ftse/ngram/search"
	"github.com/urfave/cli/v2"
)

func Cmd() *cli.Command {
	cmd := &cli.Command{
		Name:  "search",
		Usage: "search",
		Flags: []cli.Flag{
			flag.NumericFlag(),
			flag.ValueFlag(),
		},
		Action: func(c *cli.Context) error {
			value := c.String(flag.VALUE_FLAG_NAME)
			numeric := c.Int(flag.NUMERIC_FLAG_NAME)
			// temporary validation. unigram only
			if numeric != 1 {
				return fmt.Errorf("sorry. numeric is \"1\" only")
			}
			return search.Apply(numeric, value)
		},
	}
	return cmd
}
