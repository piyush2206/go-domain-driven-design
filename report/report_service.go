package report

import (
	"errors"
	"fmt"

	"github.com/piyush2206/go-domain-driven-design/admin"
	"golang.org/x/net/context"
)

type (
	ReportService struct {
		repoClass   *admin.RepositoryClass
		repoStudent *admin.RepositoryStudent
	}
)

// NewReportService returns a new ReportService object pointer
func NewReportService(
	repoClass *admin.RepositoryClass,
	repoStudent *admin.RepositoryStudent,
) ReportServer {
	return &ReportService{
		repoClass:   repoClass,
		repoStudent: repoStudent,
	}
}

func (rs *ReportService) Report(ctx context.Context, req *ReqReportStudent) (*ResReportStudent, error) {
	res := new(ResReportStudent)
	res.Report = make([]*StudentReport, 1)

	cls := rs.repoClass.Class(req.ClassId)
	if cls == nil {
		return nil, errors.New("class not found")
	}

	stud := rs.repoStudent.Student(req.StudentId)
	if stud == nil {
		return nil, errors.New("student not found")
	}

	reports := new(report)

	reports.New(cls, stud)

	formReportResponse(res, reports)

	return res, nil
}

func formReportResponse(res *ResReportStudent, reports *report) {
	res.Class = &admin.ClassInfo{
		Standard: reports.class.Standard,
		Division: reports.class.Division,
	}
	res.Report[0] = new(StudentReport)

	res.Report[0].StudentInfo = &admin.StudentInfo{
		Name: fmt.Sprintf("%s %s", reports.student.FName, reports.student.LName),
	}
	res.Report[0].GrandTotal, res.Report[0].GrandTotalMarks, res.Report[0].Percent = reports.Compute()
}
