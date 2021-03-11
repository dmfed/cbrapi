package cbrapi

var (
	BaseURL               string = "https://www.cbr.ru/scripts/"
	EndpointSingleDate           = BaseURL + "XML_daily.asp?date_req=%v"
	EndpointDateRange            = BaseURL + "XML_dynamic.asp?date_req1=%v&date_req2=%v&VAL_NM_RQ=%v"
	EndpointCurrencyCodes        = BaseURL + "XML_valFull.asp"
)
