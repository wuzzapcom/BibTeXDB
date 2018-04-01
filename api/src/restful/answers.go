package restful

import (
	"wuzzapcom/Coursework/api/src/common"
)

//Error ..
type Error struct {
	Message string `json:"errorMessage"`
}

//Search ..
type Search struct {
	Results common.Items `json:"results"`
}

//Success ..
type Success struct {
	Message string `json:"successMessage"`
}

//Courses ..
type Courses struct {
	CoursesList []common.Course `json:"courses"`
}

func (courses Courses) String() (result string) {
	for _, course := range courses.CoursesList {
		result += course.String()
		result += "\n"
	}
	return
}

// Departments ..
type Departments struct {
	DepartmentList []common.Department `json:"departments"`
}

func (departments Departments) String() (result string) {
	for _, department := range departments.DepartmentList {
		result += department.String()
		result += "\n"
	}
	return
}

// Lecturers ..
type Lecturers struct {
	LecturerList []common.Lecturer `json:"lecturers"`
}

func (lecturers Lecturers) String() (result string) {
	for _, lecturer := range lecturers.LecturerList {
		result += lecturer.String()
		result += "\n"
	}
	return
}

// LiteratureLists ..
type LiteratureLists struct {
	StoredLists []common.LiteratureList `json:"lists"`
}

func (literatureLists LiteratureLists) String() (result string) {
	for _, literatureList := range literatureLists.StoredLists {
		result += literatureList.String()
		result += "\n"
	}
	return
}

// Literature ..
type Literature struct {
	StoredLiterature []common.Literature `json:"literature"`
}

func (literature Literature) String() (result string) {
	for _, l := range literature.StoredLiterature {
		result += l.String()
		result += "\n"
	}
	return
}

//Books ..
type Books struct {
	StoredBooks common.Items `json:"storedBooks"`
}

func (books Books) String() (result string) {
	for _, book := range books.StoredBooks {
		result += book.String()
		result += "\n"
	}
	return
}
func (search Search) String() string {

	return search.Results.String()

}
