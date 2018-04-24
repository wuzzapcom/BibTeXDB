import * as FileSaver from "./dist/node_modules/file-saver/FileSaver"

// var fileSaver = require('file-saver')

export class Constants {
    static address: string = "http://localhost:8080/"
    static searchURL: string = "search?request="

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

    static getUploadURL(forTable: Table) {
        switch (forTable) {
            case Table.Textbook:
                return "addBook"
            case Table.Course:
                return "addCourse"
            case Table.Department:
                return "addDepartment"
            case Table.Lecturer:
                return "addLecturer"
            case Table.Literature:
                return "addLiterature"
            case Table.LiteratureList:
                return "addLiteratureList"
            case Table.Migrate:
                return "migrateLiteratureList"
        }
    }

    static getPrototypeURL(forTable: Table) {
        switch (forTable) {
            case Table.Textbook:
                return "getBookPrototype"
            case Table.Course:
                return "getCoursePrototype"
            case Table.Department:
                return "getDepartmentPrototype"
            case Table.Lecturer:
                return "getLecturerPrototype"
            case Table.Literature:
                return "getLiteraturePrototype"
            case Table.LiteratureList:
                return "getLiteratureListPrototype"
            case Table.Migrate:
                return "getMigratePrototype"
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
            default:
                alert("Unknown id in getTableByUploadButtonID " + id)
        }
    }

    static getTableByUploadButtonID(id: string) {
        switch (id) {
            case "TextbookUploadButtonID":
                return Table.Textbook
            case "CourseUploadButtonID":
                return Table.Course
            case "DepartmentUploadButtonID":
                return Table.Department
            case "LecturerUploadButtonID":
                return Table.Lecturer
            case "LiteratureUploadButtonID":
                return Table.Literature
            case "LiteratureListUploadButtonID":
                return Table.LiteratureList
            case "MigrateUploadButtonID":
                return Table.Migrate
            default:
                alert("Unknown id in getTableByUploadButtonID " + id)
        }
    }

    static saveFile(text: string) {
        var file = new File([text], "report.txt", { type: "text/plain;charset=utf-8" });
        FileSaver.saveAs(file)
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
    Migrate = "Migrate",
}

class Input {
    static currentState: Table = Table.Textbook
    textareaLabelID: string = "InputTextareaLabel"
    textareaID: string = "InputTextarea"
    getButtonID: string = "GetButton"
    buttonGroupID: string = "GetButtonGroup"

    addListenerOnGetButton() {
        var button = document.getElementById(this.getButtonID)
        var textarea = <HTMLTextAreaElement>document.getElementById(this.textareaID)
        button.onclick = function () {
            HTTPWrapper.Get(Constants.getSelectURL(Input.currentState), function (text: string) {
                textarea.value = JSON.stringify(JSON.parse(text), null, 2)
            })
        }
    }

