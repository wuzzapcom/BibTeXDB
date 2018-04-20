class Constants {
    static address: string = "http://localhost:8080/"

    static getSelectURL(forTable: Table) {
        switch (forTable) {
            case Table.Textbook:
                return "getBooks"
            case Table.Course:
                return "getCourses"
            case Table.Department:
                return "getDepartments"
            case Table.Lecturer:
                return "getLecturers"
            case Table.Literature:
                return "getLiterature"
            case Table.LiteratureList:
                return "getLiteratureLists"
        }
    }

    static getTableByInsertButtonID(id: string) {
        switch (id) {
            case "TextbookInsertButtonID":
                return Table.Textbook
            case "CourseInsertButtonID":
                return Table.Course
            case "DepartmentInsertButtonID":
                return Table.Department
            case "LecturerInsertButtonID":
                return Table.Lecturer
            case "LiteratureInsertButtonID":
                return Table.Literature
            case "LiteratureListInsertButtonID":
                return Table.LiteratureList
        }
    }
}

class HTTPWrapper {
    static Get(httpMethod: string, callback: (string) => void) {
        let request = new XMLHttpRequest()
        request.open("GET", Constants.address + httpMethod, true)
        request.onload = function () {
            if (request.status != 200) {
                alert(request.status + " " + request.status + " " + request.statusText)
            } else {
                callback(request.responseText)
            }
        }
        request.onerror = function () {
            alert(request.status + " " + request.status + " " + request.statusText)
        }
        request.send()
    }

    static Post(httpMethod: string, body: string, callback: (string) => void) {
        let request = new XMLHttpRequest()
        request.open("POST", Constants.address + httpMethod, true)
        request.onload = function () {
            if (request.status != 200) {
                alert(request.status + " " + request.status + " " + request.statusText)
            } else {
                callback(request.responseText)
            }
        }
        request.onerror = function () {
            alert(request.status + " " + request.status + " " + request.statusText)
        }
        request.send(body)
    }
}

enum Table {
    Textbook = "Textbook",
    Literature = "Literature",
    LiteratureList = "LiteratureList",
    Course = "Course",
    Lecturer = "Lecturer",
    Department = "Department",
}

class Input {
    static currentState: Table = Table.Textbook
    textareaLabelID: string = "InputTextareaLabel"
    textareaID: string = "InputTextarea"
    getButtonID: string = "GetButton"
    buttonGroupID: string = "GetButtonGroup"

    addListenerOnGetButton() {
        var button = document.getElementById(this.getButtonID)
        var state = Input.currentState
        var textarea = document.getElementById(this.textareaID)
        button.onclick = function () {
            HTTPWrapper.Get(Constants.getSelectURL(state), function (text: string) {
                textarea.innerText = JSON.stringify(JSON.parse(text), null, 2)
            })
        }
    }

    addListenersForSettingButtonActiveAndUpdatingTextareaLabelGet() {
        var buttonGroup = document.getElementById(this.buttonGroupID)
        var buttons = buttonGroup.getElementsByClassName("btn")
        var label = document.getElementById(this.textareaLabelID)
        var area = document.getElementById(this.textareaID)
        for (var i = 0; i < buttons.length; i++) {
            buttons[i].addEventListener("click", function () {
                for (var i = 0; i < buttons.length; i++) {
                    buttons[i].className = buttons[i].className.replace(" active", "")
                }
                this.className += " active"
                label.textContent = this.textContent
                Input.currentState = Constants.getTableByInsertButtonID(this.id)
                area.innerText = ""
            })
        }
    }

    initTablesButtonGroup() {
        let end = "InsertButtonID"
        document.getElementById(Table.Course + end).textContent = Table.Course
        document.getElementById(Table.Department + end).textContent = Table.Department
        document.getElementById(Table.Textbook + end).textContent = Table.Textbook
        document.getElementById(Table.Literature + end).textContent = Table.Literature
        document.getElementById(Table.LiteratureList + end).textContent = Table.LiteratureList
        document.getElementById(Table.Lecturer + end).textContent = Table.Lecturer
    }
}

var input = new Input()
input.initTablesButtonGroup()
