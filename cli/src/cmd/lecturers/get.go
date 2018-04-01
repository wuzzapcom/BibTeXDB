package lecturers

import (
	"fmt"
	"wuzzapcom/Coursework/api/src/restful"
	"wuzzapcom/Coursework/cli/src/cmd/helpers"

	"github.com/spf13/cobra"
)

var getCommand = &cobra.Command{
	Use:   "get",
	Run:   getLecturers,
	Short: "Получить список лекторов, сохраненных в базе данных.",
}

var getLecturersOutputFlag = "toFile"

func getLecturers(cmd *cobra.Command, args []string) {
	output := cmd.Flag(getLecturersOutputFlag).Value.String()

	url := helpers.ServerURL + "getLecturers"
	var lecturers restful.Lecturers

	answer, err := helpers.GetFromServer(url, &lecturers)
	if err != nil {
		fmt.Println(err)
		return
	}

	casted, ok := answer.(*restful.Lecturers)
	if !ok {
		fmt.Println("Type cast error")
		return
	}

	helpers.PrintResult(*casted, output)

}

func init() {
	getCommand.Flags().String(getLecturersOutputFlag, "", "Set data output. Prints to console if empty.")
}
