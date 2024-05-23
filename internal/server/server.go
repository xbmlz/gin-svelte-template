package server

// IServer Server is transport server.
type IServer interface {
	Start() error
	Shutdown() error
}
