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

const TimeFormat = "2006-02-02"

var nilTime = (time.Time{}).UnixNano()

// UnmarshalJSON ..
func (ht *HumanizedTime) UnmarshalJSON(b []byte) (err error) {
	fmt.Println("unmarshal")
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
	fmt.Println("marshal")
	if ht.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ht.Time.Format(TimeFormat))), nil
}

// IsSet ..
func (ht *HumanizedTime) IsSet() bool {
	return ht.Time.UnixNano() != nilTime
}
