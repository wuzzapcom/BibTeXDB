package literature

import (
	"fmt"
	"wuzzapcom/Coursework/api/src/common"
	"wuzzapcom/Coursework/cli/src/cmd/helpers"

	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use:   "add",
	Run:   addLiterature,
	Short: "Добавить книгу в список литературы из файла, заданного флагом.",
	Long: `Добавить книгу в список литературы. 
	Поле BookIdent задает идентификатор книги в формате BibTeX.
	Поле Year определяет, за какой год используется список литературы.
	Поля CourseTitle и DepartmentTitle определяют учебный курс.
	`,
}

const inputFileLiteratureFlag = "inputFile"

func addLiterature(cmd *cobra.Command, args []string) {
	inputFile := cmd.Flag(inputFileLiteratureFlag).Value.String()
	var items common.Literature
	data, err := helpers.LoadFromFileAndValidate(items, inputFile)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
	}

	url := helpers.ServerURL + "addLiterature"
	err = helpers.SendDataToServer(data, url)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
	}

	helpers.DeleteFile(inputFile)
}

func init() {
	addCommand.Flags().String(
		inputFileLiteratureFlag,
		literatureDefaultPath,
		"Set input file for literature",
	)
}
