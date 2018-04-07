function setRed(element) {
    element.style.color = 'red'
}

function getText(str) {
    return str
}

function setText(id) {
    document.getElementById(id).innerHTML = getText("str")
}

function testQuery() {
    var x = new XMLHttpRequest();
    x.open("GET", "https://jsonplaceholder.typicode.com/posts", true);
    x.onload = function () {
        document.getElementById("comment").value = x.responseText

    }
    x.send(null);
}

function setButtonActiveAndUpdateTextareaLabel(button, btnGroupID) {
    var btnGroup = document.getElementById(btnGroupID)
    var buttons = btnGroup.getElementsByClassName("btn")
    for (var i = 0; i < buttons.length; i++) {
        buttons[i].className = buttons[i].className.replace(" active", "")
    }
    button.className += " active"
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
        })
    }
}