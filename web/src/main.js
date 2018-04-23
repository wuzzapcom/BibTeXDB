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
    Constants.getUploadURL = function (forTable) {
        switch (forTable) {
            case Table.Textbook:
                return "addBook";
            case Table.Course:
                return "addCourse";
            case Table.Department:
                return "addDepartment";
            case Table.Lecturer:
                return "addLecturer";
            case Table.Literature:
                return "addLiterature";
            case Table.LiteratureList:
                return "addLiteratureList";
        }
    };
    Constants.getPrototypeURL = function (forTable) {
        switch (forTable) {
            case Table.Textbook:
                return "getBookPrototype";
            case Table.Course:
                return "getCoursePrototype";
            case Table.Department:
                return "getDepartmentPrototype";
            case Table.Lecturer:
                return "getLecturerPrototype";
            case Table.Literature:
                return "getLiteraturePrototype";
            case Table.LiteratureList:
                return "getLiteratureListPrototype";
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
            default:
                alert("Unknown id in getTableByUploadButtonID " + id);
        }
    };
    Constants.getTableByUploadButtonID = function (id) {
        switch (id) {
            case "TextbookUploadButtonID":
                return Table.Textbook;
            case "CourseUploadButtonID":
                return Table.Course;
            case "DepartmentUploadButtonID":
                return Table.Department;
            case "LecturerUploadButtonID":
                return Table.Lecturer;
            case "LiteratureUploadButtonID":
                return Table.Literature;
            case "LiteratureListUploadButtonID":
                return Table.LiteratureList;
            default:
                alert("Unknown id in getTableByUploadButtonID " + id);
        }
    };
    Constants.saveFile = function (text) {
        var filename = "reports.txt";
        var filetype = "text/plain";
        var a = document.createElement("a");
        var dataURI = "data:" + filetype +
            ";base64," + btoa(text);
        a.href = dataURI;
        a['download'] = filename;
        var e = document.createEvent("MouseEvents");
        // Use of deprecated function to satisfy TypeScript.
        e.initMouseEvent("click", true, false, document.defaultView, 0, 0, 0, 0, 0, false, false, false, false, 0, null);
        a.dispatchEvent(e);
        a.remove();
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
        var textarea = document.getElementById(this.textareaID);
        button.onclick = function () {
            HTTPWrapper.Get(Constants.getSelectURL(Input.currentState), function (text) {
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
var Output = /** @class */ (function () {
    function Output() {
        this.textareaLabelID = "OutputTextareaLabel";
        this.textareaID = "OutputTextarea";
        this.uploadButtonID = "OutputButton";
        this.buttonGroupID = "OutputButtonGroup";
    }
    Output.prototype.addListenerOnUploadButton = function () {
        var button = document.getElementById(this.uploadButtonID);
        var state = Input.currentState;
        var textarea = document.getElementById(this.textareaID);
        button.onclick = function () {
            HTTPWrapper.Post(Constants.getUploadURL(state), textarea.textContent, function (text) {
                alert(text);
            });
        };
    };
    Output.prototype.addListenersForSettingButtonActiveAndUpdatingTextareaLabelUpload = function () {
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
                Output.currentState = Constants.getTableByUploadButtonID(this.id);
                area.innerText = "";
                HTTPWrapper.Get(Constants.getPrototypeURL(Output.currentState), function (text) {
                    area.innerText = JSON.stringify(JSON.parse(text), null, 2);
                });
            });
        }
    };
    Output.prototype.initTablesButtonGroup = function () {
        var end = "UploadButtonID";
        document.getElementById(Table.Course + end).textContent = Table.Course;
        document.getElementById(Table.Department + end).textContent = Table.Department;
        document.getElementById(Table.Textbook + end).textContent = Table.Textbook;
        document.getElementById(Table.Literature + end).textContent = Table.Literature;
        document.getElementById(Table.LiteratureList + end).textContent = Table.LiteratureList;
        document.getElementById(Table.Lecturer + end).textContent = Table.Lecturer;
    };
    Output.currentState = Table.Textbook;
    return Output;
}());
function initMain() {
    var input = new Input();
    input.initTablesButtonGroup();
    input.addListenersForSettingButtonActiveAndUpdatingTextareaLabelGet();
    input.addListenerOnGetButton();
    var output = new Output();
    output.initTablesButtonGroup();
    output.addListenerOnUploadButton();
    output.addListenersForSettingButtonActiveAndUpdatingTextareaLabelUpload();
}
var Report = /** @class */ (function () {
    function Report() {
        this.listID = "ll";
    }
    Report.prototype.createElement = function (withID, withText) {
        var list = document.getElementById(this.listID);
        var l = document.createElement("a");
        var t = document.createTextNode(withText);
        l.appendChild(t);
        l.id = withID;
        l.className = "list-group-item list-group-item-action";
        var report = this;
        l.addEventListener("click", function () {
            console.log(report.fields[Number(this.id)].toString());
            HTTPWrapper.Post("generateReport", report.fields[Number(this.id)].toString(), function (text) {
                Constants.saveFile(text);
            });
        });
        list.appendChild(l);
    };
    return Report;
}());
function initReport() {
    var report = new Report();
    HTTPWrapper.Get(Constants.getSelectURL(Table.LiteratureList), function (text) {
        var obj = JSON.parse(text);
        var i = 0;
        report.fields = [];
        obj.lists.forEach(function (elem) {
            report.createElement(i.toString(), "Title: " + elem.CourseTitle + ", Semester: " + elem.Semester + ", Year: " + elem.Year + ", Department: " + elem.DepartmentTitle);
            report.fields[report.fields.length] = JSON.stringify(elem);
            i++;
        });
    });
}
