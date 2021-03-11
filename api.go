// Package cbrapi implements basic functionality of API
// of Central Bank of Russia (cbr.ru). It allows to quote
// exchange rates of currencies in very simple steps
package cbrapi

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html/charset"
)

var (
	// ErrIncorrectCode is returned if currency code is not supported by the API
	ErrIncorrectCode = errors.New("Provided currency code is not known to the API")
	// ErrNoData is returned if rate for requested date(s) was not found
	ErrNoData = errors.New("API did not provide requested data")
)

// Currency represents a currency known to Central Bank of Russia API
// Fields of this struct are almost identical to what API returns
type Currency struct {
	NameRUS     string
	NameENG     string
	APIID       string
	Nominal     int
	ISONumCode  int
	ISOCharCode string
}

// ExchangeRate represents exchange rate of a currency at specified date.
type ExchangeRate struct {
	ISOCode string
	Nominal int
	Date    time.Time
	Rate    float64
}

func (r ExchangeRate) String() string {
	switch r.Nominal {
	case 1:
		return fmt.Sprintf("%v %v/RUB %v", r.Date.Format("02/01/2006"), r.ISOCode, r.Rate)
	default:
		return fmt.Sprintf("%v %v%v/RUB %v", r.Date.Format("02/01/2006"), r.Nominal, r.ISOCode, r.Rate)
	}
}

// New returns instance of Currency object which can be used
// to request exchange rate of currency from the Central Bank API
// with RateAtDate() and RateAtRangeDates() methods or
// QuoteAtDate() and QuoteAtRangeDates()
// In case of error New returns nil. Make sure to always check error.
func New(ISOcode string) (*Currency, error) {
	ISOcode = strings.ToUpper(ISOcode)
	if globalVarAPICodes == nil {
		if err := initAPI(); err != nil {
			return nil, err
		}
	}
	item, ok := globalVarAPICodes[ISOcode]
	if !ok {
		return nil, ErrIncorrectCode
	}
	return &item, nil
}

// RateAtDate accept either "DD/MM/YYYY" formatted date or
// time.Time object. It sends request to the API and returns
// ExchangeRate object
func (c *Currency) RateAtDate(date interface{}) (ExchangeRate, error) {
	rate := ExchangeRate{}
	url := fmt.Sprintf(EndpointSingleDate, apidate(date).stringobject())
	resp, err := http.Get(url)
	if err != nil {
		return rate, err
	}
	defer resp.Body.Close()
	var daily ResponseDaily
	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReaderLabel
	if err := decoder.Decode(&daily); err != nil {
		return rate, err
	}
	for _, item := range daily.Elements {
		if item.APIID == c.APIID {
			r, _ := strconv.ParseFloat(strings.ReplaceAll(item.Value, ",", "."), 64)
			rate.ISOCode = c.ISOCharCode
			rate.Nominal = c.Nominal
			rate.Date = apidate(date).timeobject()
			rate.Rate = r
			break
		}
	}
	return rate, ErrNoData
}

// RateAtRangeDates accept either "DD/MM/YYYY" formatted date or
// time.Time object. It sends request to the API and returns
// slice of ExchangeRate objects
func (c *Currency) RateAtRangeDates(startdate, enddate interface{}) ([]ExchangeRate, error) {
	rates := []ExchangeRate{}
	url := fmt.Sprintf(EndpointDateRange, apidate(startdate).stringobject(), apidate(enddate).stringobject(), c.APIID)
	resp, err := http.Get(url)
	if err != nil {
		return rates, err
	}
	var result ResponseRange
	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReaderLabel
	if err := decoder.Decode(&result); err != nil {
		return rates, err
	}
	for _, item := range result.Elements {
		rate, _ := strconv.ParseFloat(strings.ReplaceAll(item.Value, ",", "."), 64)
		date := apidate(item.Date).timeobject()
		r := ExchangeRate{ISOCode: c.ISOCharCode, Nominal: c.Nominal, Date: date, Rate: rate}
		rates = append(rates, r)
	}
	switch len(rates) {
	case 0:
		return rates, ErrNoData
	default:
		return rates, nil
	}
}

// QuoteAtDate accepts ISO code of currency and either "DD/MM/YYYY" formatted date or
// time.Time object and returns ExchangeRate object
func QuoteAtDate(ISOCode string, date interface{}) (ExchangeRate, error) {
	c, err := New(ISOCode)
	if err != nil {
		return ExchangeRate{}, err
	}
	rate, err := c.RateAtDate(date)
	c = nil // cleanup
	return rate, err
}

// QuoteAtRangeDates accepts ISO code of currency and either "DD/MM/YYYY" formatted dates or
// time.Time objects where startdate and enddate are both inclusive and returns slice of
// ExchangeRate objects for corresponding dates
func QuoteAtRangeDates(ISOCode string, startdate, enddate interface{}) ([]ExchangeRate, error) {
	c, err := New(ISOCode)
	if err != nil {
		return []ExchangeRate{}, err
	}
	rate, err := c.RateAtRangeDates(startdate, enddate)
	c = nil // cleanup
	return rate, err
}
