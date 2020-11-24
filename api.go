package cbrcurrency

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
	errorIncorrectCode = errors.New("Provided currency code is not known to the API")
	errorNoData        = errors.New("API did not provide requested data")
)

// Currency represents a currency known to Central Bank API
// Fields of this struct are almost identical to what API returns
type Currency struct {
	NameRUS     string
	NameENG     string
	APIID       string
	Nominal     int
	ISONumCode  int
	ISOCharCode string
}

// ExchangeRate represents exchange rate of a currency at certain date.
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
// to request exchange rate of currency from the Central BAnk API
// with RateAT() and RateAtRangeDates() methods
func New(ISOcode string) (*Currency, error) {
	ISOcode = strings.ToUpper(ISOcode)
	if globalVarAPICodes == nil {
		err := initAPI()
		if err != nil {
			return nil, err
		}
	}
	item, ok := globalVarAPICodes[ISOcode]
	if !ok {
		return nil, errorIncorrectCode
	}
	return &item, nil
}

// NameRU returns currency name in Russian
func (c *Currency) NameRU() string {
	return c.NameRUS
}

// NameEN returnc currency name in English
func (c *Currency) NameEN() string {
	return c.NameENG
}

// NameISO returns ISO character code of currency
func (c *Currency) NameISO() string {
	return c.ISOCharCode
}

// CodeISO returns ISO numeric code of currency
func (c *Currency) CodeISO() int {
	return c.ISONumCode
}

// RateAt accept either "DD/MM/YYYY" formatted date or
// time.Time object. It sends request to the API and returns
// ExchangeRate object
func (c *Currency) RateAt(date interface{}) (ExchangeRate, error) {
	rate := ExchangeRate{}
	url := fmt.Sprintf(endpointSingleDate, apidate(date).stringobject())
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
			rate, _ := strconv.ParseFloat(strings.ReplaceAll(item.Value, ",", "."), 64)
			return ExchangeRate{ISOCode: c.ISOCharCode, Nominal: c.Nominal, Date: apidate(date).timeobject(), Rate: rate}, nil
		}
	}
	return rate, errorNoData
}

// RateAtRangeDates accept either "DD/MM/YYYY" formatted date or
// time.Time object. It sends request to the API and returns
// slice of ExchangeRate objects
func (c *Currency) RateAtRangeDates(startdate, enddate interface{}) ([]ExchangeRate, error) {
	rates := []ExchangeRate{}
	url := fmt.Sprintf(endpointDateRange, apidate(startdate).stringobject(), apidate(enddate).stringobject(), c.APIID)
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
		return rates, errorNoData
	default:
		return rates, nil
	}
}