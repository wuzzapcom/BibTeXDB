package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
	"wuzzapcom/Coursework/api/src/common"

	//SQL driver import
	_ "github.com/lib/pq"
)

//Postgres ..
type Postgres struct {
	db *sql.DB
}

// PostgresConfiguration ..
type PostgresConfiguration struct {
	Port int
}

// Configuration ..
var Configuration = PostgresConfiguration{}

//Connect ..
func (postgres *Postgres) Connect() error {

	if Configuration.Port < 0 {
		return errors.New("Wrong TCP port")
	}

	var port = 32770
	if Configuration.Port != 0 {
		port = Configuration.Port
	}

	connStr := fmt.Sprintf("user=wuzzapcom port=%d dbname=bibtex_literature sslmode=disable", port)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	postgres.db = db
	err = postgres.db.Ping()
	if err != nil {
		return err
	}

	return nil
}

//InsertTextbook ..
func (postgres *Postgres) InsertTextbook(item common.Item) error {
	result, err := postgres.db.Exec(
		"INSERT INTO schema.textbook(textbook_ident, textbook_title, "+
			"textbook_author, textbook_publisher, textbook_year, textbook_isbn, "+
			"textbook_url, textbook_timestamp) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		item.Ident,
		item.Title,
		item.Author,
		item.Publisher,
		item.Year,
		item.ISBN,
		item.URL,
		int32(time.Now().Unix()),
	)
	if err != nil {
		return err
	}

	fmt.Println(result.RowsAffected())
	return nil
}

// InsertTextbooks ..
func (postgres *Postgres) InsertTextbooks(items common.Items) error {
	for _, item := range items {
		err := postgres.InsertTextbook(item)
		if err != nil {
			return err
		}
	}
	return nil
}

// SelectTextbooks ..
func (postgres *Postgres) SelectTextbooks() (common.Items, error) {

	rows, err := postgres.db.Query("SELECT textbook_ident, textbook_title, " +
		"textbook_author, textbook_publisher, textbook_year, textbook_isbn, textbook_url FROM schema.textbook")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items common.Items
	for rows.Next() {
		item := new(common.Item)
		err = rows.Scan(&item.Ident, &item.Title, &item.Author, &item.Publisher, &item.Year, &item.ISBN, &item.URL)
		if err != nil {
			return nil, err
		}
		items.Append(item)
	}

	return items, err

}

// InsertDepartment ..
func (postgres *Postgres) InsertDepartment(department common.Department) error {
	result, err := postgres.db.Exec("INSERT INTO schema.department(department_title, department_timestamp) VALUES ($1, $2)",
		department.Title,
		int32(time.Now().Unix()),
	)
	if err != nil {
		return err
	}

	fmt.Println(result.RowsAffected())
	return nil
}

// SelectDepartments ..
func (postgres *Postgres) SelectDepartments() ([]common.Department, error) {
	rows, err := postgres.db.Query("SELECT department_title FROM schema.department")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var deparments []common.Department
	for rows.Next() {
		deparment := new(common.Department)
		err = rows.Scan(&deparment.Title)
		if err != nil {
			return nil, err
		}
		deparments = append(deparments, *deparment)
	}

	return deparments, nil
}

// InsertLecturer ..
func (postgres *Postgres) InsertLecturer(lecturer common.Lecturer) error {
	id, err := postgres.FindIDOfDepartmentWithName(lecturer.Department)
	if err != nil {
		return err
	}

	result, err := postgres.db.Exec("INSERT INTO schema.lecturer(lecturer_name, lecturer_date_of_birth, "+
		"lecturer_department_id, lecturer_timestamp) VALUES ($1, $2, $3, $4)",
		lecturer.Name,
		lecturer.DateOfBirth.Time,
		id,
		int32(time.Now().Unix()),
	)
	if err != nil {
		return err
	}

	fmt.Println(result.RowsAffected())
	return nil
}

