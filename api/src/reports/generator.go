package reports

import (
	"wuzzapcom/Coursework/api/src/common"
)

type report struct {
	books common.Items
}

func (r *report) String() string {
	var result string
	for _, item := range r.books {
		result += item.String() + "\n"
	}
	return result
}

// CreateReport ..
func CreateReport(books common.Items) string {
	r := report{books}
	return r.String()
}
