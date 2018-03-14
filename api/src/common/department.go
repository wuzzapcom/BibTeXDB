package common

import (
	"encoding/json"
)

//Department ..
type Department struct {
	Title string `json:"title"`
}

func (department Department) String() string {
	data, _ := json.MarshalIndent(department, "", "\t")
	return string(data)
}

// GetDepartmentExample ..
func GetDepartmentExample() Department {
	return Department{
		Title: "Прикладная математика и информатика",
	}
}
