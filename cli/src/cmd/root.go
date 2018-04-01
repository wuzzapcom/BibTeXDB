package cmd

import (
	"fmt"
	"wuzzapcom/Coursework/cli/src/cmd/courses"

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
	// rootCommand.AddCommand(searchCommand)
	// rootCommand.AddCommand(addBooksCommand)
	// rootCommand.AddCommand(addCourseCommand)
	// rootCommand.AddCommand(getCoursePrototypeCommand)
	// rootCommand.AddCommand(getCoursesCommand)
	// rootCommand.AddCommand(getBooksCommand)
	// rootCommand.AddCommand(getDepartmentsCommand)
	// rootCommand.AddCommand(getDepartmentPrototypeCommand)
	// rootCommand.AddCommand(addDepartmentCommand)
	// rootCommand.AddCommand(getLecturersCommand)
	// rootCommand.AddCommand(getLecturerPrototypeCommand)
	// rootCommand.AddCommand(addLecturerCommand)
	// rootCommand.AddCommand(addLiteratureListCommand)
	// rootCommand.AddCommand(getLiteratureListsCommand)
	// rootCommand.AddCommand(getLiteratureListPrototypeCommand)
	// rootCommand.AddCommand(addLiteratureCommand)
	// rootCommand.AddCommand(getLiteratureCommand)
	// rootCommand.AddCommand(getLiteraturePrototypeCommand)
	// rootCommand.AddCommand(migratePrototypeCommand)
	// rootCommand.AddCommand(migrateCommand)
	// rootCommand.AddCommand(reportCommand)
}

//Execute ..
func Execute() {
	rootCommand.Execute()
}
