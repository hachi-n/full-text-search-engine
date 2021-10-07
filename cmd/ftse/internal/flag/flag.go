package flag

import (
	"github.com/urfave/cli/v2"
)

const (
	TEXTFILE_FLAG_NAME = "textfile"
	NUMERIC_FLAG_NAME  = "numeric"
	VALUE_FLAG_NAME    = "value"
)

func TextFileFlag() *cli.StringFlag {
	return &cli.StringFlag{
		Name:     TEXTFILE_FLAG_NAME,
		Value:    "",
		Usage:    "textfile path.",
		Required: true,
	}
}

func ValueFlag() *cli.StringFlag {
	return &cli.StringFlag{
		Name:     VALUE_FLAG_NAME,
		Value:    "",
		Usage:    "search value.",
		Required: true,
	}
}

func NumericFlag() *cli.IntFlag {
	return &cli.IntFlag{
		Name:     NUMERIC_FLAG_NAME,
		Value:    1,
		Usage:    "numeric",
		Required: true,
	}
}
