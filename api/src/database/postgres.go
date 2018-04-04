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

// SQLExecutable интерфейс для унификации db и Tx структур
type SQLExecutable interface {
	Exec(string, ...interface{}) (sql.Result, error)
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
}

//Postgres ..
type Postgres struct {
	db          *sql.DB
	transaction *sql.Tx
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
	result, err := postgres.getSQLExecutable().Exec(
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

	rows, err := postgres.getSQLExecutable().Query("SELECT textbook_ident, textbook_title, " +
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
	result, err := postgres.getSQLExecutable().Exec("INSERT INTO schema.department(department_title, department_timestamp) VALUES ($1, $2)",
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
	rows, err := postgres.getSQLExecutable().Query("SELECT department_title FROM schema.department")
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

	result, err := postgres.getSQLExecutable().Exec("INSERT INTO schema.lecturer(lecturer_name, lecturer_date_of_birth, "+
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
	rows, err := postgres.getSQLExecutable().Query("SELECT lecturer_name, lecturer_date_of_birth, department_title FROM schema.lecturer l " +
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
	result, err := postgres.getSQLExecutable().Exec(
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
	courseID, err := postgres.FindIDOfCourseByCourseTitleAndDepartmentTitle(list.CourseTitle, list.DepartmentTitle, list.Semester)
	if err != nil {
		return err
	}
	fmt.Println(courseID)
	result, err := postgres.getSQLExecutable().Exec("INSERT INTO schema.literature_lists(literature_list_year, "+
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
	rows, err := postgres.getSQLExecutable().Query(`
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
	result, err := postgres.getSQLExecutable().Exec("INSERT INTO schema.course(course_title, course_lecturer_id, course_department_id, course_semester, course_timestamp) VALUES ($1, $2, $3, $4, $5)",
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
	rows, err := postgres.getSQLExecutable().Query("SELECT course_title, lecturer_name, lecturer_date_of_birth, department_title, course_semester FROM schema.lecturer " +
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
	row := postgres.getSQLExecutable().QueryRow("SELECT course_title, lecturer_name, lecturer_date_of_birth, department_title, course_semester FROM schema.lecturer "+
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
	literatureListID, err := postgres.FindIDOfLiteratureListByCourseTitleAndDepartmentTitleAndYear(
		literature.CourseTitle,
		literature.DepartmentTitle,
		literature.Semester,
		literature.Year,
	)
	if err != nil {
		return err
	}
	fmt.Println(literatureListID)
	result, err := postgres.getSQLExecutable().Exec("INSERT INTO schema.literature(literature_textbook_id, "+
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
	rows, err := postgres.getSQLExecutable().Query(`
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

// SelectBooksInList ..
func (postgres *Postgres) SelectBooksInList(list common.LiteratureList) (common.Items, error) {
	id, err := postgres.FindIDOfLiteratureListByCourseTitleAndDepartmentTitleAndYear(
		list.CourseTitle,
		list.DepartmentTitle,
		list.Semester,
		list.Year,
	)
	if err != nil {
		return nil, err
	}

	rows, err := postgres.getSQLExecutable().Query(`
		SELECT textbook_ident, textbook_title, textbook_author, textbook_publisher, textbook_year, textbook_isbn, textbook_url FROM 
		(SELECT * FROM schema.literature WHERE literature_literature_list_id = $1) l
				JOIN schema.textbook
					ON l.literature_textbook_id = textbook.textbook_id;`,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items common.Items
	for rows.Next() {
		item := common.Item{}
		err = rows.Scan(
			&item.Ident,
			&item.Title,
			&item.Author,
			&item.Publisher,
			&item.Year,
			&item.ISBN,
			&item.URL,
		)
		if err != nil {
			return nil, err
		}

		items.Append(item)
	}

	return items, err
}

// FindIDOfDepartmentWithName ..
func (postgres *Postgres) FindIDOfDepartmentWithName(name string) (int, error) {
	row := postgres.getSQLExecutable().QueryRow("SELECT department_id FROM schema.department WHERE department_title = $1", name)

	var id int
	err := row.Scan(&id)

	return id, err
}

// FindIDOfLecturerWithNameAndDateOfBirth ..
func (postgres *Postgres) FindIDOfLecturerWithNameAndDateOfBirth(name string, dateOfBirth time.Time) (int, error) {
	row := postgres.getSQLExecutable().QueryRow("SELECT lecturer_id FROM schema.lecturer WHERE lecturer_name=$1 AND lecturer_date_of_birth=$2",
		name,
		dateOfBirth,
	)

	var id int
	err := row.Scan(&id)

	return id, err
}

// FindIDOfCourseByCourseTitleAndDepartmentTitle ..
func (postgres *Postgres) FindIDOfCourseByCourseTitleAndDepartmentTitle(courseTitle string, departmentTitle string, semester int) (int, error) {
	departmentID, err := postgres.FindIDOfDepartmentWithName(departmentTitle)
	if err != nil {
		return 0, err
	}
	row := postgres.getSQLExecutable().QueryRow(`
		SELECT course_id 
			FROM schema.course 
				WHERE course_title=$1 AND 
					course_department_id=$2 AND
					course_semester=$3`,
		courseTitle,
		departmentID,
		semester,
	)

	var id int
	err = row.Scan(&id)

	return id, err
}

// FindIDOfTextbookByIdent ..
func (postgres *Postgres) FindIDOfTextbookByIdent(ident string) (int, error) {
	row := postgres.getSQLExecutable().QueryRow("SELECT textbook_id FROM schema.textbook WHERE textbook_ident = $1", ident)

	var id int
	err := row.Scan(&id)

	return id, err
}

// FindIDOfLiteratureListByCourseTitleAndDepartmentTitleAndYear ..
func (postgres *Postgres) FindIDOfLiteratureListByCourseTitleAndDepartmentTitleAndYear(
	courseTitle string,
	departmentTitle string,
	semester int,
	year int,
) (int, error) {

	courseID, err := postgres.FindIDOfCourseByCourseTitleAndDepartmentTitle(courseTitle, departmentTitle, semester)
	if err != nil {
		return 0, err
	}
	row := postgres.getSQLExecutable().QueryRow(`
		SELECT literature_list_id 
			FROM schema.literature_lists 
			WHERE literature_list_course_id = $1 AND literature_list_year = $2
		`, courseID, year)

	var id int
	err = row.Scan(&id)

	return id, err
}

// Migrate ..
func (postgres *Postgres) Migrate(migrate common.Migrate) error {

	fromLiteratureListID, err := postgres.FindIDOfLiteratureListByCourseTitleAndDepartmentTitleAndYear(
		migrate.CourseTitle,
		migrate.DepartmentTitle,
		migrate.Semester,
		migrate.From,
	)
	if err != nil {
		return err
	}

	toLiteratureListID, err := postgres.FindIDOfLiteratureListByCourseTitleAndDepartmentTitleAndYear(
		migrate.CourseTitle,
		migrate.DepartmentTitle,
		migrate.Semester,
		migrate.To,
	)
	if err != nil {
		return err
	}

	res, err := postgres.getSQLExecutable().Exec(`INSERT INTO schema.literature(literature_textbook_id, literature_literature_list_id, literature_timestamp)
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

// DeleteTextbook ..
func (postgres *Postgres) DeleteTextbook(ident string) error {
	var err error
	postgres.transaction, err = postgres.db.Begin()
	if err != nil {
		return err
	}

	row := postgres.getSQLExecutable().QueryRow(`
		UPDATE schema.textbook
			SET textbook_is_deleted = TRUE, textbook_timestamp = $1
			WHERE textbook_ident = $2
			RETURNING textbook_id
		`, getTimestamp(), ident)
	if err != nil {
		postgres.transaction.Rollback()
		postgres.transaction = nil
		return err
	}

	var textbookID int
	row.Scan(&textbookID)
	fmt.Println(textbookID)

	postgres.DeleteLiteratureWithTextbook(textbookID)

	err = postgres.transaction.Commit()
	if err != nil {
		postgres.transaction.Rollback()
		postgres.transaction = nil
		return err
	}
	postgres.transaction = nil

	return nil
}

// DeleteLiteratureWithTextbook ..
func (postgres *Postgres) DeleteLiteratureWithTextbook(textbookID int) error {

	_, err := postgres.getSQLExecutable().Exec(`
		UPDATE schema.literature
			SET literature_is_deleted = TRUE, literature_timestamp = $1
			WHERE literature_textbook_id = $2
		`, getTimestamp(), textbookID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteLiteratureWithList ..
func (postgres *Postgres) DeleteLiteratureWithList(literatureListID int) error {

	_, err := postgres.getSQLExecutable().Exec(`
		UPDATE schema.literature
			SET literature_is_deleted = TRUE, literature_timestamp = $1
			WHERE literature_literature_list_id = $2
		`, getTimestamp(), literatureListID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteLiterature ..
func (postgres *Postgres) DeleteLiterature(literature common.Literature) error {
	var err error
	postgres.transaction, err = postgres.db.Begin()
	if err != nil {
		return err
	}
	textbookID, err := postgres.FindIDOfTextbookByIdent(literature.BookIdent)
	if err != nil {
		postgres.transaction.Rollback()
		postgres.transaction = nil
		return err
	}
	literatureListID, err := postgres.FindIDOfLiteratureListByCourseTitleAndDepartmentTitleAndYear(
		literature.CourseTitle,
		literature.DepartmentTitle,
		literature.Semester,
		literature.Year,
	)
	if err != nil {
		postgres.transaction.Rollback()
		postgres.transaction = nil
		return err
	}

	_, err = postgres.getSQLExecutable().Exec(`
		UPDATE schema.literature
			SET literature_is_deleted = TRUE, literature_timestamp = $1
			WHERE literature_textbook_id = $2 AND literature_literature_list_id = $3
		`, getTimestamp(), textbookID, literatureListID)
	if err != nil {
		postgres.transaction.Rollback()
		postgres.transaction = nil
		return err
	}

	err = postgres.transaction.Commit()
	if err != nil {
		postgres.transaction.Rollback()
		postgres.transaction = nil
		return err
	}
	postgres.transaction = nil

	return nil
}

// DeleteLiteratureList ..
func (postgres *Postgres) DeleteLiteratureList(
	list common.LiteratureList,
) error {

	var err error
	postgres.transaction, err = postgres.db.Begin()
	if err != nil {
		return err
	}

	id, err := postgres.FindIDOfLiteratureListByCourseTitleAndDepartmentTitleAndYear(
		list.CourseTitle,
		list.DepartmentTitle,
		list.Semester,
		list.Year,
	)
	if err != nil {
		postgres.transaction.Rollback()
		postgres.transaction = nil
		return err
	}

	err = postgres.DeleteLiteratureListWithID(id)
	if err != nil {
		postgres.transaction.Rollback()
		postgres.transaction = nil
		return err
	}

	err = postgres.transaction.Commit()
	if err != nil {
		postgres.transaction.Rollback()
		postgres.transaction = nil
		return err
	}
	postgres.transaction = nil

	return nil
}

// DeleteLiteratureListWithID ..
func (postgres *Postgres) DeleteLiteratureListWithID(
	id int,
) error {

	_, err := postgres.getSQLExecutable().Exec(`
		UPDATE schema.literature_lists
			SET literature_list_is_deleted = TRUE, literature_list_timestamp = $1
			WHERE literature_list_id = $2
		`, getTimestamp(), id)
	if err != nil {
		fmt.Println(55)
		return err
	}

	err = postgres.DeleteLiteratureWithList(id)
	if err != nil {
		fmt.Println(56)
		return err
	}

	return nil
}

// DeleteCourse ..
func (postgres *Postgres) DeleteCourse(title string, department string, semester int) error {

	var err error
	postgres.transaction, err = postgres.db.Begin()
	if err != nil {
		return err
	}

	courseID, err := postgres.FindIDOfCourseByCourseTitleAndDepartmentTitle(title, department, semester)
	if err != nil {
		postgres.transaction.Rollback()
		postgres.transaction = nil
		return err
	}

	_, err = postgres.getSQLExecutable().Exec(`
		UPDATE schema.course
			SET course_is_deleted = TRUE, course_timestamp = $1
			WHERE course_id = $2
		`, getTimestamp(), courseID)
	if err != nil {
		postgres.transaction.Rollback()
		postgres.transaction = nil
		return err
	}

	rows, err := postgres.getSQLExecutable().Query(`
		SELECT literature_list_id FROM schema.literature_lists
			WHERE literature_list_course_id=$1
		`, courseID)
	if err != nil {
		postgres.transaction.Rollback()
		postgres.transaction = nil
		return err
	}

	literatureListIDs := make([]int, 0)
	for rows.Next() {
		var literatureListID int
		err = rows.Scan(&literatureListID)
		if err != nil {
			postgres.transaction.Rollback()
			postgres.transaction = nil
			return err
		}
		literatureListIDs = append(literatureListIDs, literatureListID)
	}
	rows.Close()

	for _, literatureListID := range literatureListIDs {
		err = postgres.DeleteLiteratureListWithID(literatureListID)
		if err != nil {
			fmt.Println(5)
			postgres.transaction.Rollback()
			postgres.transaction = nil
			return err
		}
	}

	err = postgres.transaction.Commit()
	if err != nil {
		postgres.transaction.Rollback()
		postgres.transaction = nil
		return err
	}
	postgres.transaction = nil

	return nil
}

func (postgres *Postgres) getSQLExecutable() SQLExecutable {
	if postgres.transaction != nil {
		return postgres.transaction
	}
	return postgres.db
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

func getTimestamp() int32 {
	return int32(time.Now().Unix())
}
