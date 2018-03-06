package cmd

import (
	"fmt"
	"wuzzapcom/Coursework/api/src/restful"
	"encoding/json"
	"github.com/spf13/cobra"
	"net/http"
	"io/ioutil"
)

var getBooksCommand = &cobra.Command{
	Use: "getBooks",
	Run: getBooks,
}

func getBooks(cmd *cobra.Command, args []string){

	url := "http://localhost:8080/getBooks"

	response, err := http.Get(url)
	if err != nil{
		fmt.Println(err)
		return
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil{
		fmt.Println(err)
		return
	}

	var books restful.Books
	err = json.Unmarshal(data, &books)
	if err != nil{
		fmt.Println(err)
		return
	}

	for _, book := range books.StoredBooks{
		fmt.Println(book.JSONString())
	}

}

func init() {
}