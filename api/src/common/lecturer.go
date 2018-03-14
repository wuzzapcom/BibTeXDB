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
