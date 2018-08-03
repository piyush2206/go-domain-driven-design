package admin

import (
	"fmt"

	"golang.org/x/net/context"
)

type (
	StudentService struct{}
)

func (ss *StudentService) Student(ctx context.Context, req *ReqStudent) (*ResStudent, error) {
	stud, _ := SampleStudents[req.StudentId]
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
