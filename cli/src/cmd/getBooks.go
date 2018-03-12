package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"wuzzapcom/Coursework/api/src/restful"

	"github.com/spf13/cobra"
)

var getBooksCommand = &cobra.Command{
	Use: "getBooks",
	Run: getBooks,
}

func getBooks(cmd *cobra.Command, args []string) {

	url := "http://localhost:8080/getBooks"

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	if response.StatusCode != 200 {
		handleError(data)
		return
	}
	var books restful.Books
	err = json.Unmarshal(data, &books)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	for _, book := range books.StoredBooks {
		fmt.Println(book.JSONString())
	}


}

func init() {
}
