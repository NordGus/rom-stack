package staticfileserver

import "embed"

var (
	//go:embed dist public
	fs embed.FS
)

type Service struct {
	fs embed.FS
}

func New() (*Service, error) {
	return &Service{
		fs: fs,
	}, nil
}
