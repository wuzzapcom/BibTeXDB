package cmd

import (
	"github.com/spf13/cobra"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
	"wuzzapcom/Coursework/api/src/common"
)

var getCoursePrototypeCommand = &cobra.Command{
	Use: "getCoursePrototype",
	Run: getCoursePrototype,
}

const courseOutputFile = "outputFile"
const courseDefaultPath = "course.txt"

func getCoursePrototype(cmd *cobra.Command, args []string){
	outputFile := cmd.Flag(courseOutputFile).Value.String()

	url := "http://localhost:8080/getCoursePrototype"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	answer, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if response.StatusCode != 200 {
		handleError(answer)
		return
	}
	var course common.Course
	err = json.Unmarshal(answer, &course)
	if err != nil{
		fmt.Println(err)
		return
	}
	resultFile, err := os.Create(outputFile)
	if err != nil{
		fmt.Println(err)
		return
	}

	resultFile.WriteString(course.String())
	resultFile.Close()

	fmt.Println(fmt.Sprintf("Open %s and fill prototype struct with correct data", outputFile))
}

func init(){
	getCoursePrototypeCommand.Flags().String(
		courseOutputFile,
		courseDefaultPath,
		"Set output file for Course prototype",
		)
}