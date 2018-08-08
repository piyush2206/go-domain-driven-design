package dep

import "go.uber.org/fx"

// FxModule is fx module for complete dep package
var FxModule = fx.Options(
	fx.Provide(
		NewMySQLConn,
	),
)
