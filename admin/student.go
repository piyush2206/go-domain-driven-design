package admin

import (
	"github.com/piyush2206/go-domain-driven-design/dep"
)

type (
	Student struct {
		Id    string
		FName string
		LName string
		Class *Class
	}

	RepositoryStudent struct {
		DB dep.IDb
	}
)

// NewStudentRepository initiates class repository with required external dependencies
// and returns a new class repository
func NewStudentRepository(DB dep.IDb) (repo *RepositoryStudent) {
	return &RepositoryStudent{DB: DB}
}

func (rs *RepositoryStudent) Student(studId string) (stud *Student) {
	stud, _ = SampleStudents[studId]
	return stud
}
