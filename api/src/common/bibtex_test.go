package common_test

import (
	"testing"
	"wuzzapcom/Coursework/api/src/common"
	"fmt"
)

func TestItems_Append(t *testing.T) {
	items := common.GetRandomItems()

	err := items.Append(common.GetRandomItems())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	err = items.Append(common.GetRandomItems()[0])
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	err = items.Append(&common.GetRandomItems()[0])
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
