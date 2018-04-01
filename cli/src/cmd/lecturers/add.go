package lecturers

import (
	"fmt"
	"wuzzapcom/Coursework/api/src/common"
	"wuzzapcom/Coursework/cli/src/cmd/helpers"

	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use:   "add",
	Run:   addLecturer,
	Short: "Отправить лектора на сервер из файла, заданного флагом.",
}

const inputFileLecturerFlag = "inputFile"

func addLecturer(cmd *cobra.Command, args []string) {
	inputFile := cmd.Flag(inputFileLecturerFlag).Value.String()

	var items common.Lecturer
	data, err := helpers.LoadFromFileAndValidate(items, inputFile)

	url := helpers.ServerURL + "addLecturer"
	err = helpers.SendDataToServer(data, url)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	helpers.DeleteFile(inputFile)
}

func init() {
	addCommand.Flags().String(
		inputFileLecturerFlag,
		lecturerDefaultPath,
		"Set input file for lecturer",
	)
}
