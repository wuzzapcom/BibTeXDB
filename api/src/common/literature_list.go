package common

import "encoding/json"

// LiteratureList ..
type LiteratureList struct {
	ID              int `json:"omitempty"`
	Year            int
	DepartmentTitle string
	CourseTitle     string
}

func (list LiteratureList) String() string {
	data, _ := json.MarshalIndent(list, "", "\t")
	return string(data)
}
