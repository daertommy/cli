package plugin

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"

	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/plush/v4"
	"github.com/pkg/errors"

	"github.com/ignite/cli/ignite/pkg/gocmd"
	"github.com/ignite/cli/ignite/pkg/xgenny"
)

//go:embed template/*
var fsPluginSource embed.FS

// Scaffold generates a plugin structure under dir/path.Base(moduleName).
func Scaffold(dir, moduleName string, sharedHost bool) (string, error) {
	var (
		name     = filepath.Base(moduleName)
		finalDir = path.Join(dir, name)
	)
	if _, err := os.Stat(finalDir); err == nil {
		// finalDir already exists, don't overwrite stuff
		return "", errors.Errorf("directory %q already exists, abort scaffolding", finalDir)
	}

	// Remove "files/" prefix
	subfs, err := fs.Sub(fsPluginSource, "template")
	if err != nil {
		return "", fmt.Errorf("template sub: %w", err)
	}

	g := genny.New()
	if err := g.FS(subfs); err != nil {
		return "", fmt.Errorf("template fs: %w", err)
	}

	ctx := plush.NewContext()
	ctx.Set("ModuleName", moduleName)
	ctx.Set("Name", name)
	ctx.Set("SharedHost", sharedHost)

	g.Transformer(xgenny.Transformer(ctx))
	r := genny.WetRunner(ctx)
	if err := r.With(g); err != nil {
		return "", errors.WithStack(err)
	}
	if err := r.Run(); err != nil {
		return "", errors.WithStack(err)
	}
	if err := gocmd.ModTidy(context.TODO(), finalDir); err != nil {
		return "", errors.WithStack(err)
	}
	return finalDir, nil
}
