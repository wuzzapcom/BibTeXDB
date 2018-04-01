package cmd

import (
	"fmt"
	"wuzzapcom/Coursework/cli/src/cmd/courses"
	"wuzzapcom/Coursework/cli/src/cmd/departments"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use: "cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cli")
	},
}

func init() {
	rootCommand.AddCommand(courses.CourseCommand)
	rootCommand.AddCommand(departments.DepartmentCommand)
}

//Execute ..
func Execute() {
	rootCommand.Execute()
}
