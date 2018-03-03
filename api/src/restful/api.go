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
	internalServerError(w, "Not implemented")
}

func addCourseHandler(w http.ResponseWriter, r *http.Request){
	internalServerError(w, "Not implemented")
}

func addCourseLiterature(w http.ResponseWriter, r *http.Request){
	internalServerError(w, "Not implemented")

}

func getCoursesHandler(w http.ResponseWriter, r *http.Request){
	internalServerError(w, "Not implemented")

}

func getCourseLiteratureHandler(w http.ResponseWriter, r *http.Request){
	internalServerError(w, "Not implemented")
}

func internalServerError(w http.ResponseWriter, message string){
	answer, err := json.Marshal(Error{message})
	if err != nil {
		fmt.Println(err)
		return
	}

	w.WriteHeader(500)
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