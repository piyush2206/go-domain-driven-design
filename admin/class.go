package admin

import (
	"errors"
	"fmt"

	"github.com/piyush2206/go-domain-driven-design/app"
)

type (
	Class struct {
		Id       string
		Standard string
		Division string
		Subjects []*Subject
	}

	RepositoryClass struct {
		appCtx *app.AppCtx
	}
)

func (c *Class) New(req interface{}) (err error) {
	switch reqT := req.(type) {
	case *ReqClassCreate:
		c = &Class{reqT.Id, reqT.Standard, reqT.Division, nil}
		c.Id = reqT.Id
		return nil
	default:
		return errors.New("type not supported")
	}
}

func (rc *RepositoryClass) Init(appCtx *app.AppCtx) {
	rc.appCtx = appCtx
}

func (rc *RepositoryClass) Class(classId string) (cls *Class) {
	cls, _ = SampleClasses[classId]
	return cls
}

func (rc *RepositoryClass) ClassSubjects(classId string) (subjects []*Subject) {
	cls, ok := SampleClasses[classId]
	if !ok {
		return nil
	}
	return cls.Subjects
}

func (rc *RepositoryClass) Create(class *Class) (err error) {
	qryInsert := fmt.Sprintf("INSERT INTO %s VALUES ('%s', '%s', '%s')",
		app.TblClass, class.Id, class.Standard, class.Division)

	if _, err := rc.appCtx.DB.Query(qryInsert); err != nil {
		return err
	}
	return nil
}
