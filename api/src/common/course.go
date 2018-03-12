package common

import "encoding/json"

//Course ..
type Course struct {
	Title      string `json:"title"`
	Lecturer   string `json:"lecturer"`
	Department string `json:"department"`
	Semester   int    `json:"semester"`
}

func (course Course) String() string {
	data, _ := json.MarshalIndent(course, "", "\t")
	return string(data)
}

//GetCourseExample ..
func GetCourseExample() Course {
	return Course{
		Title:      "Конструирование компиляторов",
		Lecturer:   "Коновалов Александр Владимирович",
		Department: "ИУ9",
		Semester:   6,
	}
}