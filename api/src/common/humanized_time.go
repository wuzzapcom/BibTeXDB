package common

import (
	"fmt"
	"strings"
	"time"
)

// HumanizedTime ..
type HumanizedTime struct {
	time.Time
}

//TimeFormat is string that defines format of date for user
const TimeFormat = "2006-01-02"

var nilTime = (time.Time{}).UnixNano()

// UnmarshalJSON ..
func (ht *HumanizedTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ht.Time = time.Time{}
		return nil
	}
	ht.Time, err = time.Parse(TimeFormat, s)
	return err
}

// MarshalJSON ..
func (ht HumanizedTime) MarshalJSON() ([]byte, error) {
	if ht.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ht.Time.Format(TimeFormat))), nil
}

// IsSet ..
func (ht *HumanizedTime) IsSet() bool {
	return ht.Time.UnixNano() != nilTime
}
