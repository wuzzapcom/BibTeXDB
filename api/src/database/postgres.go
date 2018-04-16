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
		return common.CreateError(errors.New("Wrong TCP port"))
	}

	var port = 32770
	if Configuration.Port != 0 {
		port = Configuration.Port
	}

	connStr := fmt.Sprintf("user=wuzzapcom port=%d dbname=bibtex_literature sslmode=disable", port)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return common.CreateError(err)
	}

	postgres.db = db
	err = postgres.db.Ping()
	if err != nil {
		return common.CreateError(err)
	}

	return nil
}

//InsertTextbook ..
func (postgres *Postgres) InsertTextbook(item common.Item) error {
	result, err := postgres.getSQLExecutable().Exec(
		`INSERT INTO schema.textbook(textbook_ident, textbook_title,
			textbook_author, textbook_publisher, textbook_year, textbook_isbn,
			textbook_url, textbook_timestamp) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
			ON CONFLICT (textbook_isbn) DO UPDATE SET
				textbook_ident=EXCLUDED.textbook_ident,
				textbook_title=EXCLUDED.textbook_title,
				textbook_author=EXCLUDED.textbook_author,
				textbook_publisher=EXCLUDED.textbook_publisher,
				textbook_year=EXCLUDED.textbook_year,
				textbook_url=EXCLUDED.textbook_url,
				textbook_timestamp=EXCLUDED.textbook_timestamp,
				textbook_is_deleted=FALSE;
			`,
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
		return common.CreateError(err)
	}

	fmt.Println(result.RowsAffected())
	return nil
}

// InsertTextbooks ..
func (postgres *Postgres) InsertTextbooks(items common.Items) error {
	for _, item := range items {
		err := postgres.InsertTextbook(item)
		if err != nil {
			return common.CreateError(err)
		}
	}
	return nil
}

// SelectTextbooks ..
func (postgres *Postgres) SelectTextbooks() (common.Items, error) {

	rows, err := postgres.getSQLExecutable().Query(`
		SELECT textbook_ident, textbook_title, textbook_author, 
		textbook_publisher, textbook_year, textbook_isbn, textbook_url 
			FROM schema.textbook
				WHERE textbook_is_deleted = FALSE
			`)
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
	result, err := postgres.getSQLExecutable().Exec(`
		INSERT INTO schema.department(department_title, department_timestamp) VALUES 
		($1, $2)
		ON CONFLICT (department_title) DO UPDATE SET
			department_title=EXCLUDED.department_title,
			department_is_deleted=FALSE,
			department_timestamp=EXCLUDED.department_timestamp`,
		department.Title,
		int32(time.Now().Unix()),
	)
	if err != nil {
		return common.CreateError(err)
	}

	fmt.Println(result.RowsAffected())
	return nil
}

// SelectDepartments ..
func (postgres *Postgres) SelectDepartments() ([]common.Department, error) {
	rows, err := postgres.getSQLExecutable().Query(`
		SELECT department_title 
			FROM schema.department
				WHERE department_is_deleted = FALSE
		`)
	if err != nil {
		return nil, common.CreateError(err)
	}
	defer rows.Close()

	var deparments []common.Department
	for rows.Next() {
		deparment := new(common.Department)
		err = rows.Scan(&deparment.Title)
		if err != nil {
			return nil, common.CreateError(err)
		}
		deparments = append(deparments, *deparment)
	}

	return deparments, nil
}

// InsertLecturer ..
func (postgres *Postgres) InsertLecturer(lecturer common.Lecturer) error {
	id, err := postgres.FindIDOfDepartmentWithName(lecturer.Department)
	if err != nil {
		return common.CreateError(err)
	}

	result, err := postgres.getSQLExecutable().Exec(`
		INSERT INTO schema.lecturer(lecturer_name, lecturer_date_of_birth,
			lecturer_department_id, lecturer_timestamp) 
			VALUES ($1, $2, $3, $4)
			ON CONFLICT(lecturer_name, lecturer_date_of_birth) DO UPDATE SET
			lecturer_department_id=EXCLUDED.lecturer_department_id,
			lecturer_timestamp=EXCLUDED.lecturer_timestamp,
			lecturer_is_deleted=FALSE;
		`,
		lecturer.Name,
		lecturer.DateOfBirth.Time,
		id,
		int32(time.Now().Unix()),
	)
	if err != nil {
		return common.CreateError(err)
	}

	fmt.Println(result.RowsAffected())
	return nil
}

