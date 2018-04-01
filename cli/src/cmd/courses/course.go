package courses

import "github.com/spf13/cobra"

var CourseCommand = &cobra.Command{
	Use: "course",
	// Run:   addCourse,
	Short: "Команда для работы с курсами",
}

func init() {
	CourseCommand.AddCommand(addCommand)
	CourseCommand.AddCommand(getCommand)
	CourseCommand.AddCommand(getCoursePrototypeCommand)
}
