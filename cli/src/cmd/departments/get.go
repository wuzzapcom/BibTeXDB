package departments

import (
	"fmt"
	"wuzzapcom/Coursework/api/src/restful"
	"wuzzapcom/Coursework/cli/src/cmd/helpers"

	"github.com/spf13/cobra"
)

var getCommand = &cobra.Command{
	Use:   "get",
	Run:   getDepartments,
	Short: "Получить список факультетов, сохраненных в базе данных.",
}

var getOutputFlag = "toFile"

func getDepartments(cmd *cobra.Command, args []string) {
	output := cmd.Flag(getOutputFlag).Value.String()

	url := helpers.ServerURL + "getDepartments"
	var departments restful.Departments
	answer, err := helpers.GetFromServer(url, &departments)
	if err != nil {
		fmt.Println(err)
		return
	}

	casted, ok := answer.(*restful.Departments)
	if !ok {
		fmt.Println("Type cast error")
		return
	}

	departments = *casted
	helpers.PrintResult(departments, output)
}

func init() {
	getCommand.Flags().String(getOutputFlag, "", "Set data output. Prints to console if empty.")
}
