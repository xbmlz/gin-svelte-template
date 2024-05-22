package server

// Server is transport server.
type IServer interface {
	Start() error
	Shutdown() error
}
