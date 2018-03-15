package cmd

import (
	"fmt"
	"os"
	"wuzzapcom/Coursework/api/src/common"

	"github.com/spf13/cobra"
)

var getLecturerPrototypeCommand = &cobra.Command{
	Use: "getLecturerPrototype",
	Run: getLecturerPrototype,
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
	getLecturerPrototypeCommand.Flags().String(
		lecturerOutputFile,
		lecturerDefaultPath,
		"Set output file for Course prototype",
	)
}