// SelectLecturers ..
func (postgres *Postgres) SelectLecturers() ([]common.Lecturer, error) {
	rows, err := postgres.db.Query("SELECT lecturer_name, lecturer_date_of_birth, department_title FROM schema.lecturer l " +
		"JOIN schema.department d ON l.lecturer_department_id = d.department_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lecturers []common.Lecturer
	for rows.Next() {
		lecturer := new(common.Lecturer)
		err = rows.Scan(&lecturer.Name, &lecturer.DateOfBirth.Time, &lecturer.Department)
		if err != nil {
			return nil, err
		}
		lecturers = append(lecturers, *lecturer)
	}

	return lecturers, nil
}

// DeleteDepartment ..
func (postgres *Postgres) DeleteDepartment(id int) error {
	result, err := postgres.db.Exec(
		"DELETE FROM schema.department WHERE department_id = $1",
		id,
	)
	if err != nil {
		return err
	}

	fmt.Println(result.RowsAffected())
	return nil
}

// InsertLiteratureList ..
func (postgres *Postgres) InsertLiteratureList(list common.LiteratureList) error {
	courseID, err := postgres.FindIDOfCourseByCourseTitleAndDepartmentTitle(list.CourseTitle, list.DepartmentTitle)
	if err != nil {
		return err
	}
	fmt.Println(courseID)
	result, err := postgres.db.Exec("INSERT INTO schema.literature_lists(literature_list_year, "+
		"literature_list_course_id, literature_list_timestamp) VALUES ($1, $2, $3)",
		list.Year,
		courseID,
		int32(time.Now().Unix()),
	)
	if err != nil {
		return err
	}

	fmt.Println(result.RowsAffected())
	return nil
}

// SelectLiteratureList ..
func (postgres *Postgres) SelectLiteratureList() ([]common.LiteratureList, error) {
	rows, err := postgres.db.Query(`
	SELECT literature_list_id, literature_list_year, course_title, department_title FROM schema.department d 
		JOIN (schema.literature_lists l 
			JOIN schema.course c 
				ON l.literature_list_course_id = c.course_id) j 
			ON d.department_id = j.course_department_id;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lists []common.LiteratureList
	for rows.Next() {
		list := new(common.LiteratureList)
		err = rows.Scan(&list.ID, &list.Year, &list.CourseTitle, &list.DepartmentTitle)
		if err != nil {
			return nil, err
		}
		lists = append(lists, *list)
	}

	return lists, nil
}

//InsertCourse ..
func (postgres *Postgres) InsertCourse(course common.Course) error {
	departmentID, err := postgres.FindIDOfDepartmentWithName(course.Department)
	if err != nil {
		return err
	}
	fmt.Println(departmentID)
	lecturerID, err := postgres.FindIDOfLecturerWithNameAndDateOfBirth(course.Lecturer, course.DateOfBirth.Time)
	if err != nil {
		return err
	}
	fmt.Println(lecturerID)
	result, err := postgres.db.Exec("INSERT INTO schema.course(course_title, course_lecturer_id, course_department_id, course_semester, course_timestamp) VALUES ($1, $2, $3, $4, $5)",
		course.Title,
		lecturerID,
		departmentID,
		course.Semester,
		int32(time.Now().Unix()),
	)
	if err != nil {
		return err
	}

	fmt.Println(result.RowsAffected())
	return nil
}

//SelectCourses ..
func (postgres *Postgres) SelectCourses() ([]common.Course, error) {
	rows, err := postgres.db.Query("SELECT course_title, lecturer_name, lecturer_date_of_birth, department_title, course_semester FROM schema.lecturer " +
		"JOIN (schema.course " +
		"JOIN schema.department " +
		"ON course_department_id = department_id) j " +
		"ON lecturer_id = j.course_lecturer_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []common.Course
	for rows.Next() {
		course := new(common.Course)
		err = rows.Scan(
			&course.Title,
			&course.Lecturer,
			&course.DateOfBirth.Time,
			&course.Department,
			&course.Semester,
		)
		if err != nil {
			return nil, err
		}
		courses = append(courses, *course)
	}

	return courses, nil
}

// SelectCourse ..
func (postgres *Postgres) SelectCourse(courseID int) (common.Course, error) {
	row := postgres.db.QueryRow("SELECT course_title, lecturer_name, lecturer_date_of_birth, department_title, course_semester FROM schema.lecturer "+
		"JOIN (schema.course "+
		"JOIN schema.department "+
		"ON course_department_id = department_id) j "+
		"ON lecturer_id = j.course_lecturer_id WHERE course_id = $1", courseID)

	course := new(common.Course)
	err := row.Scan(
		&course.Title,
		&course.Lecturer,
		&course.DateOfBirth.Time,
		&course.Department,
		&course.Semester,
	)
	return *course, err
}

// InsertLiterature ..
func (postgres *Postgres) InsertLiterature(literature common.Literature) error {
	textbookID, err := postgres.FindIDOfTextbookByIdent(literature.BookIdent)
	if err != nil {
		return err
	}
	fmt.Println(textbookID)
	literatureListID, err := postgres.FindIDOfLiteratureListByCourseTitleAndDepartmentTitleAndYear(literature.CourseTitle, literature.DepartmentTitle, literature.Year)
	if err != nil {
		return err
	}
	fmt.Println(literatureListID)
	result, err := postgres.db.Exec("INSERT INTO schema.literature(literature_textbook_id, "+
		"literature_literature_list_id, literature_timestamp) VALUES ($1, $2, $3)",
		textbookID,
		literatureListID,
		int32(time.Now().Unix()),
	)
	if err != nil {
		return err
	}

	fmt.Println(result.RowsAffected())
	return nil
}

// SelectLiterature ..
func (postgres *Postgres) SelectLiterature() ([]common.Literature, error) {
	rows, err := postgres.db.Query(`
		SELECT textbook_ident, literature_list_course_id, literature_list_year FROM schema.literature_lists l
			JOIN (schema.literature
				JOIN schema.textbook
					ON literature.literature_textbook_id = textbook.textbook_id) j
				ON j.literature_literature_list_id = l.literature_list_id;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var literatures []common.Literature
	for rows.Next() {
		literature := new(common.Literature)
		var courseID int
		err = rows.Scan(&literature.BookIdent, &courseID, &literature.Year)
		if err != nil {
			return nil, err
		}
		course, err := postgres.SelectCourse(courseID)
		if err != nil {
			return nil, err
		}
		literature.CourseTitle = course.Title
		literature.DepartmentTitle = course.Department
		literatures = append(literatures, *literature)
	}

	return literatures, nil
}

// FindIDOfDepartmentWithName ..
func (postgres *Postgres) FindIDOfDepartmentWithName(name string) (int, error) {
	row := postgres.db.QueryRow("SELECT department_id FROM schema.department WHERE department_title = $1", name)

	var id int
	err := row.Scan(&id)

	return id, err
}

// FindIDOfLecturerWithNameAndDateOfBirth ..
func (postgres *Postgres) FindIDOfLecturerWithNameAndDateOfBirth(name string, dateOfBirth time.Time) (int, error) {
	row := postgres.db.QueryRow("SELECT lecturer_id FROM schema.lecturer WHERE lecturer_name=$1 AND lecturer_date_of_birth=$2",
		name,
		dateOfBirth,
	)

	var id int
	err := row.Scan(&id)

	return id, err
}

// FindIDOfCourseByCourseTitleAndDepartmentTitle ..
func (postgres *Postgres) FindIDOfCourseByCourseTitleAndDepartmentTitle(courseTitle string, departmentTitle string) (int, error) {
	departmentID, err := postgres.FindIDOfDepartmentWithName(departmentTitle)
	if err != nil {
		return 0, err
	}
	row := postgres.db.QueryRow("SELECT course_id FROM schema.course WHERE course_title=$1 AND course_department_id=$2",
		courseTitle,
		departmentID,
	)

	var id int
	err = row.Scan(&id)

	return id, err
}

// FindIDOfTextbookByIdent ..
func (postgres *Postgres) FindIDOfTextbookByIdent(ident string) (int, error) {
	row := postgres.db.QueryRow("SELECT textbook_id FROM schema.textbook WHERE textbook_ident = $1", ident)

	var id int
	err := row.Scan(&id)

	return id, err
}

// FindIDOfLiteratureListByCourseTitleAndDepartmentTitleAndYear ..
func (postgres *Postgres) FindIDOfLiteratureListByCourseTitleAndDepartmentTitleAndYear(
	courseTitle string,
	departmentTitle string,
	year int,
) (int, error) {

	courseID, err := postgres.FindIDOfCourseByCourseTitleAndDepartmentTitle(courseTitle, departmentTitle)
	if err != nil {
		return 0, err
	}
	row := postgres.db.QueryRow("SELECT literature_list_id FROM schema.literature_lists WHERE literature_list_course_id = $1 AND literature_list_year = $2", courseID, year)

	var id int
	err = row.Scan(&id)

	return id, err
}

// Migrate ..
func (postgres *Postgres) Migrate(migrate common.Migrate) error {

	fromLiteratureListID, err := postgres.FindIDOfLiteratureListByCourseTitleAndDepartmentTitleAndYear(
		migrate.CourseTitle,
		migrate.DepartmentTitle,
		migrate.From,
	)
	if err != nil {
		return err
	}

	toLiteratureListID, err := postgres.FindIDOfLiteratureListByCourseTitleAndDepartmentTitleAndYear(
		migrate.CourseTitle,
		migrate.DepartmentTitle,
		migrate.To,
	)
	if err != nil {
		return err
	}

	res, err := postgres.db.Exec(`INSERT INTO schema.literature(literature_textbook_id, literature_literature_list_id, literature_timestamp)
  SELECT literature_textbook_id, $1, literature_timestamp FROM schema.literature WHERE literature_literature_list_id = $2;`,
		toLiteratureListID,
		fromLiteratureListID,
	)
	if err != nil {
		return err
	}

	fmt.Println(res.RowsAffected())

	return nil
}

//FindAllTextbooks ..
func (postgres *Postgres) FindAllTextbooks() (common.Items, error) {
	return nil, errors.New("Unimplemented")
}

//GetAllCourses ..
func (postgres *Postgres) GetAllCourses() ([]common.Course, error) {
	return nil, errors.New("Unimplemented")
}

// Disconnect ..
func (postgres *Postgres) Disconnect() {
	postgres.db.Close()
}
