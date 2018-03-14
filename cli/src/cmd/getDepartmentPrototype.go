package cmd

import (
	"fmt"
	"os"
	"wuzzapcom/Coursework/api/src/common"

	"github.com/spf13/cobra"
)

var getDepartmentPrototypeCommand = &cobra.Command{
	Use: "getDepartmentPrototype",
	Run: getDepartmentPrototype,
}

const departmentOutputFile = "outputFile"
const departmentDefaultPath = "department.txt"

func getDepartmentPrototype(cmd *cobra.Command, args []string) {
	outputFile := cmd.Flag(departmentOutputFile).Value.String()

	resultFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	resultFile.WriteString(common.GetDepartmentExample().String())
	resultFile.Close()

	fmt.Println(fmt.Sprintf("Open %s and fill prototype struct with correct data", outputFile))
}

func init() {
	getDepartmentPrototypeCommand.Flags().String(
		departmentOutputFile,
		departmentDefaultPath,
		"Set output file for Course prototype",
	)
}
