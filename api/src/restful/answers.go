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

//Books ..
type Books struct {
	StoredBooks common.Items `json:"storedBooks"`
}

func (search Search) String() string {

	return search.Results.String()

}
