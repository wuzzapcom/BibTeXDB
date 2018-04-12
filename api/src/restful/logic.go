package restful

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"wuzzapcom/Coursework/api/src/common"
	"wuzzapcom/Coursework/api/src/database"
	"wuzzapcom/Coursework/api/src/reports"
)

func searchCheckInput(w http.ResponseWriter, r *http.Request) (url.Values, error) {
	parameters, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		returnError(w, 500, common.ErrorMessages[common.ParsingParameterError])
		return nil, err
	}

	if len(parameters["request"]) == 0 {
		answer, err := json.Marshal(Error{common.ErrorMessages[common.NoRequestProvidedError]})
		if err != nil {
			returnError(w, 500, common.ErrorMessages[common.InternalServerError])
			return nil, err
		}

		writeAnswer(w, 400, answer)

	}
	return parameters, nil
}

func search(w http.ResponseWriter, values url.Values) {
	var result common.Items
	for _, request := range values["request"] {
		res, err := fetcher.FetchWithString(request)
		if err != nil {
			fmt.Printf("FATAL: %+v\n", err)
			returnError(w, 500, common.ErrorMessages[common.ParsingParameterError])
			return
		}
		result.Append(res)
	}

	answer, err := json.Marshal(Search{result})
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 500, common.ErrorMessages[common.InternalServerError])
		return
	}

	writeAnswer(w, 200, answer)

}

func addBookCheckInput(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	body := r.Body

	answer, err := ioutil.ReadAll(body)
	if err != nil {
		returnError(w, 400, common.ErrorMessages[common.NoJSONProvidedError])
		return nil, err
	}

	return answer, nil

}

func addBook(w http.ResponseWriter, body []byte) {
	var addingBooks common.Item

	err := json.Unmarshal(body, &addingBooks)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 400, common.ErrorMessages[common.WrongJSONInputError])
		return
	}

	postgres := database.Postgres{}
	postgres.Connect()
	defer postgres.Disconnect()

	err = postgres.InsertTextbook(addingBooks)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 500, common.ErrorMessages[common.InternalServerError])
		return
	}

	answer, err := json.Marshal(Success{"OK"})
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 500, common.ErrorMessages[common.InternalServerError])
		return
	}

	writeAnswer(w, 200, answer)

}

func getBooks(w http.ResponseWriter) {
	postgres := &database.Postgres{}

	postgres.Connect()
	defer postgres.Disconnect()

	textbooks, err := postgres.SelectTextbooks()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 500, common.ErrorMessages[common.InternalServerError])
		return
	}

	data, err := json.Marshal(Books{textbooks})
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 500, common.ErrorMessages[common.InternalServerError])
		return
	}

	writeAnswer(w, 200, data)

}

func getCoursePrototype(w http.ResponseWriter) {

	data, err := json.Marshal(common.GetCourseExample())
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 500, common.ErrorMessages[common.InternalServerError])
		return
	}

	writeAnswer(w, 200, data)

}

func getDepartmentPrototype(w http.ResponseWriter) {

	data, err := json.Marshal(common.GetDepartmentExample())
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 500, common.ErrorMessages[common.InternalServerError])
		return
	}

	writeAnswer(w, 200, data)

}

func addDepartmentCheckInput(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	body := r.Body
	answer, err := ioutil.ReadAll(body)
	if err != nil {
		returnError(w, 400, common.ErrorMessages[common.NoJSONProvidedError])
		return nil, err
	}
	return answer, nil
}

func addDepartment(w http.ResponseWriter, data []byte) {
	var department common.Department

	err := json.Unmarshal(data, &department)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 400, common.ErrorMessages[common.WrongJSONInputError])
		return
	}

	postgres := database.Postgres{}
	postgres.Connect()
	defer postgres.Disconnect()

	err = postgres.InsertDepartment(department)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 500, common.ErrorMessages[common.InternalServerError])
		return
	}

	answer, err := json.Marshal(Success{"OK"})
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 500, common.ErrorMessages[common.InternalServerError])
		return
	}

	writeAnswer(w, 200, answer)

}

func getDepartments(w http.ResponseWriter) {
	postgres := &database.Postgres{}
	postgres.Connect()
	defer postgres.Disconnect()

	departments, err := postgres.SelectDepartments()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 500, common.ErrorMessages[common.InternalServerError])
		return
	}

	data, err := json.Marshal(Departments{departments})
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 500, common.ErrorMessages[common.InternalServerError])
		return
	}

	writeAnswer(w, 200, data)

}

func getLecturerPrototype(w http.ResponseWriter) {

	data, err := json.Marshal(common.GetLecturerExample())
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 500, common.ErrorMessages[common.InternalServerError])
		return
	}

	writeAnswer(w, 200, data)

}

