package common

import "encoding/json"

// Literature ..
type Literature struct {
	BookIdent       string
	Year            int
	CourseTitle     string
	DepartmentTitle string
	Semester        int `json:"semester,omitempty"`
}

func (literature Literature) String() string {
	data, _ := json.MarshalIndent(literature, "", "\t")
	return string(data)
}

// GetLiteratureExample ..
func GetLiteratureExample() Literature {
	return Literature{
		BookIdent:       "КнигаДракона",
		CourseTitle:     "Компиляторы",
		Year:            2017,
		DepartmentTitle: "Прикладная математика и информатика",
		Semester:        6,
	}
}
