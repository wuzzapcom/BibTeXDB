package fetchers

//BibTexItem ..
type BibTexItem struct {
	Ident     string
	Title     string
	Author    string
	Publisher string
	Year      string
	Language  string
	ISBN      string
	URL       string
}

func (b BibTexItem) String() string {
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
