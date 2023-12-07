package staticfileserver

import "embed"

var (
	//go:embed dist
	dist embed.FS
)

type Service struct {
	dist embed.FS
}

func New() (*Service, error) {
	return &Service{
		dist: dist,
	}, nil
}
