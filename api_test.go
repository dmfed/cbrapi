package cbrapi

import (
	"fmt"
	"testing"
	"time"
)

func Test_New(t *testing.T) {
	curr, err := New("USD")
	fmt.Println(curr)
	if err != nil {
		t.Fail()
	}
}

func Test_RateAt(t *testing.T) {
	curr, err := New("USD")
	if err != nil {
		t.Fail()
	}
	fmt.Println(curr.RateAt(time.Now()))
	slice, _ := curr.RateAtRangeDates("10/10/2019", "12/10/2019")
	for _, item := range slice {
		fmt.Println(item)
	}
	curr2, err := New("BEF")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(curr2.RateAt("11/08/2000"))
	slice, _ = curr2.RateAtRangeDates("10/10/2019", "12/10/2019")
	for _, item := range slice {
		fmt.Println(item)
	}
}
