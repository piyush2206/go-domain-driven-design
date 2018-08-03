package report

import "github.com/piyush2206/go-domain-driven-design/admin"

var SampleReports = map[string]*report{
	"S1": &report{
		student: admin.SampleStudents["S1"],
		class:   admin.SampleStudents["S1"].Class,
		reportSubjects: []reportSubject{
			{
				subject:   admin.SampleStudents["S1"].Class.Subjects[0],
				subjScore: &subjectScore{67},
			},
			{
				subject:   admin.SampleStudents["S1"].Class.Subjects[1],
				subjScore: &subjectScore{88},
			},
		},
	},
	"S2": &report{
		student: admin.SampleStudents["S2"],
		class:   admin.SampleStudents["S2"].Class,
		reportSubjects: []reportSubject{
			{
				subject:   admin.SampleStudents["S2"].Class.Subjects[0],
				subjScore: &subjectScore{99},
			},
			{
				subject:   admin.SampleStudents["S2"].Class.Subjects[1],
				subjScore: &subjectScore{69},
			},
		},
	},
	"S3": &report{
		student: admin.SampleStudents["S3"],
		class:   admin.SampleStudents["S3"].Class,
		reportSubjects: []reportSubject{
			{
				subject:   admin.SampleStudents["S3"].Class.Subjects[0],
				subjScore: &subjectScore{65},
			},
			{
				subject:   admin.SampleStudents["S3"].Class.Subjects[1],
				subjScore: &subjectScore{77},
			},
		},
	},
	"S4": &report{
		student: admin.SampleStudents["S4"],
		class:   admin.SampleStudents["S4"].Class,
		reportSubjects: []reportSubject{
			{
				subject:   admin.SampleStudents["S4"].Class.Subjects[0],
				subjScore: &subjectScore{90},
			},
			{
				subject:   admin.SampleStudents["S4"].Class.Subjects[1],
				subjScore: &subjectScore{59},
			},
			{
				subject:   admin.SampleStudents["S4"].Class.Subjects[2],
				subjScore: &subjectScore{91},
			},
		},
	},
	"S5": &report{
		student: admin.SampleStudents["S5"],
		class:   admin.SampleStudents["S5"].Class,
		reportSubjects: []reportSubject{
			{
				subject:   admin.SampleStudents["S5"].Class.Subjects[0],
				subjScore: &subjectScore{90},
			},
			{
				subject:   admin.SampleStudents["S5"].Class.Subjects[1],
				subjScore: &subjectScore{60},
			},
			{
				subject:   admin.SampleStudents["S5"].Class.Subjects[2],
				subjScore: &subjectScore{68},
			},
		},
	},
}
