package common

import (
	"encoding/json"
	"time"
)

//Course ..
type Course struct {
	Title       string    `json:"title"`
	Department  string    `json:"department"`
	Lecturer    string    `json:"lecturer"`
	DateOfBirth time.Time `json:"date_of_birth,omitempty"`
	Semester    int       `json:"semester"`
}

func (course Course) String() string {
	data, _ := json.MarshalIndent(course, "", "\t")
	return string(data)
}

//GetCourseExample ..
func GetCourseExample() Course {
	dateOfBirth, _ := time.Parse("2006-Jan-02", "2013-Feb-03")

	return Course{
		Title:       "Конструирование компиляторов",
		Lecturer:    "Коновалов Александр Владимирович",
		Department:  "ИУ9",
		DateOfBirth: dateOfBirth,
		Semester:    6,
	}
}
