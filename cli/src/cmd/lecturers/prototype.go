package lecturers

import (
	"fmt"
	"os"
	"wuzzapcom/Coursework/api/src/common"

	"github.com/spf13/cobra"
)

var prototypeCommand = &cobra.Command{
	Use:   "prototype",
	Run:   getLecturerPrototype,
	Short: "Получить заготовку JSON для лектора в файл, определяемый флагом.",
	Long:  "Получить заготовку JSON для лектора. После чего следует заполнить его вручную и отправить соответствующей командой.",
}

const lecturerOutputFile = "outputFile"
const lecturerDefaultPath = "lecturer.txt"

func getLecturerPrototype(cmd *cobra.Command, args []string) {
	outputFile := cmd.Flag(lecturerOutputFile).Value.String()

	resultFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	resultFile.WriteString(common.GetLecturerExample().String())
	resultFile.Close()

	fmt.Println(fmt.Sprintf("Open %s and fill prototype struct with correct data", outputFile))
}

func init() {
	prototypeCommand.Flags().String(
		lecturerOutputFile,
		lecturerDefaultPath,
		"Set output file for Course prototype",
	)
}
