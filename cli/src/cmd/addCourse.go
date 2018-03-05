package cmd

import (
	"github.com/spf13/cobra"
	"io/ioutil"
	"fmt"
	"wuzzapcom/Coursework/api/src/common"
	"encoding/json"
	"net/http"
	"bytes"
	"wuzzapcom/Coursework/api/src/restful"
)

var addCourseCommand = &cobra.Command{
	Use: "addCourse",
	Run: addCourse,
}

const inputFileCourseFlag = "inputFile"

func addCourse(cmd *cobra.Command, args []string){
	inputFile := cmd.Flag(inputFileFlag).Value.String()
	data, err := ioutil.ReadFile(inputFile)
	if err != nil{
		fmt.Println(err)
		return
	}

	var items common.Course
	err = json.Unmarshal(data, &items)
	if err != nil{
		fmt.Println(err)
		return
	}

	url := "http://localhost:8080/addCourse"
	response, err := http.Post(url, "application/json", bytes.NewReader(data))

	answer, err := ioutil.ReadAll(response.Body)
	if err != nil{
		fmt.Println(err)
	}

	if response.StatusCode != 200 {
		handleError(answer)
		return
	}

	var success restful.Success
	err = json.Unmarshal(answer, &success)
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(success)
}

func init(){
	addCourseCommand.Flags().String(
		inputFileFlag,
		courseDefaultPath,
		"Set input file for course",
		)
}