package departments

import (
	"fmt"
	"wuzzapcom/Coursework/api/src/common"
	"wuzzapcom/Coursework/cli/src/cmd/helpers"

	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use:   "add",
	Run:   addDepartment,
	Short: "Отправить факультет на сервер из файла, заданного флагом.",
}

const inputFileDepartmentFlag = "inputFile"

func addDepartment(cmd *cobra.Command, args []string) {
	inputFile := cmd.Flag(inputFileDepartmentFlag).Value.String()
	var items common.Department

	data, err := helpers.LoadFromFileAndValidate(items, inputFile)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	url := helpers.ServerURL + "addDepartment"
	err = helpers.SendDataToServer(data, url)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	helpers.DeleteFile(inputFile)
}

func init() {
	addCommand.Flags().String(
		inputFileDepartmentFlag,
		departmentDefaultPath,
		"Set input file for department",
	)
}
