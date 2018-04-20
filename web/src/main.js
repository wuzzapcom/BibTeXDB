var Constants = /** @class */ (function () {
    function Constants() {
    }
    Constants.getSelectURL = function (forTable) {
        switch (forTable) {
            case Table.Textbook:
                return "getBooks";
            case Table.Course:
                return "getCourses";
            case Table.Department:
                return "getDepartments";
            case Table.Lecturer:
                return "getLecturers";
            case Table.Literature:
                return "getLiterature";
            case Table.LiteratureList:
                return "getLiteratureLists";
        }
    };
    Constants.getTableByInsertButtonID = function (id) {
        switch (id) {
            case "TextbookInsertButtonID":
                return Table.Textbook;
            case "CourseInsertButtonID":
                return Table.Course;
            case "DepartmentInsertButtonID":
                return Table.Department;
            case "LecturerInsertButtonID":
                return Table.Lecturer;
            case "LiteratureInsertButtonID":
                return Table.Literature;
            case "LiteratureListInsertButtonID":
                return Table.LiteratureList;
        }
    };
    Constants.address = "http://localhost:8080/";
    return Constants;
}());
var HTTPWrapper = /** @class */ (function () {
    function HTTPWrapper() {
    }
    HTTPWrapper.Get = function (httpMethod, callback) {
        var request = new XMLHttpRequest();
        request.open("GET", Constants.address + httpMethod, true);
        request.onload = function () {
            if (request.status != 200) {
                alert(request.status + " " + request.status + " " + request.statusText);
            }
            else {
                callback(request.responseText);
            }
        };
        request.onerror = function () {
            alert(request.status + " " + request.status + " " + request.statusText);
        };
        request.send();
    };
    HTTPWrapper.Post = function (httpMethod, body, callback) {
        var request = new XMLHttpRequest();
        request.open("POST", Constants.address + httpMethod, true);
        request.onload = function () {
            if (request.status != 200) {
                alert(request.status + " " + request.status + " " + request.statusText);
            }
            else {
                callback(request.responseText);
            }
        };
        request.onerror = function () {
            alert(request.status + " " + request.status + " " + request.statusText);
        };
        request.send(body);
    };
    return HTTPWrapper;
}());
var Table;
(function (Table) {
    Table["Textbook"] = "Textbook";
    Table["Literature"] = "Literature";
    Table["LiteratureList"] = "LiteratureList";
    Table["Course"] = "Course";
    Table["Lecturer"] = "Lecturer";
    Table["Department"] = "Department";
})(Table || (Table = {}));
var Input = /** @class */ (function () {
    function Input() {
        this.textareaLabelID = "InputTextareaLabel";
        this.textareaID = "InputTextarea";
        this.getButtonID = "GetButton";
        this.buttonGroupID = "GetButtonGroup";
    }
    Input.prototype.addListenerOnGetButton = function () {
        var button = document.getElementById(this.getButtonID);
        var state = Input.currentState;
        var textarea = document.getElementById(this.textareaID);
        button.onclick = function () {
            HTTPWrapper.Get(Constants.getSelectURL(state), function (text) {
                textarea.innerText = JSON.stringify(JSON.parse(text), null, 2);
            });
        };
    };
    Input.prototype.addListenersForSettingButtonActiveAndUpdatingTextareaLabelGet = function () {
        var buttonGroup = document.getElementById(this.buttonGroupID);
        var buttons = buttonGroup.getElementsByClassName("btn");
        var label = document.getElementById(this.textareaLabelID);
        var area = document.getElementById(this.textareaID);
        for (var i = 0; i < buttons.length; i++) {
            buttons[i].addEventListener("click", function () {
                for (var i = 0; i < buttons.length; i++) {
                    buttons[i].className = buttons[i].className.replace(" active", "");
                }
                this.className += " active";
                label.textContent = this.textContent;
                Input.currentState = Constants.getTableByInsertButtonID(this.id);
                area.innerText = "";
            });
        }
    };
    Input.prototype.initTablesButtonGroup = function () {
        var end = "InsertButtonID";
        document.getElementById(Table.Course + end).textContent = Table.Course;
        document.getElementById(Table.Department + end).textContent = Table.Department;
        document.getElementById(Table.Textbook + end).textContent = Table.Textbook;
        document.getElementById(Table.Literature + end).textContent = Table.Literature;
        document.getElementById(Table.LiteratureList + end).textContent = Table.LiteratureList;
        document.getElementById(Table.Lecturer + end).textContent = Table.Lecturer;
    };
    Input.currentState = Table.Textbook;
    return Input;
}());
var input = new Input();
input.initTablesButtonGroup();
