package main

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type doer struct {
	filesPath string
}

func newDoer(folderPath, filesType string) (*doer, error) {
	fpath := filepath.Join(folderPath, filesType)
	if _, err := os.Stat(fpath); err != nil {
		return nil, errors.Wrapf(err, "file path:%s detect failed", fpath)
	}

	return &doer{
		filesPath: fpath,
	}, nil
}

// func (d *doer) do() error {

// }
