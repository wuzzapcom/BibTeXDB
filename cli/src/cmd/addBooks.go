package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"wuzzapcom/Coursework/api/src/common"
	"wuzzapcom/Coursework/api/src/restful"

	"github.com/spf13/cobra"
)

var addBooksCommand = &cobra.Command{
	Use:   "addBooks",
	Run:   addBooks,
	Short: "Отправить книгу на сервер из файла, заданного флагом.",
}

const inputFileBooksFlag = "inputFile"

func addBooks(cmd *cobra.Command, args []string) {
	resultFilePath = cmd.Flag(inputFileBooksFlag).Value.String()
	data, err := ioutil.ReadFile(resultFilePath)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	var items common.Items
	err = json.Unmarshal(data, &items)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	for _, item := range items {
		if item.Ident == "empty" {
			fmt.Println("Заполните поле Ident латинскими буквами без пробелов. Это будет идентификатор, по которому " +
				"вы будете добавлять библиографическую ссылку в документ.")
			return
		}

		matches, err := regexp.MatchString("[a-zA-Z]+", item.Ident)
		if err != nil {
			fmt.Printf("FATAL: %+v\n", err)
			return
		}

		if !matches {
			fmt.Printf("Идентификатор %s не подходит под требования для идентификаторов в LaTeX.\n", item.Ident)
			return
		}

	}

	for _, item := range items {
		itemBytes, err := json.Marshal(item)
		if err != nil {
			fmt.Printf("FATAL: %+v\n", err)
			continue
		}
		sendRequestForBook(itemBytes)
	}

}

func sendRequestForBook(data []byte) {
	url := "http://localhost:8080/addBook"
	response, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
	}

	answer, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
	}

	if response.StatusCode != 200 {
		handleError(answer)
		return
	}

	var success restful.Success
	err = json.Unmarshal(answer, &success)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
	}

	fmt.Println(success)

}

func init() {
	addBooksCommand.Flags().String(inputFileBooksFlag, resultFilePath, "Set input file")
}
