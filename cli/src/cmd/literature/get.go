package literature

import (
	"fmt"
	"wuzzapcom/Coursework/api/src/common"
	"wuzzapcom/Coursework/api/src/restful"
	"wuzzapcom/Coursework/cli/src/cmd/helpers"

	"github.com/spf13/cobra"
)

var getCommand = &cobra.Command{
	Use:   "get",
	Run:   getLiterature,
	Short: "Получить список связей книг и списков литературы, сохраненных в базе данных.",
}

const getLiteratureForListFlag = "forList"
const getLiteratureFromFileFlag = "inputFile"
const getLiteratureOutputFlag = "toFile"

var isForList = false

func getLiterature(cmd *cobra.Command, args []string) {
	output := cmd.Flag(getLiteratureOutputFlag).Value.String()
	url := helpers.ServerURL + "getLiterature"

	if !isForList {
		requestWithoutBody(output, url)
	} else {
		requestWithBody(cmd, output, url)
	}

}

func requestWithoutBody(output string, url string) {
	var literatures restful.Literature

	answer, err := helpers.GetFromServer(url, &literatures)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	casted, ok := answer.(*restful.Literature)
	if !ok {
		fmt.Println("Type cast error")
		return
	}
	helpers.PrintResult(*casted, output)
}

func requestWithBody(cmd *cobra.Command, output string, url string) {
	inputFile := cmd.Flag(inputFileLiteratureFlag).Value.String()
	var list common.LiteratureList

	data, err := helpers.LoadFromFileAndValidate(&list, inputFile)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	var literatures restful.LiteratureLists

	answer, err := helpers.GetDataFromServerWithBody(data, &literatures, url)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	casted, ok := answer.(*restful.LiteratureLists)
	if !ok {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	helpers.PrintResult(casted, output)
}

func init() {
	getCommand.Flags().BoolVar(&isForList, getLiteratureForListFlag, false, "usage string")
	getCommand.Flags().String(
		getLiteratureFromFileFlag,
		helpers.LiteratureListDefaultPath,
		"Set input file for literature",
	)
	getCommand.Flags().String(getLiteratureOutputFlag, "", "Set data output. Prints to console if empty.")
}
