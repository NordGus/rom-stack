package staticfileserver

import "embed"

var (
	//go:embed dist
	dist embed.FS
)

type App struct {
	dist embed.FS
}

func New() (*App, error) {
	return &App{
		dist: dist,
	}, nil
}
