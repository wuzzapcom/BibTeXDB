package restful

import (
	"wuzzapcom/Coursework/api/src/common"
)

type Error struct {
	Message string `json:"errorMessage"`
}

type Search struct {
	Results common.Items `json:"results"`
}

type Success struct {
	Message string `json:"successMessage"`
}

type Courses struct {
	CoursesList []common.Course `json:"courses"`
}

type Books struct {
	StoredBooks common.Items `json:"storedBooks"`
}

func (search Search) String() string {

	return search.Results.String()

}