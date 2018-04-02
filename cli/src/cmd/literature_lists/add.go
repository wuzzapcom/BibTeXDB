package literature_lists

import (
	"fmt"
	"wuzzapcom/Coursework/api/src/common"
	"wuzzapcom/Coursework/cli/src/cmd/helpers"

	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use:   "add",
	Run:   addLiteratureList,
	Short: "Добавить список литературы из файла, заданного флагом.",
	Long: `Добавить книгу в список литературы. 
	Поле Year определяет, за какой год создается список литературы.
	Поля CourseTitle и DepartmentTitle определяют учебный курс, для которого создается список литературы.`,
}

const inputFileLiteratureListFlag = "inputFile"

func addLiteratureList(cmd *cobra.Command, args []string) {
	inputFile := cmd.Flag(inputFileLiteratureListFlag).Value.String()
	var items common.LiteratureList

	data, err := helpers.LoadFromFileAndValidate(items, inputFile)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	url := helpers.ServerURL + "addLiteratureList"
	err = helpers.SendDataToServer(data, url)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	helpers.DeleteFile(inputFile)
}

func init() {
	addCommand.Flags().String(
		inputFileLiteratureListFlag,
		literatureListDefaultPath,
		"Set input file for literatureList",
	)
}
