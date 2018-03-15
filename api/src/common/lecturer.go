package common

import (
	"encoding/json"
	"time"
)

// Lecturer ..
type Lecturer struct {
	Name        string
	DateOfBirth time.Time
	Department  string
}

func (lecturer Lecturer) String() string {
	data, _ := json.MarshalIndent(lecturer, "", "\t")
	return string(data)
}

// GetLecturerExample ..
func GetLecturerExample() Lecturer {
	dateOfBirth, _ := time.Parse("2006-Jan-02", "2013-Feb-03")

	return Lecturer{
		Name:        "Коновалов Александр Владимирович",
		DateOfBirth: dateOfBirth,
		Department:  "Прикладная математика и информатика",
	}
}
