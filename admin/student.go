package admin

import "github.com/piyush2206/go-domain-driven-design/app"

type (
	Student struct {
		Id    string
		FName string
		LName string
		Class *Class
	}

	RepositoryStudent struct {
		appCtx *app.AppCtx
	}
)

func (rs *RepositoryStudent) Init(appCtx *app.AppCtx) {
	rs.appCtx = appCtx
}

func (rs *RepositoryStudent) Student(studId string) (stud *Student) {
	stud, _ = SampleStudents[studId]
	return stud
}
