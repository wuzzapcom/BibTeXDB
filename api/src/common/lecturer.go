package common

import (
	"encoding/json"
	"time"
)

// Lecturer ..
type Lecturer struct {
	Name        string
	DateOfBirth HumanizedTime
	Department  string
}

func (lecturer Lecturer) String() string {
	data, _ := json.MarshalIndent(lecturer, "", "\t")
	return string(data)
}

// GetLecturerExample ..
func GetLecturerExample() Lecturer {
	dateOfBirth := HumanizedTime{}
	dateOfBirth.Time, _ = time.Parse(TimeFormat, "2013-Feb-03")

	return Lecturer{
		Name:        "Коновалов Александр Владимирович",
		DateOfBirth: dateOfBirth,
		Department:  "Прикладная математика и информатика",
	}
}
