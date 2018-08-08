// admin package is responible for all the school entities
// and their repositories and services

package admin

import (
	"go.uber.org/fx"
)

// FxModule is fx module for complete admin package
var FxModule = fx.Options(
	ClassModule,
	StudentModule,
)
