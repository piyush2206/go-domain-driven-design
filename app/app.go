// app package is responible for initialising
// all the external dependencies and their mocking

package app

type (
	// AppCtx struct contains
	AppCtx struct {
		DB IDb
	}
)

// Init etablishes connections to multiple external dependency of the project
// and returns them in a single context object
func Init() (ctx *AppCtx, err error) {
	ctx = &AppCtx{}

	if ctx.DB, err = initMySQL(); err != nil {
		return nil, err
	}

	return ctx, nil
}
