package courses

import (
	"fmt"
	"wuzzapcom/Coursework/api/src/common"
	"wuzzapcom/Coursework/cli/src/cmd/helpers"

	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use:   "add",
	Run:   addCourse,
	Short: "Отправить курс на сервер из файла, заданного флагом.",
}

const inputFileCourseFlag = "inputFile"

func addCourse(cmd *cobra.Command, args []string) {
	inputFile := cmd.Flag(inputFileCourseFlag).Value.String()

	var items common.Course
	data, err := helpers.LoadFromFileAndValidate(items, inputFile)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	url := helpers.ServerURL + "addCourse"
	err = helpers.SendDataToServer(data, url)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	helpers.DeleteFile(inputFile)
}

func init() {
	addCommand.Flags().String(
		inputFileCourseFlag,
		courseDefaultPath,
		"Set input file for course",
	)
}
