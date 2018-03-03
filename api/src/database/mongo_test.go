package database_test

import (
	"testing"

	"wuzzapcom/Coursework/pkg/bibtex"
	"wuzzapcom/Coursework/pkg/database"
)

func TestTextbookInsert(t *testing.T) {
	mongo := database.Mongo{}

	mongo.Connect()
	defer mongo.Disconnect()

	mongo.DropTextbooks()

	err := mongo.InsertTextbooks(bibtex.GetRandomItems())
	if err != nil {
		t.Fatal(err)
	}

	textbooks, err := mongo.FindAllTextbooks()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(textbooks)

	if len(textbooks) == 0 {
		t.Fatal("Empty textbooks")
	}

	for i, textbook := range textbooks {
		if textbook != bibtex.GetRandomItems()[i] {
			t.Errorf("Not equal. Got result %s", textbooks.String())
		}
	}

}

func TestFindAllTextbooks(t *testing.T) {
	mongo := database.Mongo{}
	mongo.Connect()
	defer mongo.Disconnect()

	t.Log(mongo.FindAllTextbooks())
}
