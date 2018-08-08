package admin

import (
	"errors"
	"fmt"

	"github.com/piyush2206/go-domain-driven-design/dep"
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
		DB dep.IDb
	}
)

// NewClass initiates a class entity object
func NewClass(req interface{}) (class *Class, err error) {
	switch reqT := req.(type) {
	case *ReqClassCreate:
		class = &Class{reqT.Id, reqT.Standard, reqT.Division, nil}
		return class, nil
	default:
		return nil, errors.New("param type not supported")
	}
}

// NewClassRepository initiates class repository with required external dependencies
// and returns a new class repository
func NewClassRepository(DB dep.IDb) (repo *RepositoryClass) {
	return &RepositoryClass{DB: DB}
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
		dep.TblClass, class.Id, class.Standard, class.Division)

	if _, err := rc.DB.Query(qryInsert); err != nil {
		return err
	}
	return nil
}
