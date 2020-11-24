package cbrapi

var (
	baseURL               string = "https://www.cbr.ru/scripts/"
	endpointSingleDate           = baseURL + "XML_daily.asp?date_req=%v"
	endpointDateRange            = baseURL + "XML_dynamic.asp?date_req1=%v&date_req2=%v&VAL_NM_RQ=%v"
	endpointCurrencyCodes        = baseURL + "XML_valFull.asp"
)
