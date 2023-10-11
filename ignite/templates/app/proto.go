package app

import (
	"embed"
	"fmt"
	"io/fs"

	"github.com/gobuffalo/genny/v2"
)

//go:embed files/proto/* files/buf.work.yaml
var fsProto embed.FS

// NewBufGenerator returns the generator to buf build files.
func NewBufGenerator(appPath string) (*genny.Generator, error) {
	g := genny.New()
	// Remove "files/" prefix
	subfs, err := fs.Sub(fsProto, "files")
	if err != nil {
		return nil, fmt.Errorf("generator sub: %w", err)
	}

	if err := g.FS(subfs); err != nil {
		return g, fmt.Errorf("generator fs: %w", err)
	}

	return g, nil
}
