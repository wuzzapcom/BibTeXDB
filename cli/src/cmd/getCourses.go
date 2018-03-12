package cmd

import (
	"github.com/spf13/cobra"
	"net/http"
	"fmt"
	"wuzzapcom/Coursework/api/src/restful"
	"encoding/json"
	"io/ioutil"
	"os"
)

var getCoursesCommand = &cobra.Command{
	Use: "getCourses",
	Run: getCourses,
}

var getCoursesOutputFlag = "toFile"

func getCourses(cmd *cobra.Command, args []string){
	output := cmd.Flag(getCoursesOutputFlag).Value.String()

	url := "http://localhost:8080/getCourses"

	response, err := http.Get(url)
	if err != nil{
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil{
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	if response.StatusCode != 200 {
		handleError(data)
		return
	}

	var courses restful.Courses
	err = json.Unmarshal(data, &courses)
	if err != nil{
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	var result string
	for _, course := range courses.CoursesList{
		result += course.String()
		result += "\n"
	}

	if output == ""{
		fmt.Println(result)
	}else{
		file, err := os.Create(output)
		if err != nil{
			fmt.Printf("FATAL: %+v\n", err)
			return
		}
		file.Write([]byte(result))
	}
}

func init() {
	getCoursesCommand.Flags().String(getCoursesOutputFlag, "", "Set data output. Prints to console if empty.")
}