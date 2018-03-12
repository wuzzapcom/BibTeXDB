package database

import "time"

type Subject struct {
	Title        string
	LecturerID   string
	DepartmentID string
	Course       int
}

type LiteratureList struct {
	ISBNs []string
}

type Lecturer struct {
	ID           string     `bson:"_id,omitempty"`
	LecturerName string     `bson:"lecturerName"`
	DateOfBirth  *time.Time `bson:"dateOfBirth"`
}

type Department struct {
	ID         string `bson:"_id,omitempty"`
	Name       string
	FacilityID string `bson:"facilityID"`
}

type Facility struct {
	ID   string `bson:"_id,omitempty"`
	Name string `bson:"name"`
}
