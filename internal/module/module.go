package module

import "go.uber.org/fx"

var Options = fx.Options(
	fx.Provide(LoadConfig),
)
