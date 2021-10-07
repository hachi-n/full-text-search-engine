package ngram

import (
	"github.com/hachi-n/full-text-search-engine/cmd/ftse/ngram/indexing"
	"github.com/hachi-n/full-text-search-engine/cmd/ftse/ngram/search"
	"github.com/urfave/cli/v2"
)

func Cmd() *cli.Command {
	cmd := &cli.Command{
		Name:  "ngram",
		Usage: "ngram commands",
		Subcommands: []*cli.Command{
			indexing.Cmd(),
			search.Cmd(),
		},
	}
	return cmd
}
