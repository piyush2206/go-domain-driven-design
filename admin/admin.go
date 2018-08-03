package admin

import "github.com/piyush2206/go-domain-driven-design/app"

var (
	RepoClass   *RepositoryClass
	RepoStudent *RepositoryStudent
)

func Init(appCtx *app.AppCtx) {
	RepoClass = new(RepositoryClass)
	RepoClass.Init(appCtx)

	RepoStudent = new(RepositoryStudent)
	RepoStudent.Init(appCtx)
}
