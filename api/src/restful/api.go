package restful

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wuzzapcom/Coursework/api/src/common"
	"wuzzapcom/Coursework/api/src/fetchers"
)

var fetcher fetchers.GoogleFetcher

func searchHandler(w http.ResponseWriter, r *http.Request) {
	common.LogRequest(*r)
	addHeaders(w)
	parameters, err := searchCheckInput(w, r)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}
	search(w, parameters)
}

func addBookHandler(w http.ResponseWriter, r *http.Request) {
	common.LogRequest(*r)
	addHeaders(w)
	body, err := addBookCheckInput(w, r)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}
	addBook(w, body)
}

func getBooksHandler(w http.ResponseWriter, r *http.Request) {
	common.LogRequest(*r)
	addHeaders(w)
	getBooks(w)
}

func getCoursePrototypeHandler(w http.ResponseWriter, r *http.Request) {
	common.LogRequest(*r)
	addHeaders(w)
	getCoursePrototype(w)
}

func getDepartmentPrototypeHandler(w http.ResponseWriter, r *http.Request) {
	common.LogRequest(*r)
	addHeaders(w)
	getDepartmentPrototype(w)
}

func addDepartmentHandler(w http.ResponseWriter, r *http.Request) {
	common.LogRequest(*r)
	addHeaders(w)
	body, err := addDepartmentCheckInput(w, r)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}
	addDepartment(w, body)
}

func getDepartmentsHandler(w http.ResponseWriter, r *http.Request) {
	common.LogRequest(*r)
	addHeaders(w)
	getDepartments(w)
}

func getLecturerPrototypeHandler(w http.ResponseWriter, r *http.Request) {
	common.LogRequest(*r)
	addHeaders(w)
	getLecturerPrototype(w)
}

func addLecturerHandler(w http.ResponseWriter, r *http.Request) {
	common.LogRequest(*r)
	addHeaders(w)
	body, err := addLecturerCheckInput(w, r)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}
	addLecturer(w, body)
}

func getLecturersHandler(w http.ResponseWriter, r *http.Request) {
	common.LogRequest(*r)
	addHeaders(w)
	getLecturers(w)
}

func getLiteratureListPrototypeHandler(w http.ResponseWriter, r *http.Request) {
	common.LogRequest(*r)
	addHeaders(w)
	getLiteratureListPrototype(w)
}

func addLiteratureListHandler(w http.ResponseWriter, r *http.Request) {
	common.LogRequest(*r)
	addHeaders(w)
	body, err := addLiteratureListCheckInput(w, r)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}
	addLiteratureList(w, body)
}

func getLiteratureListsHandler(w http.ResponseWriter, r *http.Request) {
	common.LogRequest(*r)
	addHeaders(w)
	getLiteratureLists(w)
}

func addLiteratureHandler(w http.ResponseWriter, r *http.Request) {
	common.LogRequest(*r)
	addHeaders(w)
	body, err := addLiteratureCheckInput(w, r)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}
	addLiterature(w, body)
}

func getLiteraturePrototypeHandler(w http.ResponseWriter, r *http.Request) {
	common.LogRequest(*r)
	addHeaders(w)
	getLiteraturePrototype(w)
}

func getLiteratureHandler(w http.ResponseWriter, r *http.Request) {
	common.LogRequest(*r)
	addHeaders(w)
	request, err := addLiteratureCheckInput(w, r)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}
	getLiterature(request, w)
}

func addCourseHandler(w http.ResponseWriter, r *http.Request) {
	common.LogRequest(*r)
	addHeaders(w)
	body, err := addCourseCheckInput(w, r)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}
	addCourse(w, body)
}

func getCoursesHandler(w http.ResponseWriter, r *http.Request) {
	common.LogRequest(*r)
	addHeaders(w)
	getCourses(w)
}

func migrateLiteratureListHandler(w http.ResponseWriter, r *http.Request) {
	common.LogRequest(*r)
	addHeaders(w)
	body, err := migrateLiteratureListCheckInput(w, r)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}
	migrateLiteratureList(w, body)
}

func generateBibTexHandler(w http.ResponseWriter, r *http.Request) {
	common.LogRequest(*r)
	addHeaders(w)
	body, err := generateBibTexCheckInput(w, r)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}
	generateBibTex(w, body)
}

func addHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func returnError(w http.ResponseWriter, code int, message string) {
	answer, err := json.Marshal(Error{message})
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	common.LogErrorResponseWriter(code, message)

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

	http.HandleFunc("/getDepartmentPrototype", getDepartmentPrototypeHandler)
	http.HandleFunc("/addDepartment", addDepartmentHandler)
	http.HandleFunc("/getDepartments", getDepartmentsHandler)

	http.HandleFunc("/getLecturerPrototype", getLecturerPrototypeHandler)
	http.HandleFunc("/addLecturer", addLecturerHandler)
	http.HandleFunc("/getLecturers", getLecturersHandler)

	http.HandleFunc("/getLiteratureListPrototype", getLiteratureListPrototypeHandler)
	http.HandleFunc("/addLiteratureList", addLiteratureListHandler)
	http.HandleFunc("/getLiteratureLists", getLiteratureListsHandler)

	http.HandleFunc("/getLiteraturePrototype", getLiteraturePrototypeHandler)
	http.HandleFunc("/addLiterature", addLiteratureHandler)
	http.HandleFunc("/getLiterature", getLiteratureHandler)

	http.HandleFunc("/migrateLiteratureList", migrateLiteratureListHandler)
	http.HandleFunc("/generateReport", generateBibTexHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
	}
}
