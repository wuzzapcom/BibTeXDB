package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"wuzzapcom/Coursework/api/src/restful"

	"github.com/spf13/cobra"
)

var getDepartmentsCommand = &cobra.Command{
	Use: "getDepartments",
	Run: getDepartments,
	Short: "Получить список факультетов, сохраненных в базе данных.",	
}

func getDepartments(cmd *cobra.Command, args []string) {

	url := "http://localhost:8080/getDepartments"

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	if response.StatusCode != 200 {
		handleError(data)
		return
	}
	var departments restful.Departments
	err = json.Unmarshal(data, &departments)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	for _, department := range departments.DepartmentList {
		fmt.Println(department.String())
	}

}
