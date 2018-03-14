package restful

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wuzzapcom/Coursework/api/src/fetchers"
)

var fetcher fetchers.GoogleFetcher

func searchHandler(w http.ResponseWriter, r *http.Request) {
	parameters, err := searchCheckInput(w, r)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}
	search(w, parameters)
}

func addBookHandler(w http.ResponseWriter, r *http.Request) {
	body, err := addBookCheckInput(w, r)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}
	addBook(w, body)
}

func getBooksHandler(w http.ResponseWriter, _ *http.Request) {
	getBooks(w)
}

func getCoursePrototypeHandler(w http.ResponseWriter, _ *http.Request) {
	getCoursePrototype(w)
}

func addCourseHandler(w http.ResponseWriter, r *http.Request) {
	body, err := addCourseCheckInput(w, r)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}
	addCourse(w, body)
}

func addCourseLiteratureHandler(w http.ResponseWriter, _ *http.Request) {
	returnError(w, 501, "Not implemented")

}

func getCoursesHandler(w http.ResponseWriter, _ *http.Request) {
	getCourses(w)
}

func getCourseLiteratureHandler(w http.ResponseWriter, _ *http.Request) {
	returnError(w, 501, "Not implemented")
}

func addDepartmentHandler(w http.ResponseWriter, r *http.Request) {
	body, err := addDepartmentCheckInput(w, r)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}
	addDepartment(w, body)
}

func getDepartmentsHandler(w http.ResponseWriter, r *http.Request) {
	getDepartments(w)
}

func returnError(w http.ResponseWriter, code int, message string) {
	answer, err := json.Marshal(Error{message})
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	w.WriteHeader(code)
	w.Write(answer)
}

//Run ..
func Run(f fetchers.GoogleFetcher) {
	fetcher = f

	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/addBook", addBookHandler)
	http.HandleFunc("/getBooks", getBooksHandler)

	http.HandleFunc("/getCoursePrototype", getCoursePrototypeHandler)
	http.HandleFunc("/addCourse", addCourseHandler)
	http.HandleFunc("/getCourses", getCoursesHandler)

	http.HandleFunc("/getCourseLiterature", getCourseLiteratureHandler)
	http.HandleFunc("/addCourseLiterature", addCourseLiteratureHandler)

	http.HandleFunc("/addDepartment", addDepartmentHandler)
	http.HandleFunc("/getDepartments", getDepartmentsHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
	}
}
