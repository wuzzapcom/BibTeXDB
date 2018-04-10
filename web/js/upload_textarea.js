var currentStateUpload = "textbookButtonUpload"

var statesUpload = new Map() //Maps button ID to Text
statesUpload.set("textbookButtonUpload", "Textbook")
statesUpload.set("literatureButtonUpload", "Literature")
statesUpload.set("literatureListButtonUpload", "LiteratureList")
statesUpload.set("courseButtonUpload", "Course")
statesUpload.set("lecturerButtonUpload", "Lecturer")
statesUpload.set("departmentButtonUpload", "Department")

var uploadUrls = new Map() //ID to RequestURL projection
uploadUrls.set("textbookButtonUpload", ["addBook", "getBookPrototype"])
uploadUrls.set("literatureButtonUpload", ["addLiterature", "getLiteraturePrototype"])
uploadUrls.set("literatureListButtonUpload", ["addLiteratureList", "getLiteratureListPrototype"])
uploadUrls.set("courseButtonUpload", ["addCourse", "getCoursePrototype"])
uploadUrls.set("lecturerButtonUpload", ["addLecturer", "getLecturerPrototype"])
uploadUrls.set("departmentButtonUpload", ["addDepartment", "getDepartmentPrototype"])

function getPrototypeQuery(id) {
    alert(currentStateUpload)
    var x = new XMLHttpRequest()
    x.open("GET", address + uploadUrls.get(id)[1], true)
    x.onload = function () {
        if (x.status != 200) {
            alert(x.status + ': ' + x.statusText);
        } else {
            // вывести результат
            document.getElementById("textareaUpload").value = JSON.stringify(JSON.parse(x.responseText), null, 2)
        }
    }
    x.send()
}

function addListenerToGetPrototypeButton(btnID) {
    var btn = document.getElementById(btnID)
    btn.addEventListener("click", function () {
        getPrototypeQuery(currentStateUpload)
    })
}

function sendDataQuery(id) {
    var x = new XMLHttpRequest()
    x.open("POST", address + uploadUrls.get(id)[0], true)
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
    x.send(document.getElementById("textareaUpload").value)
}

function addListenerToUploadButton(btnID) {
    var btn = document.getElementById(btnID)
    btn.addEventListener("click", function () {
        sendDataQuery(currentStateUpload)
    })
}

function addListenersForSettingButtonActiveAndUpdatingTextareaLabelUpload(btnGroupID, textareaLabelID, textareaID) {
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
            currentStateUpload = this.id
            alert("update state with " + this.id)
            // currentSelectedState = this.id
            area.value = ""
        })
    }
}