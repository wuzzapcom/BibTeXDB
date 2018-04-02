package literature_lists

import (
	"fmt"
	"wuzzapcom/Coursework/api/src/restful"
	"wuzzapcom/Coursework/cli/src/cmd/helpers"

	"github.com/spf13/cobra"
)

var getCommand = &cobra.Command{
	Use:   "get",
	Run:   getLiteratureLists,
	Short: "Получить все списки литературы, сохраненных в базе данных.",
}

var getLiteratureListsOutputFlag = "toFile"

func getLiteratureLists(cmd *cobra.Command, args []string) {
	output := cmd.Flag(getLiteratureListsOutputFlag).Value.String()

	url := helpers.ServerURL + "getLiteratureLists"
	var lists restful.LiteratureLists

	answer, err := helpers.GetFromServer(url, &lists)
	if err != nil {
		fmt.Println(err)
		return
	}

	casted, ok := answer.(*restful.LiteratureLists)
	if !ok {
		fmt.Println("Type cast error")
		return
	}

	helpers.PrintResult(*casted, output)

}

func init() {
	getCommand.Flags().String(getLiteratureListsOutputFlag, "", "Set data output. Prints to console if empty.")
}
