var currentSelectedState = "Textbook"
var states = new Map() //Maps button ID to Text
states.set("textbookButton", "Textbook")
states.set("literatureButton", "Literature")
states.set("literatureListButton", "LiteratureList")
states.set("courseButton", "Course")
states.set("lecturerButton", "Lecturer")
states.set("departmentButton", "Department")


function testQuery() {
    var x = new XMLHttpRequest();
    x.open("GET", "https://jsonplaceholder.typicode.com/posts", true);
    x.onload = function () {
        document.getElementById("textarea").value += "\n" + x.responseText

    }
    x.send(null);
}

function addListenersForSettingButtonActiveAndUpdatingTextareaLabel(btnGroupID, textareaID) {
    var btnGroup = document.getElementById(btnGroupID)
    var buttons = btnGroup.getElementsByClassName("btn")
    var label = document.getElementById(textareaID)
    for (var i = 0; i < buttons.length; i++) {
        buttons[i].addEventListener("click", function () {
            for (var i = 0; i < buttons.length; i++) {
                buttons[i].className = buttons[i].className.replace(" active", "")
            }
            this.className += " active"
            label.textContent = this.textContent
            currentSelectedState = label.textContent
        })
    }
}

function addListenerToGetPrototypeButton(btnID, textareaID) {
    var btn = document.getElementById(btnID)
    var area = document.getElementById(textareaID)
    btn.addEventListener("click", function () {
        area.value = currentSelectedState
        testQuery()
    })
}

function initTablesButtonGroup(btnGroupID) {
    var btnGroup = document.getElementById(btnGroupID)
    var buttons = btnGroup.getElementsByClassName("btn")
    for (var i = 0; i < buttons.length; i++) {
        buttons[i].textContent = states.get(buttons[i].id)
    }
}