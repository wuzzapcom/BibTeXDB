package database_test

import (
	"fmt"
	"testing"
	"time"
	"wuzzapcom/Coursework/api/src/common"
	"wuzzapcom/Coursework/api/src/database"
)

var testPort = 32768

func TestPostgres_Connect(t *testing.T) {
	database.Configuration.Port = testPort
	postgres := database.Postgres{}
	err := postgres.Connect()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}
	postgres.Disconnect()
}

func TestPostgres_InsertTextbook(t *testing.T) {
	database.Configuration.Port = testPort
	postgres := database.Postgres{}
	err := postgres.Connect()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}
	defer postgres.Disconnect()

	err = postgres.InsertTextbook(
		common.Item{
			Ident:     "FlowersForAlgernon",
			Title:     "Flowers for Algernon",
			Author:    "D.D. Keyes",
			Publisher: "Houghton Mifflin Harcourt",
			Year:      "2007",
			Language:  "English",
			ISBN:      "9780547539638",
			URL:       "https://www.googleapis.com/books/v1/volumes/_oG_iTxP1pIC",
		})
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}
}

func TestPostgres_InsertTextbooks(t *testing.T) {
	database.Configuration.Port = testPort
	postgres := database.Postgres{}
	err := postgres.Connect()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}
	defer postgres.Disconnect()

	err = postgres.InsertTextbooks(common.GetRandomItems())
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}
}

func TestPostgres_SelectTextbooks(t *testing.T) {
	database.Configuration.Port = testPort
	postgres := database.Postgres{}
	err := postgres.Connect()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}
	defer postgres.Disconnect()

	result, err := postgres.SelectTextbooks()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}

	fmt.Println(result)
}

func TestPostgres_InsertDepartment(t *testing.T) {
	database.Configuration.Port = testPort
	postgres := database.Postgres{}
	err := postgres.Connect()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}
	defer postgres.Disconnect()

	err = postgres.InsertDepartment(common.Department{
		Title: "Прикладная математика и информатика",
	})
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}
}

func TestPostgres_SelectDepartments(t *testing.T) {
	database.Configuration.Port = testPort
	postgres := database.Postgres{}
	err := postgres.Connect()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}
	defer postgres.Disconnect()

	result, err := postgres.SelectDepartments()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}

	fmt.Println(result)
}

func TestPostgres_FindIDOfDepartmentWithName(t *testing.T) {
	database.Configuration.Port = testPort
	postgres := database.Postgres{}
	err := postgres.Connect()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}
	defer postgres.Disconnect()

	result, err := postgres.FindIDOfDepartmentWithName("Прикладная математика и информатика")
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}
	fmt.Println(result)
}

func TestPostgres_InsertLecturer(t *testing.T) {
	database.Configuration.Port = testPort
	postgres := database.Postgres{}
	err := postgres.Connect()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}
	defer postgres.Disconnect()

	dateOfBirth, err := time.Parse("2006-Jan-02", "2013-Feb-03")
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}

	err = postgres.InsertLecturer(common.Lecturer{
		Name:        "Коновалов Александр Владимирович",
		DateOfBirth: dateOfBirth,
		Department:  "Прикладная математика и информатика",
	})
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}
}

func TestPostgres_SelectLecturers(t *testing.T) {
	database.Configuration.Port = testPort
	postgres := database.Postgres{}
	err := postgres.Connect()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}
	defer postgres.Disconnect()

	result, err := postgres.SelectLecturers()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}

	fmt.Println(result)
}

func TestPostgres_DeleteDepartment(t *testing.T) {
	database.Configuration.Port = 32768
	postgres := database.Postgres{}
	err := postgres.Connect()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}
	defer postgres.Disconnect()

	err = postgres.DeleteDepartment(5)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}
}

func TestPostgres_InsertLiteratureList(t *testing.T) {
	database.Configuration.Port = testPort
	postgres := database.Postgres{}
	err := postgres.Connect()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}
	defer postgres.Disconnect()

	err = postgres.InsertLiteratureList(common.LiteratureList{
		Year:            2017,
		DepartmentTitle: "ИУ9",
		CourseTitle:     "Компиляторы",
	})
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}
}

func TestPostgres_SelectLiteratureLists(t *testing.T) {
	database.Configuration.Port = testPort
	postgres := database.Postgres{}
	err := postgres.Connect()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}
	defer postgres.Disconnect()

	result, err := postgres.SelectLiteratureList()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}

	fmt.Println(result)
}

func TestPostgres_InsertCourse(t *testing.T) {
	database.Configuration.Port = testPort
	postgres := database.Postgres{}
	err := postgres.Connect()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}
	defer postgres.Disconnect()

	dateOfBirth := common.HumanizedTime{}
	dateOfBirth.Time, err = time.Parse(common.TimeFormat, "2013-Feb-03")
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}

	err = postgres.InsertCourse(common.Course{
		Title:       "Объектно-функциональное программирование",
		Lecturer:    "Скоробогатов Сергей Юрьевич",
		DateOfBirth: dateOfBirth,
		Department:  "Прикладная математика и информатика",
		Semester:    6,
	})
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}
}

func TestPostgres_SelectCourses(t *testing.T) {
	database.Configuration.Port = testPort
	postgres := database.Postgres{}
	err := postgres.Connect()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}
	defer postgres.Disconnect()

	result, err := postgres.SelectCourses()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}

	fmt.Println(result)
}

func TestPostgres_InsertLiterature(t *testing.T) {
	database.Configuration.Port = testPort
	postgres := database.Postgres{}
	err := postgres.Connect()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}
	defer postgres.Disconnect()

	err = postgres.InsertLiterature(common.Literature{
		BookIdent:       "FlowersForAlgernon",
		Year:            2017,
		CourseTitle:     "Компиляторы",
		DepartmentTitle: "ИУ9",
	})
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}
}

func TestPostgres_SelectLiterature(t *testing.T) {
	database.Configuration.Port = testPort
	postgres := database.Postgres{}
	err := postgres.Connect()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}
	defer postgres.Disconnect()

	result, err := postgres.SelectLiterature()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}

	fmt.Println(result)
}
