// report package is responsible for students progress card reports

package report

import (
	"math"

	"github.com/piyush2206/go-domain-driven-design/admin"
)

type (
	subjectScore struct {
		score float64
	}

	reportSubject struct {
		subject   *admin.Subject
		subjScore *subjectScore
	}

	report struct {
		student        *admin.Student
		class          *admin.Class
		reportSubjects []reportSubject
	}
)

func (ss *subjectScore) round() (score *subjectScore) {
	score = &subjectScore{score: math.Ceil(ss.score*100) / 100}
	return score
}

func (r *report) New(cls *admin.Class, stud *admin.Student) {
	r.student = stud
	r.class = cls

	for i, subj := range cls.Subjects {
		reportSubj := reportSubject{
			subject:   subj,
			subjScore: SampleReports[stud.Id].reportSubjects[i].subjScore,
		}
		r.reportSubjects = append(r.reportSubjects, reportSubj)
	}
}

func (r *report) Compute() (grandtotal, totalScore, percentScored float64) {
	for _, subjs := range r.reportSubjects {
		roundScore := subjs.subjScore.round()
		grandtotal += subjs.subject.TotalMarks
		totalScore += roundScore.score
	}
	percentScored = totalScore / grandtotal * 100

	return grandtotal, totalScore, percentScored
}
