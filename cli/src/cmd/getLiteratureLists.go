package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"wuzzapcom/Coursework/api/src/restful"

	"github.com/spf13/cobra"
)

var getLiteratureListsCommand = &cobra.Command{
	Use:   "getLiteratureLists",
	Run:   getLiteratureLists,
	Short: "Получить все списки литературы, сохраненных в базе данных.",
}

func getLiteratureLists(cmd *cobra.Command, args []string) {

	url := "http://localhost:8080/getLiteratureLists"

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
	var departments restful.LiteratureLists
	err = json.Unmarshal(data, &departments)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	for _, department := range departments.StoredLists {
		fmt.Println(department.String())
	}

}