func addLecturerCheckInput(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	body := r.Body
	answer, err := ioutil.ReadAll(body)
	if err != nil {
		returnError(w, 400, common.ErrorMessages[common.NoJSONProvidedError])
		return nil, err
	}
	return answer, nil
}

func addLecturer(w http.ResponseWriter, data []byte) {
	var lecturer common.Lecturer

	err := json.Unmarshal(data, &lecturer)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 400, common.ErrorMessages[common.WrongJSONInputError])
		return
	}

	postgres := database.Postgres{}
	postgres.Connect()
	defer postgres.Disconnect()

	err = postgres.InsertLecturer(lecturer)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 500, common.ErrorMessages[common.InternalServerError])
		return
	}

	answer, err := json.Marshal(Success{"OK"})
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 500, common.ErrorMessages[common.InternalServerError])
		return
	}

	writeAnswer(w, 200, answer)

}

func getLecturers(w http.ResponseWriter) {
	postgres := &database.Postgres{}
	postgres.Connect()
	defer postgres.Disconnect()

	lecturers, err := postgres.SelectLecturers()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 500, common.ErrorMessages[common.InternalServerError])
		return
	}

	data, err := json.Marshal(Lecturers{lecturers})
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 500, common.ErrorMessages[common.InternalServerError])
		return
	}

	writeAnswer(w, 200, data)

}

func getLiteratureListPrototype(w http.ResponseWriter) {

	data, err := json.Marshal(common.GetLiteratureListExample())
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 500, common.ErrorMessages[common.InternalServerError])
		return
	}

	writeAnswer(w, 200, data)

}

func addLiteratureListCheckInput(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	body := r.Body
	answer, err := ioutil.ReadAll(body)
	if err != nil {
		returnError(w, 400, common.ErrorMessages[common.NoJSONProvidedError])
		return nil, err
	}
	return answer, nil
}

func addLiteratureList(w http.ResponseWriter, data []byte) {
	var literatureList common.LiteratureList

	err := json.Unmarshal(data, &literatureList)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 400, common.ErrorMessages[common.WrongJSONInputError])
		return
	}

	postgres := database.Postgres{}
	postgres.Connect()
	defer postgres.Disconnect()

	err = postgres.InsertLiteratureList(literatureList)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 500, common.ErrorMessages[common.InternalServerError])
		return
	}

	answer, err := json.Marshal(Success{"OK"})
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 500, common.ErrorMessages[common.InternalServerError])
		return
	}

	writeAnswer(w, 200, answer)

}

func getLiteratureLists(w http.ResponseWriter) {
	postgres := &database.Postgres{}
	postgres.Connect()
	defer postgres.Disconnect()

	lists, err := postgres.SelectLiteratureList()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 500, common.ErrorMessages[common.InternalServerError])
		return
	}

	data, err := json.Marshal(LiteratureLists{lists})
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 500, common.ErrorMessages[common.InternalServerError])
		return
	}

	writeAnswer(w, 200, data)

}

func getLiteraturePrototype(w http.ResponseWriter) {

	data, err := json.Marshal(common.GetLiteratureExample())
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 500, common.ErrorMessages[common.InternalServerError])
		return
	}

	writeAnswer(w, 200, data)
}

func addLiteratureCheckInput(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	body := r.Body
	answer, err := ioutil.ReadAll(body)
	if err != nil {
		returnError(w, 400, common.ErrorMessages[common.NoJSONProvidedError])
		return nil, err
	}
	return answer, nil
}

func addLiterature(w http.ResponseWriter, data []byte) {
	var literature common.Literature

	err := json.Unmarshal(data, &literature)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 400, common.ErrorMessages[common.WrongJSONInputError])
		return
	}

	postgres := database.Postgres{}
	postgres.Connect()
	defer postgres.Disconnect()

	err = postgres.InsertLiterature(literature)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 500, common.ErrorMessages[common.InternalServerError])
		return
	}

	answer, err := json.Marshal(Success{"OK"})
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 500, common.ErrorMessages[common.InternalServerError])
		return
	}

	writeAnswer(w, 200, answer)
}

func getLiteratureCheckInput(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	if r.ContentLength > 0 {
		body := r.Body
		answer, err := ioutil.ReadAll(body)
		if err != nil {
			returnError(w, 400, common.ErrorMessages[common.NoJSONProvidedError])
			return nil, err
		}
		return answer, nil
	}

	return nil, nil

}

