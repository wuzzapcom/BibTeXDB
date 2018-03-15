package cmd

import (
	"encoding/json"
	"fmt"
	"wuzzapcom/Coursework/api/src/restful"

	"github.com/spf13/cobra"
)

func handleError(answer []byte) {
	var errorAnswer restful.Error
	err := json.Unmarshal(answer, &errorAnswer)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
	}

	fmt.Println(errorAnswer.Message)
}

var rootCommand = &cobra.Command{
	Use: "cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cli")
	},
}

func init() {
	rootCommand.AddCommand(searchCommand)
	rootCommand.AddCommand(addBooksCommand)
	rootCommand.AddCommand(addCourseCommand)
	rootCommand.AddCommand(getCoursePrototypeCommand)
	rootCommand.AddCommand(getCoursesCommand)
	rootCommand.AddCommand(getBooksCommand)
	rootCommand.AddCommand(getDepartmentsCommand)
	rootCommand.AddCommand(getDepartmentPrototypeCommand)
	rootCommand.AddCommand(addDepartmentCommand)
	rootCommand.AddCommand(getLecturersCommand)
}

//Execute ..
func Execute() {
	rootCommand.Execute()
}
