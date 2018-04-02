package cmd

import (
	"fmt"
	"wuzzapcom/Coursework/cli/src/cmd/books"
	"wuzzapcom/Coursework/cli/src/cmd/courses"
	"wuzzapcom/Coursework/cli/src/cmd/departments"
	"wuzzapcom/Coursework/cli/src/cmd/lecturers"
	"wuzzapcom/Coursework/cli/src/cmd/literature"
	"wuzzapcom/Coursework/cli/src/cmd/literature_lists"

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
	rootCommand.AddCommand(books.BookCommand)
	rootCommand.AddCommand(lecturers.LecturerCommand)
	rootCommand.AddCommand(literature.LiteratureCommand)
	rootCommand.AddCommand(literature_lists.LiteratureListCommand)
}

//Execute ..
func Execute() {
	rootCommand.Execute()
}
