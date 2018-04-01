package courses

import (
	"fmt"
	"wuzzapcom/Coursework/api/src/restful"
	"wuzzapcom/Coursework/cli/src/cmd/helpers"

	"github.com/spf13/cobra"
)

var getCommand = &cobra.Command{
	Use:   "get",
	Run:   getCourses,
	Short: "Получить список курсов, сохраненных в базе данных.",
}

var getCoursesOutputFlag = "toFile"

func getCourses(cmd *cobra.Command, args []string) {
	output := cmd.Flag(getCoursesOutputFlag).Value.String()

	url := helpers.ServerURL + "getCourses"
	var courses restful.Courses

	answer, err := helpers.GetFromServer(url, &courses)
	if err != nil {
		fmt.Println(err)
		return
	}

	castedCourses, ok := answer.(*restful.Courses)
	if !ok {
		fmt.Println("Type cast error")
		return
	}
	courses = *castedCourses

	helpers.PrintResult(courses, output)
}

func init() {
	getCommand.Flags().String(getCoursesOutputFlag, "", "Set data output. Prints to console if empty.")
}
