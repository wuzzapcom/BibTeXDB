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

// Departments ..
type Departments struct {
	DepartmentList []common.Department `json:"departments"`
}

// Lecturers ..
type Lecturers struct {
	LecturerList []common.Lecturer `json:"lecturers"`
}

// LiteratureLists ..
type LiteratureLists struct {
	StoredLists []common.LiteratureList `json:"lists"`
}

// Literature ..
type Literature struct {
	StoredLiterature []common.Literature `json:"literature"`
}

//Books ..
type Books struct {
	StoredBooks common.Items `json:"storedBooks"`
}

func (search Search) String() string {

	return search.Results.String()

}
