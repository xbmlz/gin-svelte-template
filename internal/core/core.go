package core

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewHTTPServer),
	fx.Provide(NewConfig),
	fx.Provide(NewLogger),
	fx.Provide(NewDatabase),
)
