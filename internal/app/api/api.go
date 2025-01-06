package api

type Server interface {
	Start() error
}

type API struct {
	Server Server
}
