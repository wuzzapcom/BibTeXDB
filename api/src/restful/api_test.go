package restful

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
	"wuzzapcom/Coursework/api/src/common"
)

func TestAPI_AddLecturer(t *testing.T) {
	dateOfBirth := common.HumanizedTime{}
	var err error
	dateOfBirth.Time, err = time.Parse(common.TimeFormat, "2013-02-03")
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}
	req, err := http.NewRequest("POST", "/addLecturer", strings.NewReader(common.Lecturer{
		Name:        "Скоробогатов Сергей Юрьевич",
		DateOfBirth: dateOfBirth,
		Department:  "Прикладная математика и информатика",
	}.String()))
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		t.Fail()
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(addLecturerHandler)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	fmt.Println(recorder.Body.String())
}