func getLiterature(request []byte, w http.ResponseWriter) {
	postgres := &database.Postgres{}
	postgres.Connect()
	defer postgres.Disconnect()

	if request == nil {
		literature, err := postgres.SelectLiterature()
		if err != nil {
			fmt.Printf("FATAL: %+v\n", err)
			returnError(w, 500, common.ErrorMessages[common.InternalServerError])
			return
		}

		data, err := json.Marshal(Literature{literature})
		if err != nil {
			fmt.Printf("FATAL: %+v\n", err)
			returnError(w, 500, common.ErrorMessages[common.InternalServerError])
			return
		}

		writeAnswer(w, 200, data)

	} else {
		var list common.LiteratureList
		err := json.Unmarshal(request, &list)
		if err != nil {
			fmt.Printf("FATAL: %+v\n", err)
			returnError(w, 500, common.ErrorMessages[common.InternalServerError])
			return
		}
		items, err := postgres.SelectBooksInList(list)
		if err != nil {
			fmt.Printf("FATAL: %+v\n", err)
			returnError(w, 500, common.ErrorMessages[common.InternalServerError])
			return
		}

		data, err := json.Marshal(Books{items})
		if err != nil {
			fmt.Printf("FATAL: %+v\n", err)
			returnError(w, 500, common.ErrorMessages[common.InternalServerError])
			return
		}

		writeAnswer(w, 200, data)
	}
}

func addCourseCheckInput(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	body := r.Body
	answer, err := ioutil.ReadAll(body)
	if err != nil {
		returnError(w, 400, common.ErrorMessages[common.NoJSONProvidedError])
		return nil, err
	}
	return answer, nil
}

func addCourse(w http.ResponseWriter, data []byte) {
	var course common.Course

	err := json.Unmarshal(data, &course)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 400, common.ErrorMessages[common.WrongJSONInputError])
		return
	}

	postgres := database.Postgres{}
	postgres.Connect()
	defer postgres.Disconnect()

	err = postgres.InsertCourse(course)
	if err != nil {
		handleDatabaseErrors(w, err)
		return
	}

	answer, err := json.Marshal(Success{"OK"})
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 500, common.ErrorMessages[common.InternalServerError])
		return
	}

	writeAnswer(w, 200, answer)

}

func getCourses(w http.ResponseWriter) {

	postgres := &database.Postgres{}

	postgres.Connect()
	defer postgres.Disconnect()

	courses, err := postgres.SelectCourses()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 500, common.ErrorMessages[common.InternalServerError])
		return
	}

	data, err := json.Marshal(Courses{courses})
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 500, common.ErrorMessages[common.InternalServerError])
		return
	}

	writeAnswer(w, 200, data)

}

func migrateLiteratureListCheckInput(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	body := r.Body
	answer, err := ioutil.ReadAll(body)
	if err != nil {
		returnError(w, 400, common.ErrorMessages[common.NoJSONProvidedError])
		return nil, err
	}
	return answer, nil
}

func migrateLiteratureList(w http.ResponseWriter, data []byte) {
	log.Println("migrateLiteratureList")
	var migrate common.Migrate

	err := json.Unmarshal(data, &migrate)
	if err != nil {
		log.Println("Error in json umparshalling")
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 400, common.ErrorMessages[common.WrongJSONInputError])
		return
	}

	postgres := database.Postgres{}
	postgres.Connect()
	defer postgres.Disconnect()

	err = postgres.Migrate(migrate)
	if err != nil {
		log.Println("Error in migrate")
		handleDatabaseErrors(w, err)
		return
	}

	answer, err := json.Marshal(Success{"OK"})
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 500, common.ErrorMessages[common.InternalServerError])
		return
	}

	writeAnswer(w, 200, answer)
}

func generateBibTexCheckInput(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	body := r.Body
	answer, err := ioutil.ReadAll(body)
	if err != nil {
		returnError(w, 400, common.ErrorMessages[common.NoJSONProvidedError])
		return nil, err
	}
	return answer, nil
}

func generateBibTex(w http.ResponseWriter, data []byte) {
	var list common.LiteratureList
	err := json.Unmarshal(data, &list)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 400, common.ErrorMessages[common.WrongJSONInputError])
		return
	}

	postgres := database.Postgres{}
	postgres.Connect()
	defer postgres.Disconnect()

	books, err := postgres.SelectBooksInList(list)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		returnError(w, 400, common.ErrorMessages[common.WrongJSONInputError])
		return
	}

	report := reports.CreateReport(books)
	writeAnswer(w, 200, []byte(report))
}

func writeAnswer(w http.ResponseWriter, code int, answer []byte) {
	log.Printf(
		"\n---------\nAnswer:\n\tCode: %d\n\tMessage: %s\n---------\n",
		code,
		string(answer),
	)
	w.WriteHeader(code)
	w.Write(answer)
}

func handleDatabaseErrors(w http.ResponseWriter, err error) {
	log.Println(err)
	full, ok := err.(*common.Error)
	if ok {
		returnError(w, 500, full.GetMessageForUser())
	} else {
		returnError(w, 500, common.ErrorMessages[common.InternalServerError])
	}
}
