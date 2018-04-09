var currentSelectedState = "textbookButton"
const address = "http://localhost:8080/"
var states = new Map() //Maps button ID to Text
states.set("textbookButton", "Textbook")
states.set("literatureButton", "Literature")
states.set("literatureListButton", "LiteratureList")
states.set("courseButton", "Course")
states.set("lecturerButton", "Lecturer")
states.set("departmentButton", "Department")

var urls = new Map() //ID to RequestURL projection
urls.set("textbookButton", ["addBook", "getBooks", "getBookPrototype"])
urls.set("literatureButton", ["addLiterature", "getLiterature", "getLiteraturePrototype"])
urls.set("literatureListButton", ["addLiteratureList", "getLiteratureLists", "getLiteratureListPrototype"])
urls.set("courseButton", ["addCourse", "getCourses", "getCoursePrototype"])
urls.set("lecturerButton", ["addLecturer", "getLecturers", "getLecturerPrototype"])
urls.set("departmentButton", ["addDepartment", "getDepartments", "getDepartmentPrototype"])

function sendDataQuery(id) {
    var x = new XMLHttpRequest()
    x.open("POST", address + urls.get(id)[0], true)
    x.onload = function () {
        if (x.status != 200) {
            var label = document.getElementById("uploadResultLabel")
            label.className = label.className.replace("invisible", "visible")
            label.className += " alert-danger"
            label.innerText = x.status + ': ' + x.statusText

            // alert(x.status + ': ' + x.statusText);
        } else {
            var label = document.getElementById("uploadResultLabel")
            label.className = label.className.replace("invisible", "visible")
            label.className += " alert-success"
            label.innerText = x.responseText
        }
    }
    x.send(document.getElementById("textarea").value)
}

function getDataQuery(id) {
    var x = new XMLHttpRequest()
    x.open("GET", address + urls.get(id)[1], true)
    x.onload = function () {
        if (x.status != 200) {
            alert(x.status + ': ' + x.statusText);
        } else {
            // вывести результат
            document.getElementById("textarea").value = JSON.stringify(JSON.parse(x.responseText), null, 2)
        }
    }
    x.send()
}

function getPrototypeQuery(id) {
    var x = new XMLHttpRequest()
    x.open("GET", address + urls.get(id)[2], true)
    x.onload = function () {
        if (x.status != 200) {
            alert(x.status + ': ' + x.statusText);
        } else {
            // вывести результат
            document.getElementById("textarea").value = JSON.stringify(JSON.parse(x.responseText), null, 2)
        }
    }
    x.send()
}

function addListenersForSettingButtonActiveAndUpdatingTextareaLabel(btnGroupID, textareaLabelID, textareaID) {
    var btnGroup = document.getElementById(btnGroupID)
    var buttons = btnGroup.getElementsByClassName("btn")
    var label = document.getElementById(textareaLabelID)
    var area = document.getElementById(textareaID)
    for (var i = 0; i < buttons.length; i++) {
        buttons[i].addEventListener("click", function () {
            for (var i = 0; i < buttons.length; i++) {
                buttons[i].className = buttons[i].className.replace(" active", "")
            }
            this.className += " active"
            label.textContent = this.textContent
            currentSelectedState = this.id
            area.value = ""
        })
    }
}

function addListenerToUploadButton(btnID) {
    var btn = document.getElementById(btnID)
    btn.addEventListener("click", function () {
        sendDataQuery(currentSelectedState)
    })
}

function addListenerToGetButton(btnID) {
    var btn = document.getElementById(btnID)
    btn.addEventListener("click", function () {
        getDataQuery(currentSelectedState)
    })
}

function addListenerToGetPrototypeButton(btnID) {
    var btn = document.getElementById(btnID)
    btn.addEventListener("click", function () {
        getPrototypeQuery(currentSelectedState)
    })
}

function initTablesButtonGroup(btnGroupID) {
    var btnGroup = document.getElementById(btnGroupID)
    var buttons = btnGroup.getElementsByClassName("btn")
    for (var i = 0; i < buttons.length; i++) {
        buttons[i].textContent = states.get(buttons[i].id)
    }
}