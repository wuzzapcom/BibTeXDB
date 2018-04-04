package common

import "encoding/json"

//Migrate defines JSON for migration literatureList from year to year
type Migrate struct {
	CourseTitle     string
	DepartmentTitle string
	Semester        int
	From            int
	To              int
}

func (migrate Migrate) String() string {
	data, _ := json.MarshalIndent(migrate, "", "\t")
	return string(data)
}

// GetMigrateExample ..
func GetMigrateExample() Migrate {
	return Migrate{
		CourseTitle:     "Компиляторы",
		DepartmentTitle: "Прикладная математика и информатика",
		From:            2017,
		To:              2018,
	}
}
