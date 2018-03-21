package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"wuzzapcom/Coursework/api/src/common"
	"wuzzapcom/Coursework/api/src/restful"

	"github.com/spf13/cobra"
)

var addBooksCommand = &cobra.Command{
	Use:   "addBooks",
	Run:   addBooks,
	Short: "Отправить книгу на сервер из файла, заданного флагом.",
}

const inputFileFlag = "inputFile"

func addBooks(cmd *cobra.Command, args []string) {
	resultFilePath = cmd.Flag(inputFileFlag).Value.String()
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
	addBooksCommand.Flags().String(inputFileFlag, resultFilePath, "Set input file")
}
