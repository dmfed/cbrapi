package cbrapi

import (
	"fmt"
	"testing"
	"time"
)

func Test_dateconverter_string_to_time_converts_as_expected(t *testing.T) {
	fmt.Println("Testing that apidate() converts string to time.Time")
	stringToTime := []struct {
		input        string
		resultstring string
		resulttime   time.Time
	}{
		{input: "16/02/1980",
			resultstring: "16/02/1980",
		},
		{input: "03.04.2007",
			resultstring: "03/04/2007"},
	}
	for _, test := range stringToTime {
		stt := apidate(test.input)
		test.resulttime, _ = time.Parse("02/01/2006", test.resultstring)
		if stt.str != test.resultstring ||
			stt.time != test.resulttime ||
			stt.stringobject() != test.resultstring ||
			stt.timeobject() != test.resulttime {
			fmt.Printf("Resulting string is: %v, expected: %v\nresulting time is: %v, expected %v\n", stt.str, test.resultstring, stt.time, test.resulttime)
			t.Fail()
		}
	}
}

func Test_dateconverter_string_to_time_fails_when_allowed(t *testing.T) {
	fmt.Println("Testing that apidate() fails to convert incorrect string when needed")
	stringToTime := []struct {
		input        string
		resultstring string
		resulttime   time.Time
	}{
		{input: "6/02/1980",
			resultstring: "",
		},
		{input: "01/4/2020",
			resultstring: "",
		},
		{input: "03|04|2007",
			resultstring: ""},
		{input: "03/04|2007",
			resultstring: ""},
		{input: "03/04/207",
			resultstring: ""},
		{input: "03/04/07",
			resultstring: ""},
	}
	for _, test := range stringToTime {
		stt := apidate(test.input)
		test.resulttime, _ = time.Parse("02/01/2006", "")
		if stt.str != test.resultstring ||
			stt.stringobject() != test.resultstring ||
			stt.time != test.resulttime ||
			stt.timeobject() != test.resulttime {
			fmt.Printf("stt.str: %v, stt.time: %v\n", stt.str, stt.time)
			t.Fail()
		}
	}
}
func Test_dateconverter_time_to_string(t *testing.T) {
	fmt.Println("Testing that date converts time.Time to string of needed format")
	timenow := time.Now()
	stt := apidate(timenow)
	if stt.str != timenow.Format("02/01/2006") {
		t.Fail()
	}
	timeToString := []struct {
		input  string
		result string
	}{
		{input: "16/02/1980",
			result: "16/02/1980"},
		{input: "03/04/2007",
			result: "03/04/2007"},
		{input: "03/04/2007",
			result: "03/04/2007"},
	}
	for _, test := range timeToString {
		testtime, _ := time.Parse("02/01/2006", test.input)
		stt := apidate(testtime)
		if stt.time != testtime ||
			stt.timeobject() != testtime ||
			stt.stringobject() != test.result ||
			stt.str != test.result {
			t.Fail()
		}
	}
}

func Test_dateconverter_returns_empty_object(t *testing.T) {
	fmt.Println("Test apidate() returns empty object if wrong interface supplied")
	stt := apidate(56)
	resulttime, _ := time.Parse("02/01/2006", "")
	if stt.str != "" || stt.time != resulttime {
		t.Fail()
	}
}
