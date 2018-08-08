package admin

import (
	"fmt"

	"golang.org/x/net/context"
)

type (
	StudentService struct {
		repoStudent *RepositoryStudent
	}
)

// NewStudentService returns a new empty object of StudentService
func NewStudentService(repoStudent *RepositoryStudent) StudentServer {
	return &StudentService{repoStudent: repoStudent}
}

// Student rpc of Student service returns requested student data from student ID
func (ss *StudentService) Student(ctx context.Context, req *ReqStudent) (*ResStudent, error) {
	stud := ss.repoStudent.Student(req.StudentId)

	resStud := &ResStudent{
		Student: &StudentInfo{
			Name: fmt.Sprintf("%s %s", stud.FName, stud.LName),
		},
		Class: &ClassInfo{
			Standard: stud.Class.Standard,
			Division: stud.Class.Division,
		},
	}
	return resStud, nil
}