    addListenersForSettingButtonActiveAndUpdatingTextareaLabelGet() {
        var buttonGroup = document.getElementById(this.buttonGroupID)
        var buttons = buttonGroup.getElementsByClassName("btn")
        var label = document.getElementById(this.textareaLabelID)
        var area = <HTMLTextAreaElement>document.getElementById(this.textareaID)
        for (var i = 0; i < buttons.length; i++) {
            buttons[i].addEventListener("click", function () {
                for (var i = 0; i < buttons.length; i++) {
                    buttons[i].className = buttons[i].className.replace(" active", "")
                }
                this.className += " active"
                label.textContent = this.textContent
                Input.currentState = Constants.getTableByInsertButtonID(this.id)
                area.value = ""
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

class Output {
    static currentState: Table = Table.Textbook
    textareaLabelID: string = "OutputTextareaLabel"
    textareaID: string = "OutputTextarea"
    uploadButtonID: string = "OutputButton"
    buttonGroupID: string = "OutputButtonGroup"

    addListenerOnUploadButton() {
        var button = document.getElementById(this.uploadButtonID)
        var textarea = <HTMLTextAreaElement>document.getElementById(this.textareaID)
        button.onclick = function () {
            console.log("uploadButton: " + textarea.value)
            HTTPWrapper.Post(Constants.getUploadURL(Output.currentState), textarea.value, function (text: string) {
                alert(text)
            })
        }
    }

    addListenersForSettingButtonActiveAndUpdatingTextareaLabelUpload() {
        var buttonGroup = document.getElementById(this.buttonGroupID)
        var buttons = buttonGroup.getElementsByClassName("btn")
        var label = document.getElementById(this.textareaLabelID)
        var area = <HTMLTextAreaElement>document.getElementById(this.textareaID)
        for (var i = 0; i < buttons.length; i++) {
            buttons[i].addEventListener("click", function () {
                for (var i = 0; i < buttons.length; i++) {
                    buttons[i].className = buttons[i].className.replace(" active", "")
                }
                this.className += " active"
                label.textContent = this.textContent
                Output.currentState = Constants.getTableByUploadButtonID(this.id)
                area.value = ""
                if (Output.currentState == Table.Textbook) {
                    return
                }
                HTTPWrapper.Get(Constants.getPrototypeURL(Output.currentState), function (text: string) {
                    area.value = JSON.stringify(JSON.parse(text), null, 2)
                })
            })
        }
    }

    initTablesButtonGroup() {
        let end = "UploadButtonID"
        document.getElementById(Table.Course + end).textContent = Table.Course
        document.getElementById(Table.Department + end).textContent = Table.Department
        document.getElementById(Table.Textbook + end).textContent = Table.Textbook
        document.getElementById(Table.Literature + end).textContent = Table.Literature
        document.getElementById(Table.LiteratureList + end).textContent = Table.LiteratureList
        document.getElementById(Table.Lecturer + end).textContent = Table.Lecturer
        document.getElementById(Table.Migrate + end).textContent = Table.Migrate
    }
}

class Search {

    static searchLabelID: string = "SearchLabel"
    static requestInputID: string = "RequestInput"
    static submitButtonID: string = "SubmitButton"
    static searchTextareaID: string = "SearchTextarea"

    addHandlerOnSubmitButton() {
        var submit = document.getElementById(Search.submitButtonID)
        submit.addEventListener("click", function () {
            var req = <HTMLInputElement>document.getElementById(Search.requestInputID)
            var textarea = document.getElementById(Search.searchTextareaID)
            textarea.innerText = req.value
            HTTPWrapper.Get(Constants.searchURL + encodeURIComponent(req.value), function (text: string) {
                textarea.innerText = JSON.stringify(JSON.parse(text), null, 2)
            })
        })
    }
}

function initMain() {
    var input = new Input()
    input.initTablesButtonGroup()
    input.addListenersForSettingButtonActiveAndUpdatingTextareaLabelGet()
    input.addListenerOnGetButton()

    var output = new Output()
    output.initTablesButtonGroup()
    output.addListenerOnUploadButton()
    output.addListenersForSettingButtonActiveAndUpdatingTextareaLabelUpload()

    var search = new Search()
    search.addHandlerOnSubmitButton()
}

class Report {
    listID: string = "ll"
    fields: string[]

    createElement(withID: string, withText: string) {
        var list = document.getElementById(this.listID)

        var l = document.createElement("a")
        var t = document.createTextNode(withText)
        l.appendChild(t)
        l.id = withID
        l.className = "list-group-item list-group-item-action"

        var report = this

        l.addEventListener("click", function () {
            console.log(report.fields[Number(this.id)].toString())
            HTTPWrapper.Post("generateReport", report.fields[Number(this.id)].toString(), function (text: string) {
                Constants.saveFile(text)
            })
        })

        list.appendChild(l)
    }

}

function initReport() {
    var report = new Report()

    HTTPWrapper.Get(Constants.getSelectURL(Table.LiteratureList), function (text: string) {
        let obj = JSON.parse(text)
        var i: number = 0
        report.fields = []

        obj.lists.forEach(elem => {
            report.createElement(i.toString(), `Title: ${elem.CourseTitle}, Semester: ${elem.Semester}, Year: ${elem.Year}, Department: ${elem.DepartmentTitle}`)
            report.fields[report.fields.length] = JSON.stringify(elem)
            i++
        });
    })
}

/*
    Костыль, который позволяет вызывать методы из HTML. Как сделать лучше -- неизвестно.
*/
window['initMain'] = function () {
    initMain()
}

window['initReport'] = function () {
    initReport()
}
