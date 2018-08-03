package admin

import (
	"errors"
	"fmt"

	"github.com/piyush2206/go-domain-driven-design/app"
)

type (
	// Class entity
	Class struct {
		Id       string
		Standard string
		Division string
		Subjects []*Subject
	}

	// RepositoryClass repository
	RepositoryClass struct {
		appCtx *app.AppCtx
	}
)

// New initiates a class entity object
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

// Init initiates class repository with required external dependencies
func (rc *RepositoryClass) Init(appCtx *app.AppCtx) {
	rc.appCtx = appCtx
}

// Class returns class entity object of the requested classId
func (rc *RepositoryClass) Class(classId string) (cls *Class) {
	cls, _ = SampleClasses[classId]
	return cls
}

// ClassSubjects returns all the subjects info of the requested classId
func (rc *RepositoryClass) ClassSubjects(classId string) (subjects []*Subject) {
	cls, ok := SampleClasses[classId]
	if !ok {
		return nil
	}
	return cls.Subjects
}

// Create creates a new class record in the database
func (rc *RepositoryClass) Create(class *Class) (err error) {
	qryInsert := fmt.Sprintf("INSERT INTO %s VALUES ('%s', '%s', '%s')",
		app.TblClass, class.Id, class.Standard, class.Division)

	if _, err := rc.appCtx.DB.Query(qryInsert); err != nil {
		return err
	}
	return nil
}
