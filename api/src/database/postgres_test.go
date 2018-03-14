package database_test

import (
	"fmt"
	"testing"
	"wuzzapcom/Coursework/api/src/common"
	"wuzzapcom/Coursework/api/src/database"
)

func TestPostgres_Connect(t *testing.T) {
	postgres := database.Postgres{}
	err := postgres.Connect()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}
	postgres.Disconnect()
}

func TestPostgres_InsertTextbook(t *testing.T) {
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
