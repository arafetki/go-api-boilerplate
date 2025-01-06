package api

type server interface {
	Start() error
}
type api struct {
	Server server
}

func New(srv server) *api {
	return &api{
		Server: srv,
	}
}
