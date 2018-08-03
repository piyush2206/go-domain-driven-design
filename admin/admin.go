// admin package is responible for all the school entities
// and their repositories and services

package admin

import "github.com/piyush2206/go-domain-driven-design/app"

var (
	RepoClass   *RepositoryClass
	RepoStudent *RepositoryStudent
)

// Init initialises all the admin package repositories
func Init(appCtx *app.AppCtx) {
	RepoClass = new(RepositoryClass)
	RepoClass.Init(appCtx)

	RepoStudent = new(RepositoryStudent)
	RepoStudent.Init(appCtx)
}
