var currentStateGet = "textbookButtonGet"

var statesGet = new Map() //Maps button ID to Text
statesGet.set("textbookButtonGet", "Textbook")
statesGet.set("literatureButtonGet", "Literature")
statesGet.set("literatureListButtonGet", "LiteratureList")
statesGet.set("courseButtonGet", "Course")
statesGet.set("lecturerButtonGet", "Lecturer")
statesGet.set("departmentButtonGet", "Department")

var getUrls = new Map() //ID to RequestURL projection
getUrls.set("textbookButtonGet", "getBooks")
getUrls.set("literatureButtonGet", "getLiterature")
getUrls.set("literatureListButtonGet", "getLiteratureLists")
getUrls.set("courseButtonGet", "getCourses")
getUrls.set("lecturerButtonGet", "getLecturers")
getUrls.set("departmentButtonGet", "getDepartments")


function getDataQuery(id) {
    var x = new XMLHttpRequest()
    x.open("GET", address + getUrls.get(id), true)
    x.onload = function () {
        if (x.status != 200) {
            alert(x.status + ': ' + x.statusText);
        } else {
            // вывести результат
            document.getElementById("textareaGet").value = JSON.stringify(JSON.parse(x.responseText), null, 2)
        }
    }
    x.send()
}

function addListenerToGetButton(btnID) {
    var btn = document.getElementById(btnID)
    btn.addEventListener("click", function () {
        getDataQuery(currentStateGet)
    })
}

function addListenersForSettingButtonActiveAndUpdatingTextareaLabelGet(btnGroupID, textareaLabelID, textareaID) {
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
            currentStateGet = this.id
            area.value = ""
        })
    }
}