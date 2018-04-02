package cmd

import (
	"fmt"
	"net/url"
	"wuzzapcom/Coursework/api/src/restful"
	"wuzzapcom/Coursework/cli/src/cmd/helpers"

	"github.com/spf13/cobra"
)

var searchCommand = &cobra.Command{
	Use:   "search",
	Run:   runSearch,
	Short: "Выполнить поиск книг в онлайн-источниках.",
	Long: `Выполнить поиск книг в онлайн источниках. Флаг --request является обязательным и необходим для задания запроса. 
Найденные книги добавляются в файл, переданный флагом --outputFile`,
	Example: "cli search --request=Компиляторы --outputFile=searchResults.txt",
}

var resultFilePath = "searchResults.txt"

const requestFlag = "request"
const outputFileFlag = "outputFile"

func runSearch(cmd *cobra.Command, args []string) {
	resultFilePath = cmd.Flag(outputFileFlag).Value.String()
	encodedRequest := url.QueryEscape(cmd.Flag(requestFlag).Value.String())
	url := "http://localhost:8080/search?request=" + encodedRequest

	var search restful.Search
	answer, err := helpers.GetFromServer(url, &search)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}
	casted, ok := answer.(*restful.Search)
	if !ok {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	helpers.PrintResult(*casted, resultFilePath)

	fmt.Println(fmt.Sprintf("Open %s, view results, remove wrong items and fix incorrect data.", resultFilePath))
}

func init() {
	searchCommand.Flags().String(requestFlag, "Compilers", "Insert request for searching")
	searchCommand.Flags().String(outputFileFlag, resultFilePath, "Set output file")
}
