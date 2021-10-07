package indexing

import (
	"github.com/hachi-n/full-text-search-engine/cmd/ftse/internal/flag"
	"github.com/hachi-n/full-text-search-engine/internal/entrypoint/ftse/ngram/indexing"
	"github.com/urfave/cli/v2"
)

func Cmd() *cli.Command {
	cmd := &cli.Command{
		Name:  "indexing",
		Usage: "indexing ",
		Flags: []cli.Flag{
			flag.NumericFlag(),
			flag.TextFileFlag(),
		},
		Action: func(c *cli.Context) error {
			textFilePath := c.String(flag.TEXTFILE_FLAG_NAME)
			numeric := c.Int(flag.NUMERIC_FLAG_NAME)
			return indexing.Apply(numeric, textFilePath)
		},
	}
	return cmd
}
