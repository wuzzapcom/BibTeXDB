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

var getLiteratureCommand = &cobra.Command{
	Use:   "getLiterature",
	Run:   getLiterature,
	Short: "Получить список связей книг и списков литературы, сохраненных в базе данных.",
}

const getLiteratureForListFlag = "forList"
const getLiteratureFromFileFlag = "inputFile"

var isForList = false

func getLiterature(cmd *cobra.Command, args []string) {

	if !isForList {
		url := "http://localhost:8080/getLiterature"
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
		var literatures restful.Literature
		err = json.Unmarshal(data, &literatures)
		if err != nil {
			fmt.Printf("FATAL: %+v\n", err)
			return
		}
		for _, literature := range literatures.StoredLiterature {
			fmt.Println(literature.String())
		}
	} else {
		inputFile := cmd.Flag(inputFileLiteratureFlag).Value.String()
		data, err := ioutil.ReadFile(inputFile)
		if err != nil {
			fmt.Printf("FATAL: %+v\n", err)
			return
		}

		var list common.LiteratureList
		err = json.Unmarshal(data, &list)
		if err != nil {
			fmt.Printf("FATAL: %+v\n", err)
			return
		}
		url := "http://localhost:8080/getLiterature"
		response, err := http.Post(url, "application/json", bytes.NewReader(data))
		if err != nil {
			fmt.Printf("FATAL: %+v\n", err)
			return
		}
		answer, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("FATAL: %+v\n", err)
			return
		}
		if response.StatusCode != 200 {
			handleError(answer)
			return
		}
		var items restful.Books
		err = json.Unmarshal(answer, &items)
		if err != nil {
			fmt.Printf("FATAL: %+v\n", err)
			return
		}
		fmt.Println(items)
	}

}

func init() {
	getLiteratureCommand.Flags().BoolVar(&isForList, getLiteratureForListFlag, false, "usage string")
	getLiteratureCommand.Flags().String(
		getLiteratureFromFileFlag,
		literatureListDefaultPath,
		"Set input file for literature",
	)
}
