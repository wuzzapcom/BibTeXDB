package courses

import (
	"fmt"
	"os"
	"wuzzapcom/Coursework/api/src/common"

	"github.com/spf13/cobra"
)

var getCoursePrototypeCommand = &cobra.Command{
	Use:   "prototype",
	Run:   getCoursePrototype,
	Short: "Получить заготовку JSON для курса в файл, определяемый флагом.",
	Long:  "Получить заготовку JSON для курса. После чего следует заполнить его вручную и отправить соответствующей командой.",
}

const courseOutputFile = "outputFile"
const courseDefaultPath = "course.txt"

func getCoursePrototype(cmd *cobra.Command, args []string) {
	outputFile := cmd.Flag(courseOutputFile).Value.String()

	resultFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	resultFile.WriteString(common.GetCourseExample().String())
	resultFile.Close()

	fmt.Println(fmt.Sprintf("Open %s and fill prototype struct with correct data", outputFile))
}

func init() {
	getCoursePrototypeCommand.Flags().String(
		courseOutputFile,
		courseDefaultPath,
		"Set output file for Course prototype",
	)
}
