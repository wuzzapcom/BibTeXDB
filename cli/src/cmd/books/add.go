package books

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"wuzzapcom/Coursework/api/src/common"
	"wuzzapcom/Coursework/cli/src/cmd/helpers"

	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use:   "add",
	Run:   addBooks,
	Short: "Отправить книгу на сервер из файла, заданного флагом.",
}

const inputFileBooksFlag = "inputFile"

var resultFilePath = "searchResults.txt"

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
		url := helpers.ServerURL + "addBook"
		helpers.SendDataToServer(itemBytes, url)
	}

}

func init() {
	addCommand.Flags().String(inputFileBooksFlag, resultFilePath, "Set input file")
}
