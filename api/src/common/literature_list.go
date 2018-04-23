package common

import "encoding/json"

// LiteratureList ..
type LiteratureList struct {
	ID              int `json:"id,omitempty"`
	Year            int
	DepartmentTitle string
	CourseTitle     string
	Semester        int
}

func (list LiteratureList) String() string {
	data, _ := json.MarshalIndent(list, "", "\t")
	return string(data)
}

// GetLiteratureListExample ..
func GetLiteratureListExample() LiteratureList {
	return LiteratureList{
		Year:            2017,
		DepartmentTitle: "Прикладная математика и информатика",
		CourseTitle:     "Компиляторы",
		Semester:        6,
	}
}
