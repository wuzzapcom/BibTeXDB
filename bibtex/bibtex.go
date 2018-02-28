package bibtex

//Items type is alias for []Item
type Items []Item

func (items Items) String() string {

	result := ""

	for _, val := range items {
		result += val.String()
		result += "\n"
	}

	return result
}

//Item represents legal BibTeX object. Author may contain many authors with \"and\" separator
type Item struct {
	Ident     string
	Title     string
	Author    string
	Publisher string
	Year      string
	Language  string
	ISBN      string
	URL       string
}

func (b Item) String() string {
	return "@Book{" + b.Ident + ",\n" +
		"\ttitle = {" + b.Title + "},\n" +
		"\tauthor = {" + b.Author + "},\n" +
		"\tpublisher = {" + b.Publisher + "},\n" +
		"\tyear = " + b.Year + ",\n" +
		"\tlanguage = " + b.Language + ",\n" +
		"\turl = " + b.URL + ",\n" +
		"\tisbn = " + b.ISBN + "\n" +
		"}"
}

//GetRandomItems returns few filled and real BibTeX structures
func GetRandomItems() Items {

	/*
		Item{
			Ident:     "",
			Title:     "",
			Author:    "",
			Publisher: "",
			Year:      "",
			Language:  "",
			ISBN:      "",
			URL:       "",
		},
	*/

	return []Item{
		Item{
			Ident:     "FlowersForAlgernon",
			Title:     "Flowers for Algernon",
			Author:    "D.D. Keyes",
			Publisher: "Houghton Mifflin Harcourt",
			Year:      "2007",
			Language:  "English",
			ISBN:      "9780547539638",
			URL:       "https://www.googleapis.com/books/v1/volumes/_oG_iTxP1pIC",
		},
		Item{
			Ident:     "Difuri",
			Title:     "Дифференциальное исчисление функций многих переменных",
			Author:    "А.Н. Канатников and А.П. Крищенко and В.Н. Четвериков",
			Publisher: "Изд. МГТУ им. Н. Э. Баумана",
			Year:      "2000",
			Language:  "Russian",
			ISBN:      "9785703816820",
			URL:       "https://www.googleapis.com/books/v1/volumes/DG9wAAAACAAJ",
		},
		Item{
			Ident:     "GolangTextbook",
			Title:     "Программирование на Go. Разработка приложений XXI века",
			Author:    "Марк Саммерфильд",
			Publisher: "Litres",
			Year:      "2017",
			Language:  "Russian",
			ISBN:      "9785457427532",
			URL:       "https://www.googleapis.com/books/v1/volumes/a279DQAAQBAJ",
		},
		Item{
			Ident:     "RustTextbookOld",
			Title:     "Rust",
			Author:    "Jonathan Waldman",
			Publisher: "Simon and Schuster",
			Year:      "2016",
			Language:  "English",
			ISBN:      "9781451691603",
			URL:       "https://www.googleapis.com/books/v1/volumes/Xoe_CwAAQBAJ",
		},
		Item{
			Ident:     "RustTextbookNew",
			Title:     "Rust",
			Author:    "S. B. Joseph",
			Publisher: "Ward Hill Press",
			Year:      "1989",
			Language:  "English",
			ISBN:      "9780962338045",
			URL:       "https://www.googleapis.com/books/v1/volumes/lgwfI0GNAHkC",
		},
	}

}
