package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"wuzzapcom/Coursework/api/src/restful"

	"github.com/spf13/cobra"
)

var getLiteratureCommand = &cobra.Command{
	Use: "getLiterature",
	Run: getLiterature,
	Short: "Получить список связей книг и списков литературы, сохраненных в базе данных.",	
}

func getLiterature(cmd *cobra.Command, args []string) {

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

}
