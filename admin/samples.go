package admin

// SampleClasses is sample class data for mocking database
var SampleClasses = map[string]*Class{
	"C1": &Class{
		Id:       "C1",
		Standard: "1st",
		Division: "A",
		Subjects: []*Subject{
			SampleSubjects["English"],
			SampleSubjects["Maths"],
		},
	},
	"C2": &Class{
		Id:       "C2",
		Standard: "2nd",
		Division: "A",
		Subjects: []*Subject{
			SampleSubjects["English"],
			SampleSubjects["Maths"],
			SampleSubjects["Science"],
		},
	},
}

// SampleSubjects is sample subjects data for mocking database
var SampleSubjects = map[string]*Subject{
	"English": &Subject{
		Name:       "English",
		TotalMarks: 100,
	},
	"Maths": &Subject{
		Name:       "Maths",
		TotalMarks: 100,
	},
	"Science": &Subject{
		Name:       "Science",
		TotalMarks: 100,
	},
}

// SampleStudents is sample students data for mocking database
var SampleStudents = map[string]*Student{
	"S1": &Student{
		Id:    "S1",
		FName: "Piyush",
		LName: "Vishwakarma",
		Class: SampleClasses["C1"],
	},
	"S2": &Student{
		Id:    "S2",
		FName: "Abhinav",
		LName: "Gaur",
		Class: SampleClasses["C1"],
	},
	"S3": &Student{
		Id:    "S3",
		FName: "Niraj",
		LName: "Bhattad",
		Class: SampleClasses["C1"],
	},
	"S4": &Student{
		Id:    "S4",
		FName: "Suresh",
		LName: "Rajput",
		Class: SampleClasses["C2"],
	},
	"S5": &Student{
		Id:    "S5",
		FName: "Disha",
		LName: "Duggal",
		Class: SampleClasses["C2"],
	},
}
