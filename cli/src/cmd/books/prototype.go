package books

import (
	"fmt"
	"os"
	"wuzzapcom/Coursework/api/src/common"

	"github.com/spf13/cobra"
)

var prototypeCommand = &cobra.Command{
	Use:   "prototype",
	Run:   getBookPrototype,
	Short: "Получить заготовку JSON для книги в файл, определяемый флагом.",
	Long:  "Получить заготовку JSON для книги. После чего следует заполнить его вручную и отправить соответствующей командой.",
}

const outputFile = "outputFile"
const defaultPath = "book.txt"

func getBookPrototype(cmd *cobra.Command, args []string) {
	outputFile := cmd.Flag(outputFile).Value.String()

	resultFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	resultFile.WriteString(common.GetRandomItems()[0].String())
	resultFile.Close()

	fmt.Println(fmt.Sprintf("Open %s and fill prototype struct with correct data", outputFile))
}

func init() {
	prototypeCommand.Flags().String(
		outputFile,
		defaultPath,
		"Set output file for Course prototype",
	)
}
