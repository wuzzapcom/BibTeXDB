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
        })
    }
}

function addListenerToGetPrototypeButton(btnID, textareaLabelID, textareaID) {
    var btn = document.getElementById(btnID)
    var label = document.getElementById(textareaLabelID)
    var area = document.getElementById(textareaID)
    btn.addEventListener("click", function () {
        area.value = label.textContent
        testQuery()
    })
}