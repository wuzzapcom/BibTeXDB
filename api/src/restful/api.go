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

func addLecturerHandler(w http.ResponseWriter, r *http.Request) {
	body, err := addLecturerCheckInput(w, r)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}
	addLecturer(w, body)
}

func getLecturersHandler(w http.ResponseWriter, r *http.Request) {
	getLecturers(w)
}

func addLiteratureListHandler(w http.ResponseWriter, r *http.Request) {
	body, err := addLiteratureListCheckInput(w, r)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}
	addLiteratureList(w, body)
}

func getLiteratureListsHandler(w http.ResponseWriter, r *http.Request) {
	getLiteratureLists(w)
}

func addLiteratureHandler(w http.ResponseWriter, r *http.Request) {
	body, err := addLiteratureCheckInput(w, r)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}
	addLiterature(w, body)
}

func getLiteratureHandler(w http.ResponseWriter, r *http.Request) {
	getLiterature(w)
}

func addCourseHandler(w http.ResponseWriter, r *http.Request) {
	body, err := addCourseCheckInput(w, r)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}
	addCourse(w, body)
}

func getCoursesHandler(w http.ResponseWriter, r *http.Request) {
	getCourses(w)
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

	http.HandleFunc("/addDepartment", addDepartmentHandler)
	http.HandleFunc("/getDepartments", getDepartmentsHandler)

	http.HandleFunc("/addLecturer", addLecturerHandler)
	http.HandleFunc("/getLecturers", getLecturersHandler)

	http.HandleFunc("/addLiteratureList", addLiteratureListHandler)
	http.HandleFunc("/getLiteratureLists", getLiteratureListsHandler)

	http.HandleFunc("/addLiterature", addLiteratureHandler)
	http.HandleFunc("/getLiterature", getLiteratureHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
	}
}
