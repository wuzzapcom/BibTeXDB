package database

import (
	"database/sql"
	"errors"
	"fmt"
	"wuzzapcom/Coursework/api/src/common"

	//SQL driver import
	_ "github.com/lib/pq"
)

//Postgres ..
type Postgres struct {
	db *sql.DB
}

//Connect ..
func (postgres *Postgres) Connect() error {
	connStr := "user=wuzzapcom port=32768 dbname=postgres sslmode=disable"
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
		"INSERT INTO schema.textbook(ident, title, author, publisher, year, isbn, url) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		item.Ident,
		item.Title,
		item.Author,
		item.Publisher,
		item.Year,
		item.ISBN,
		item.URL,
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

	rows, err := postgres.db.Query("SELECT ident, title, author, publisher, year, isbn, url FROM schema.textbook")
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
	result, err := postgres.db.Exec("INSERT INTO schema.department(title) VALUES ($1)",
		department.Title,
	)
	if err != nil {
		return err
	}

	fmt.Println(result.RowsAffected())
	return nil
}

// SelectDepartments ..
func (postgres *Postgres) SelectDepartments() ([]common.Department, error) {
	rows, err := postgres.db.Query("SELECT title FROM schema.department")
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

//FindAllTextbooks ..
func (postgres *Postgres) FindAllTextbooks() (common.Items, error) {
	return nil, errors.New("Unimplemented")
}

//InsertCourse ..
func (postgres *Postgres) InsertCourse(course common.Course) error {
	return errors.New("Unimplemented")
}

//GetAllCourses ..
func (postgres *Postgres) GetAllCourses() ([]common.Course, error) {
	return nil, errors.New("Unimplemented")
}

// Disconnect ..
func (postgres *Postgres) Disconnect() {
	postgres.db.Close()
}
