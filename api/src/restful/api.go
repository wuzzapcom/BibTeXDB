package restful

import (
	"net/http"
	"encoding/json"
	"fmt"
	"wuzzapcom/Coursework/api/src/fetchers"
)

var fetcher fetchers.GoogleFetcher

func searchHandler(w http.ResponseWriter, r *http.Request){
	parameters, err := searchCheckInput(w, r)
	if err != nil{
		fmt.Println(err)
		return
	}
	search(w, parameters)
}

func addBookHandler(w http.ResponseWriter, r *http.Request){
	body, err := addBookCheckInput(w, r)
	if err != nil{
		fmt.Println(err)
		return
	}
	addBook(w, body)
}

func addCourseHandler(w http.ResponseWriter, r *http.Request){
	returnError(w, 501, "Not implemented")
}

func addCourseLiterature(w http.ResponseWriter, r *http.Request){
	returnError(w, 501, "Not implemented")

}

func getCoursesHandler(w http.ResponseWriter, r *http.Request){
	returnError(w, 501, "Not implemented")

}

func getCourseLiteratureHandler(w http.ResponseWriter, r *http.Request){
	returnError(w, 501, "Not implemented")
}

func returnError(w http.ResponseWriter, code int, message string){
	answer, err := json.Marshal(Error{message})
	if err != nil {
		fmt.Println(err)
		return
	}

	w.WriteHeader(code)
	w.Write(answer)
}

func Run(f fetchers.GoogleFetcher){
	fetcher = f
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/addBook", addBookHandler)
	http.HandleFunc("/addCourse", addCourseHandler)
	http.HandleFunc("/addCourseLiterature", addCourseLiterature)
	http.HandleFunc("/getCourses", getCoursesHandler)
	http.HandleFunc("/getCourseLiterature", getCourseLiteratureHandler)
	http.ListenAndServe(":8080", nil)
}