// SelectLecturers ..
func (postgres *Postgres) SelectLecturers() ([]common.Lecturer, error) {
	rows, err := postgres.getSQLExecutable().Query(`
		SELECT lecturer_name, lecturer_date_of_birth, department_title 
		FROM schema.lecturer l 
			JOIN schema.department d 
				ON l.lecturer_department_id = d.department_id
			WHERE lecturer_is_deleted = FALSE
		`)
	if err != nil {
		return nil, common.CreateError(err)
	}
	defer rows.Close()

	var lecturers []common.Lecturer
	for rows.Next() {
		lecturer := new(common.Lecturer)
		err = rows.Scan(&lecturer.Name, &lecturer.DateOfBirth.Time, &lecturer.Department)
		if err != nil {
			return nil, common.CreateError(err)
		}
		lecturers = append(lecturers, *lecturer)
	}

	return lecturers, nil
}

// InsertLiteratureList ..
func (postgres *Postgres) InsertLiteratureList(list common.LiteratureList) error {
	courseID, err := postgres.FindIDOfCourseByCourseTitleAndDepartmentTitle(list.CourseTitle, list.DepartmentTitle, list.Semester)
	if err != nil {
		return common.CreateError(err)
	}
	fmt.Println(courseID)
	result, err := postgres.getSQLExecutable().Exec("INSERT INTO schema.literature_lists(literature_list_year, "+
		"literature_list_course_id, literature_list_timestamp) VALUES ($1, $2, $3)",
		list.Year,
		courseID,
		int32(time.Now().Unix()),
	)
	if err != nil {
		return common.CreateError(err)
	}

	fmt.Println(result.RowsAffected())
	return nil
}

// SelectLiteratureList ..
func (postgres *Postgres) SelectLiteratureList() ([]common.LiteratureList, error) {
	rows, err := postgres.getSQLExecutable().Query(`
	SELECT literature_list_id, literature_list_year, course_title, department_title, course_semester 
		FROM schema.department d 
			JOIN (schema.literature_lists l 
				JOIN schema.course c 
					ON l.literature_list_course_id = c.course_id) j 
				ON d.department_id = j.course_department_id
			WHERE literature_list_is_deleted = FALSE
		`)
	if err != nil {
		return nil, common.CreateError(err)
	}
	defer rows.Close()

	var lists []common.LiteratureList
	for rows.Next() {
		list := new(common.LiteratureList)
		err = rows.Scan(&list.ID, &list.Year, &list.CourseTitle, &list.DepartmentTitle, &list.Semester)
		if err != nil {
			return nil, common.CreateError(err)
		}
		lists = append(lists, *list)
	}

	return lists, nil
}

//InsertCourse ..
func (postgres *Postgres) InsertCourse(course common.Course) error {
	departmentID, err := postgres.FindIDOfDepartmentWithName(course.Department)
	if err != nil {
		return common.CreateError(err)
	}
	fmt.Println(departmentID)
	lecturerID, err := postgres.FindIDOfLecturerWithNameAndDateOfBirth(course.Lecturer, course.DateOfBirth.Time)
	if err != nil {
		return common.CreateError(err)
	}
	fmt.Println(lecturerID)
	result, err := postgres.getSQLExecutable().Exec(`I
		NSERT INTO schema.course(course_title, course_lecturer_id, 
			course_department_id, course_semester, course_timestamp) 
			VALUES ($1, $2, $3, $4, $5)
			ON CONFLICT(course_title, course_department_id, course_semester) DO UPDATE SET
			course_lecturer_id=EXCLUDED.course_lecturer_id,
			course_timestamp=EXCLUDED.course_timestamp,
			course_is_deleted=FALSE;
			`,
		course.Title,
		lecturerID,
		departmentID,
		course.Semester,
		int32(time.Now().Unix()),
	)
	if err != nil {
		return common.CreateError(err)
	}

	fmt.Println(result.RowsAffected())
	return nil
}

