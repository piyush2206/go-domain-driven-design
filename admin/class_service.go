package admin

import (
	"fmt"

	"golang.org/x/net/context"
)

type (
	// ClassService implements class service interface of grpc
	ClassService struct {
		err error
	}
)

// Class is the handler of grpc endpoint 'Class'
// It returns class info of the requested classId
func (cs *ClassService) Class(ctx context.Context, req *ReqClass) (*ResClass, error) {
	cls, _ := SampleClasses[req.ClassId]

	resCls := &ResClass{
		Class: &ClassInfo{
			Standard: cls.Standard,
			Division: cls.Division,
		},
	}

	for _, stud := range SampleStudents {
		if stud.Class.Id != req.ClassId {
			continue
		}
		resCls.Students = append(resCls.Students, &StudentInfo{
			Name: fmt.Sprintf("%s %s", stud.FName, stud.LName),
		})
	}

	return resCls, nil
}

// Create is the handler of grpc endpoint 'Create'
// It creates a class record in database with the passed request payload
func (cs *ClassService) Create(ctx context.Context, req *ReqClassCreate) (*ResSuccess, error) {
	res := new(ResSuccess)
	class := new(Class)

	class.New(req)
	defer cs.formResponse(res)

	cs.err = RepoClass.Create(class)

	return res, nil
}

func (cs *ClassService) formResponse(res *ResSuccess) {
	if cs.err != nil {
		res.Success = false
		res.Msg = "Failed to add class"
	} else {
		res.Success = true
		res.Msg = "Successfully created a new class"
	}
}
