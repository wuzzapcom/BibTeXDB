package cmd

import (
	"fmt"
	"os"
	"wuzzapcom/Coursework/api/src/common"

	"github.com/spf13/cobra"
)

var getLiteraturePrototypeCommand = &cobra.Command{
	Use:   "getLiteraturePrototype",
	Run:   getLiteraturePrototype,
	Short: "Получить заготовку JSON для добавления книги в список литературы в файл, определяемый флагом.",
	Long:  "Получить заготовку JSON для добавления книги в список литературы. После чего следует заполнить его вручную и отправить соответствующей командой.",
}

const literatureOutputFile = "outputFile"
const literatureDefaultPath = "literature.txt"

func getLiteraturePrototype(cmd *cobra.Command, args []string) {
	outputFile := cmd.Flag(literatureOutputFile).Value.String()

	resultFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	resultFile.WriteString(common.GetLiteratureExample().String())
	resultFile.Close()

	fmt.Println(fmt.Sprintf("Open %s and fill prototype struct with correct data", outputFile))
}

func init() {
	getLiteraturePrototypeCommand.Flags().String(
		literatureOutputFile,
		literatureDefaultPath,
		"Set output file for Course prototype",
	)
}