//SelectCourses ..
func (postgres *Postgres) SelectCourses() ([]common.Course, error) {
	rows, err := postgres.getSQLExecutable().Query(`
		SELECT course_title, lecturer_name, lecturer_date_of_birth, 
		department_title, course_semester 
			FROM schema.lecturer
				JOIN (schema.course 
					JOIN schema.department 
						ON course_department_id = department_id) j 
					ON lecturer_id = j.course_lecturer_id
			WHERE course_is_deleted = FALSE`)
	if err != nil {
		return nil, common.CreateError(err)
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
			return nil, common.CreateError(err)
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
	if err != nil {
		return *course, common.CreateError(err)
	}
	return *course, nil
}

// InsertLiterature ..
func (postgres *Postgres) InsertLiterature(literature common.Literature) error {
	textbookID, err := postgres.FindIDOfTextbookByIdent(literature.BookIdent)
	if err != nil {
		return common.CreateError(err)
	}
	fmt.Println(textbookID)
	literatureListID, err := postgres.FindIDOfLiteratureListByCourseTitleAndDepartmentTitleAndYear(
		literature.CourseTitle,
		literature.DepartmentTitle,
		literature.Semester,
		literature.Year,
	)
	if err != nil {
		return common.CreateError(err)
	}
	fmt.Println(literatureListID)
	result, err := postgres.getSQLExecutable().Exec("INSERT INTO schema.literature(literature_textbook_id, "+
		"literature_literature_list_id, literature_timestamp) VALUES ($1, $2, $3)",
		textbookID,
		literatureListID,
		int32(time.Now().Unix()),
	)
	if err != nil {
		return common.CreateError(err)
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
				ON j.literature_literature_list_id = l.literature_list_id
			WHERE literature_is_deleted = FALSE`)
	if err != nil {
		return nil, common.CreateError(err)
	}
	defer rows.Close()

	var literatures []common.Literature
	for rows.Next() {
		literature := new(common.Literature)
		var courseID int
		err = rows.Scan(&literature.BookIdent, &courseID, &literature.Year)
		if err != nil {
			return nil, common.CreateError(err)
		}
		course, err := postgres.SelectCourse(courseID)
		if err != nil {
			return nil, common.CreateError(err)
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
		return nil, common.CreateError(err)
	}

	rows, err := postgres.getSQLExecutable().Query(`
		SELECT textbook_ident, textbook_title, textbook_author, textbook_publisher, textbook_year, textbook_isbn, textbook_url FROM 
		(SELECT * FROM schema.literature 
			WHERE literature_literature_list_id = $1 AND literature_is_deleted = FALSE) l
				JOIN schema.textbook
					ON l.literature_textbook_id = textbook.textbook_id
				WHERE textbook_is_deleted = FALSE`,
		id,
	)
	if err != nil {
		return nil, common.CreateError(err)
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
			return nil, common.CreateError(err)
		}

		items.Append(item)
	}

	return items, nil
}

// FindIDOfDepartmentWithName ..
func (postgres *Postgres) FindIDOfDepartmentWithName(name string) (int, error) {
	row := postgres.getSQLExecutable().QueryRow("SELECT department_id FROM schema.department WHERE department_title = $1", name)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return id, common.CreateError(err)
	}

	return id, nil
}

// FindIDOfLecturerWithNameAndDateOfBirth ..
func (postgres *Postgres) FindIDOfLecturerWithNameAndDateOfBirth(name string, dateOfBirth time.Time) (int, error) {
	row := postgres.getSQLExecutable().QueryRow("SELECT lecturer_id FROM schema.lecturer WHERE lecturer_name=$1 AND lecturer_date_of_birth=$2",
		name,
		dateOfBirth,
	)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return id, common.CreateError(err)
	}

	return id, nil
}

// FindIDOfCourseByCourseTitleAndDepartmentTitle ..
func (postgres *Postgres) FindIDOfCourseByCourseTitleAndDepartmentTitle(
	courseTitle string,
	departmentTitle string,
	semester int,
) (int, error) {

	departmentID, err := postgres.FindIDOfDepartmentWithName(departmentTitle)
	if err != nil {
		return 0, common.CreateError(err)
	}
	row := postgres.getSQLExecutable().QueryRow(`
		SELECT course_id 
			FROM schema.course 
				WHERE course_title=$1 AND 
					course_department_id=$2 AND
					course_semester=$3 AND
					course_is_deleted = FALSE`,
		courseTitle,
		departmentID,
		semester,
	)

	var id int
	err = row.Scan(&id)
	if err != nil {
		return id, common.CreateError(err)
	}

	return id, nil
}

// FindIDOfTextbookByIdent ..
func (postgres *Postgres) FindIDOfTextbookByIdent(ident string) (int, error) {
	row := postgres.getSQLExecutable().QueryRow(`
		SELECT textbook_id 
			FROM schema.textbook 
				WHERE textbook_ident = $1 AND textbook_is_deleted = FALSE
		`, ident)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return id, common.CreateError(err)
	}

	return id, nil
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
		return 0, common.CreateError(err)
	}
	row := postgres.getSQLExecutable().QueryRow(`
		SELECT literature_list_id 
			FROM schema.literature_lists 
			WHERE literature_list_course_id = $1 AND 
				literature_list_year = $2 AND 
				literature_list_is_deleted = FALSE
		`, courseID, year)

	var id int
	err = row.Scan(&id)
	if err != nil {
		return id, common.CreateError(err)
	}

	return id, nil
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
		return common.CreateError(err)
	}

	toLiteratureListID, err := postgres.FindIDOfLiteratureListByCourseTitleAndDepartmentTitleAndYear(
		migrate.CourseTitle,
		migrate.DepartmentTitle,
		migrate.Semester,
		migrate.To,
	)
	if err != nil {
		return common.CreateError(err)
	}

	res, err := postgres.getSQLExecutable().Exec(`
		INSERT INTO schema.literature(literature_textbook_id, literature_literature_list_id, literature_timestamp)
			  SELECT literature_textbook_id, $1, $3
				  FROM schema.literature 
				  	WHERE literature_literature_list_id = $2;`,
		toLiteratureListID,
		fromLiteratureListID,
		getTimestamp(),
	)
	if err != nil {
		return common.CreateError(err)
	}

	fmt.Println(res.RowsAffected())

	return nil
}

// DeleteTextbook ..
func (postgres *Postgres) DeleteTextbook(ident string) error {
	var err error
	postgres.transaction, err = postgres.db.Begin()
	if err != nil {
		return common.CreateError(err)
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
		return common.CreateError(err)
	}

	var textbookID int
	err = row.Scan(&textbookID)
	if err != nil {
		postgres.transaction.Rollback()
		postgres.transaction = nil
		return common.CreateError(err)
	}
	fmt.Println(textbookID)

	err = postgres.DeleteLiteratureWithTextbook(textbookID)
	if err != nil {
		postgres.transaction.Rollback()
		postgres.transaction = nil
		return common.CreateError(err)
	}

	err = postgres.transaction.Commit()
	if err != nil {
		postgres.transaction.Rollback()
		postgres.transaction = nil
		return common.CreateError(err)
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
		return common.CreateError(err)
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
		return common.CreateError(err)
	}

	return nil
}

// DeleteLiterature ..
func (postgres *Postgres) DeleteLiterature(literature common.Literature) error {
	var err error
	postgres.transaction, err = postgres.db.Begin()
	if err != nil {
		return common.CreateError(err)
	}
	textbookID, err := postgres.FindIDOfTextbookByIdent(literature.BookIdent)
	if err != nil {
		postgres.transaction.Rollback()
		postgres.transaction = nil
		return common.CreateError(err)
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
		return common.CreateError(err)
	}

	_, err = postgres.getSQLExecutable().Exec(`
		UPDATE schema.literature
			SET literature_is_deleted = TRUE, literature_timestamp = $1
			WHERE literature_textbook_id = $2 AND literature_literature_list_id = $3
		`, getTimestamp(), textbookID, literatureListID)
	if err != nil {
		postgres.transaction.Rollback()
		postgres.transaction = nil
		return common.CreateError(err)
	}

	err = postgres.transaction.Commit()
	if err != nil {
		postgres.transaction.Rollback()
		postgres.transaction = nil
		return common.CreateError(err)
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
		return common.CreateError(err)
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
		return common.CreateError(err)
	}

	err = postgres.DeleteLiteratureListWithID(id)
	if err != nil {
		postgres.transaction.Rollback()
		postgres.transaction = nil
		return common.CreateError(err)
	}

	err = postgres.transaction.Commit()
	if err != nil {
		postgres.transaction.Rollback()
		postgres.transaction = nil
		return common.CreateError(err)
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
		return common.CreateError(err)
	}

	err = postgres.DeleteLiteratureWithList(id)
	if err != nil {
		return common.CreateError(err)
	}

	return nil
}

// DeleteCourseWithID ..
func (postgres *Postgres) DeleteCourseWithID(courseID int) error {

	_, err := postgres.getSQLExecutable().Exec(`
		UPDATE schema.course
			SET course_is_deleted = TRUE, course_timestamp = $1
			WHERE course_id = $2
		`, getTimestamp(), courseID)
	if err != nil {
		return common.CreateError(err)
	}

	rows, err := postgres.getSQLExecutable().Query(`
		SELECT literature_list_id FROM schema.literature_lists
			WHERE literature_list_course_id=$1
		`, courseID)
	if err != nil {
		return common.CreateError(err)
	}

	literatureListIDs := make([]int, 0)
	for rows.Next() {
		var literatureListID int
		err = rows.Scan(&literatureListID)
		if err != nil {
			return common.CreateError(err)
		}
		literatureListIDs = append(literatureListIDs, literatureListID)
	}
	rows.Close()

	for _, literatureListID := range literatureListIDs {
		err = postgres.DeleteLiteratureListWithID(literatureListID)
		if err != nil {
			return common.CreateError(err)
		}
	}

	return nil
}

// DeleteCourse ..
func (postgres *Postgres) DeleteCourse(title string, department string, semester int) error {

	var err error
	postgres.transaction, err = postgres.db.Begin()
	if err != nil {
		return common.CreateError(err)
	}

	courseID, err := postgres.FindIDOfCourseByCourseTitleAndDepartmentTitle(title, department, semester)
	if err != nil {
		postgres.transaction.Rollback()
		postgres.transaction = nil
		return common.CreateError(err)
	}

	err = postgres.DeleteCourseWithID(courseID)
	if err != nil {
		postgres.transaction.Rollback()
		postgres.transaction = nil
		return common.CreateError(err)
	}

	err = postgres.transaction.Commit()
	if err != nil {
		postgres.transaction.Rollback()
		postgres.transaction = nil
		return common.CreateError(err)
	}
	postgres.transaction = nil

	return nil
}

// DeleteLecturerWithID ..
func (postgres *Postgres) DeleteLecturerWithID(lecturerID int) error {

	_, err := postgres.getSQLExecutable().Exec(`
		UPDATE schema.lecturer
			SET lecturer_is_deleted = TRUE, lecturer_timestamp = $1
			WHERE lecturer_id = $2
		`, getTimestamp(), lecturerID)
	if err != nil {
		return common.CreateError(err)
	}

	rows, err := postgres.getSQLExecutable().Query(`
		SELECT course_id FROM schema.course
			WHERE course_lecturer_id=$1
		`, lecturerID)
	if err != nil {
		return common.CreateError(err)
	}

	courseIDs := make([]int, 0)
	for rows.Next() {
		var courseID int
		err = rows.Scan(&courseID)
		if err != nil {
			return common.CreateError(err)
		}
		courseIDs = append(courseIDs, courseID)
	}
	rows.Close()

	for _, courseID := range courseIDs {
		err = postgres.DeleteCourseWithID(courseID)
		if err != nil {
			return common.CreateError(err)
		}
	}
	return nil
}

// DeleteLecturer ..
func (postgres *Postgres) DeleteLecturer(name string, dateOfBirth time.Time) error {
	var err error
	postgres.transaction, err = postgres.db.Begin()
	if err != nil {
		return common.CreateError(err)
	}

	lecturerID, err := postgres.FindIDOfLecturerWithNameAndDateOfBirth(name, dateOfBirth)
	if err != nil {
		postgres.transaction.Rollback()
		postgres.transaction = nil
		return common.CreateError(err)
	}

	err = postgres.DeleteLecturerWithID(lecturerID)
	if err != nil {
		postgres.transaction.Rollback()
		postgres.transaction = nil
		return common.CreateError(err)
	}

	err = postgres.transaction.Commit()
	if err != nil {
		postgres.transaction.Rollback()
		postgres.transaction = nil
		return common.CreateError(err)
	}
	postgres.transaction = nil

	return nil
}

// DeleteDepartmentWithID ..
func (postgres *Postgres) DeleteDepartmentWithID(departmentID int) error {
	_, err := postgres.getSQLExecutable().Exec(`
		UPDATE schema.department
			SET department_is_deleted = TRUE, department_timestamp = $1
			WHERE department_id = $2
		`, getTimestamp(), departmentID)
	if err != nil {
		return common.CreateError(err)
	}
	//delete courses
	rows, err := postgres.getSQLExecutable().Query(`
		SELECT course_id FROM schema.course
			WHERE course_department_id=$1
		`, departmentID)
	if err != nil {
		return common.CreateError(err)
	}

	courseIDs := make([]int, 0)
	for rows.Next() {
		var courseID int
		err = rows.Scan(&courseID)
		if err != nil {
			return common.CreateError(err)
		}
		courseIDs = append(courseIDs, courseID)
	}
	rows.Close()

	for _, courseID := range courseIDs {
		err = postgres.DeleteCourseWithID(courseID)
		if err != nil {
			return common.CreateError(err)
		}
	}
	//delete lecturers
	rows, err = postgres.getSQLExecutable().Query(`
		SELECT lecturer_id FROM schema.lecturer
			WHERE lecturer_department_id=$1
		`, departmentID)
	if err != nil {
		return common.CreateError(err)
	}

	lecturerIDs := make([]int, 0)
	for rows.Next() {
		var lecturerID int
		err = rows.Scan(&lecturerID)
		if err != nil {
			return common.CreateError(err)
		}
		lecturerIDs = append(lecturerIDs, lecturerID)
	}
	rows.Close()

	for _, lecturerID := range lecturerIDs {
		err = postgres.DeleteLecturerWithID(lecturerID)
		if err != nil {
			return common.CreateError(err)
		}
	}
	return nil
}

// DeleteDepartment ..
func (postgres *Postgres) DeleteDepartment(title string) error {
	var err error
	postgres.transaction, err = postgres.db.Begin()
	if err != nil {
		return common.CreateError(err)
	}

	lecturerID, err := postgres.FindIDOfDepartmentWithName(title)
	if err != nil {
		postgres.transaction.Rollback()
		postgres.transaction = nil
		return common.CreateError(err)
	}

	err = postgres.DeleteDepartmentWithID(lecturerID)
	if err != nil {
		postgres.transaction.Rollback()
		postgres.transaction = nil
		return common.CreateError(err)
	}

	err = postgres.transaction.Commit()
	if err != nil {
		postgres.transaction.Rollback()
		postgres.transaction = nil
		return common.CreateError(err)
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
