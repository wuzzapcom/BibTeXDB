package database

import "time"

//Subject ..
type Subject struct {
	Title        string
	LecturerID   string
	DepartmentID string
	Course       int
}

//LiteratureList ..
type LiteratureList struct {
	ISBNs []string
}

//Lecturer ..
type Lecturer struct {
	ID           string     `bson:"_id,omitempty"`
	LecturerName string     `bson:"lecturerName"`
	DateOfBirth  *time.Time `bson:"dateOfBirth"`
}

//Department ..
type Department struct {
	ID         string `bson:"_id,omitempty"`
	Name       string
	FacilityID string `bson:"facilityID"`
}

//Facility ..
type Facility struct {
	ID   string `bson:"_id,omitempty"`
	Name string `bson:"name"`
}
