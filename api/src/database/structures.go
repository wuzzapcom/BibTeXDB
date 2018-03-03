package database

type Subject struct {
	Title        string
	LecturerID   string
	DepartmentID string
	Cource       int
}

type LiteratureList struct {
	ISBNs []string
}
