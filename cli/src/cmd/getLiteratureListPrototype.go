package cmd

import (
	"fmt"
	"os"
	"wuzzapcom/Coursework/api/src/common"

	"github.com/spf13/cobra"
)

var getLiteratureListPrototypeCommand = &cobra.Command{
	Use:   "getLiteratureListPrototype",
	Run:   getLiteratureListPrototype,
	Short: "Получить заготовку JSON для списка литературы в файл, определяемый флагом.",
	Long:  "Получить заготовку JSON для списка литературы. После чего следует заполнить его вручную и отправить соответствующей командой.",
}

const literatureListOutputFile = "outputFile"
const literatureListDefaultPath = "literatureList.txt"

func getLiteratureListPrototype(cmd *cobra.Command, args []string) {
	outputFile := cmd.Flag(literatureListOutputFile).Value.String()

	resultFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	resultFile.WriteString(common.GetLiteratureListExample().String())
	resultFile.Close()

	fmt.Println(fmt.Sprintf("Open %s and fill prototype struct with correct data", outputFile))
}

func init() {
	getLiteratureListPrototypeCommand.Flags().String(
		literatureListOutputFile,
		literatureListDefaultPath,
		"Set output file for Course prototype",
	)
}
