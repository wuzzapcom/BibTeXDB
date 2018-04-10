// var currentSelectedState = "textbookButton"
const address = "http://localhost:8080/"



function initTablesButtonGroup(btnGroupID, states) {
    var btnGroup = document.getElementById(btnGroupID)
    var buttons = btnGroup.getElementsByClassName("btn")
    for (var i = 0; i < buttons.length; i++) {
        buttons[i].textContent = states.get(buttons[i].id)
    }
}