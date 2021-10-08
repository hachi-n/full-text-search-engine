package architecture

import (
	"os"
	"path/filepath"
)

type Architecture struct{}

var (
	DatabaseDir = filepath.Join(os.Getenv("HOME"), ".ftse")
	DataDir     = filepath.Join(DatabaseDir, "data")
	IndexDir    = filepath.Join(DatabaseDir, "index")
)

func (a *Architecture) defaultDirectories() []string {
	directories := []string{
		DataDir,
		IndexDir,
	}
	return directories
}

func (a *Architecture) Create() error {
	directories := a.defaultDirectories()
	if err := a.validate(directories); err == nil {
		return nil
	}
	for _, directory := range directories {
		if err := os.MkdirAll(directory, 0755); err != nil {
			return err
		}
	}
	return nil
}

func (a *Architecture) Validate() error {
	directories := a.defaultDirectories()
	return a.validate(directories)
}

func (a *Architecture) validate(directories []string) error {
	for _, d := range directories {
		if _, err := os.Stat(d); err != nil {
			return err
		}
	}
	return nil
}