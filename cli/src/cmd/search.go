package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"wuzzapcom/Coursework/api/src/restful"
	"os"
)

var searchCommand = &cobra.Command{
	Use: "search",
	Run: runSearch,
}

var resultFilePath = "searchResults.txt"

const requestFlag = "request"
const outputFileFlag = "outputFile"

func runSearch(cmd *cobra.Command, args []string){
	resultFilePath = cmd.Flag(outputFileFlag).Value.String()
	encodedRequest := url.QueryEscape(cmd.Flag(requestFlag).Value.String())
	url := "http://localhost:8080/search?request=" + encodedRequest
	response, err := http.Get(url)
	if err != nil{
		fmt.Println(err)
	}
	answer, err := ioutil.ReadAll(response.Body)
	if err != nil{
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		handleError(answer)
		return
	}
	var search restful.Search
	err = json.Unmarshal(answer, &search)
	if err != nil{
		fmt.Println(err)
	}
	resultFile, err := os.Create(resultFilePath)
	if err != nil{
		fmt.Println(err)
	}

	resultFile.WriteString(search.String())
	resultFile.Close()

	fmt.Println("Open searchResults.txt, view results, remove wrong items and fix incorrect data.")
}

func init(){
	searchCommand.Flags().String(requestFlag, "Compilers", "Insert request for searching")
	searchCommand.Flags().String(outputFileFlag, resultFilePath, "Set output file")
}
