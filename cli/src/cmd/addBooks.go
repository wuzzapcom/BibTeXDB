package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"encoding/json"
	"wuzzapcom/Coursework/api/src/common"
	"net/http"
	"bytes"
	"io/ioutil"
	"wuzzapcom/Coursework/api/src/restful"
)

var addBooksCommand = &cobra.Command{
	Use: "addBooks",
	Run: addBooks,
}

const inputFileFlag = "inputFile"

func addBooks(cmd *cobra.Command, args []string) {
	resultFilePath = cmd.Flag(inputFileFlag).Value.String()
	data, err := ioutil.ReadFile(resultFilePath)
	if err != nil{
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	var items common.Item
	err = json.Unmarshal(data, &items)
	if err != nil{
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	url := "http://localhost:8080/addBook"
	response, err := http.Post(url, "application/json", bytes.NewReader(data))

	answer, err := ioutil.ReadAll(response.Body)
	if err != nil{
		fmt.Printf("FATAL: %+v\n", err)
	}

	if response.StatusCode != 200 {
		handleError(answer)
		return
	}

	var success restful.Success
	err = json.Unmarshal(answer, &success)
	if err != nil{
		fmt.Printf("FATAL: %+v\n", err)
	}

	fmt.Println(success)

}

func init(){
	addBooksCommand.Flags().String(inputFileFlag, resultFilePath, "Set input file")
}