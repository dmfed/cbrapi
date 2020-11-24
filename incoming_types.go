package cbrapi

import "encoding/xml"

// The following structures are used to unmarshall response from
// http://www.cbr.ru/scripts/XML_valFull.asp

// ForeignCurrencyAPICodes is a list of elements representing
// currencies known to the API of Central Bank.
type ForeignCurrencyAPICodes struct {
	XMLName  xml.Name                      `xml:"Valuta"`
	Elements []ForeignCurrencyAPICodesItem `xml:"Item"`
}

// ForeignCurrencyAPICodesItem represents XML structure of
// currency item returned by the API
type ForeignCurrencyAPICodesItem struct {
	XMLName     xml.Name `xml:"Item"`
	APIID       string   `xml:"ID,attr"`
	Name        string   `xml:"Name"`
	EngName     string   `xml:"EngName"`
	Nominal     int      `xml:"Nominal"`
	ISONumCode  int      `xml:"ISO_Num_Code"`
	ISOCharCode string   `xml:"ISO_Char_Code"`
}

// Below two structures used to unmarshall response from
// http://www.cbr.ru/scripts/XML_daily.asp
// Allows request of exchange rates of all currencies known to the API
// on a specified date.

// ResponseDaily represents XML structure of response to request
// of singe date exchange rates from the endpoint
// This response contains all currencies rates as of the requested
// date
type ResponseDaily struct {
	XMLName  xml.Name               `xml:"ValCurs"`
	Date     string                 `xml:"Date,attr"`
	Elements []ResponseDailyElement `xml:"Valute"`
}

// ResponseDailyElement represents XML element of ResponseDaily
type ResponseDailyElement struct {
	XMLName       xml.Name `xml:"Valute"`
	APIID         string   `xml:"ID,attr"`
	NumericCode   int      `xml:"NumCode"`
	CharacterCode string   `xml:"CharCode"`
	Nominal       int      `xml:"Nominal"`
	Name          string   `xml:"Name"`
	Value         string   `xml:"Value"`
}

// Following structures are used to unmarshall response from
// http://www.cbr.ru/scripts/XML_dynamic.asp
// Allows to request exchange rate of currency specified by API custom currency code
// for range of dates.

// ResponseRange list of exchange rates for requested dates
type ResponseRange struct {
	XMLName   xml.Name               `xml:"ValCurs"`
	APIID     string                 `xml:"ID,attr"`
	DateStart string                 `xml:"DateRange1,attr"`
	DateEnd   string                 `xml:"DateRange2,attr"`
	Elements  []ResponseRangeElement `xml:"Record"`
}

// ResponseRangeElement represents exchange rate for single date
type ResponseRangeElement struct {
	XMLName xml.Name `xml:"Record"`
	Date    string   `xml:"Date,attr"`
	APIID   string   `xml:"Id,attr"`
	Nominal int      `xml:"Nominal"`
	Value   string   `xml:"Value"`
}
