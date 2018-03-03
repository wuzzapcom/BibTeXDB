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

func runSearch(cmd *cobra.Command, args []string){
	encodedRequest := url.QueryEscape(cmd.Flag("request").Value.String())
	url := "http://localhost:8080/search?request=" + encodedRequest
	response, err := http.Get(url)
	if err != nil{
		fmt.Println(err)
	}
	answer, err := ioutil.ReadAll(response.Body)
	if err != nil{
		fmt.Println(err)
	}
	var search restful.Search
	err = json.Unmarshal(answer, &search)
	if err != nil{
		fmt.Println(err)
	}
	resultFile, err := os.Create("searchResults.txt")
	if err != nil{
		fmt.Println(err)
	}

	resultFile.WriteString(search.String())
	resultFile.Close()

	fmt.Println("Open searchResults.txt, view results, remove wrong items and fix incorrect data.")
}

func init(){
	searchCommand.Flags().String("request", "Compilers", "Insert request for searching")
}
