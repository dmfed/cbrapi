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

func Test_RateAtDate(t *testing.T) {
	curr, err := New("USD")
	if err != nil {
		t.Fail()
	}
	fmt.Println(curr.RateAtDate(time.Now()))
	slice, _ := curr.RateAtRangeDates("10/10/2019", "12/10/2019")
	for _, item := range slice {
		fmt.Println(item)
	}
	curr2, err := New("BEF")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(curr2.RateAtDate("11/08/2000"))
	slice, _ = curr2.RateAtRangeDates("10/10/2019", "12/10/2019")
	for _, item := range slice {
		fmt.Println(item)
	}
}

func Test_Quote_functions(t *testing.T) {
	if rate, err := QuoteAtDate("USD", "01/09/2020"); err != nil {
		t.Fail()
	} else {
		fmt.Println(rate)
	}
	if rate, err := QuoteAtRangeDates("USD", "25/11/2020", time.Now()); err != nil {
		t.Fail()
	} else {
		for _, r := range rate {
			fmt.Println(r)
		}
	}
}
