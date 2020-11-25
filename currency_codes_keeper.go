package cbrapi

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
)

var basePATH = "/tmp/" // Program saves CBR_API_codes.json to this path

// globalVarAPICodes points to codeKeeper type initialized on start
// This enables other parts of program to find API codes, names etc. of
// currencies known to the Central Bank of Russia API
var globalVarAPICodes codeKeeper

type codeKeeper map[string]Currency

// InitAPI looks for a local copy of CBR_API_codes.json and tries to load it
// into memory. IF successful it points global variable APICodes to usable codeKeeper struct.
// If it failes to read local file then it tries to fetch codes from cbr.ru API
// IF that fails it returns an error indicating either that out request to API failed,
// or that received xml failed to deserialize.
func initAPI() error {
	ck, err := codeKeeperFromFile()
	if err == nil {
		globalVarAPICodes = ck
		return nil
	}
	ck, err = codeKeeperFromAPI()
	if err == nil {
		globalVarAPICodes = ck
		return nil
	}
	return err
}

func codeKeeperFromFile() (codeKeeper, error) {
	ck := make(codeKeeper)
	data, err := ioutil.ReadFile(basePATH + "CBR_API_codes.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &ck)
	return ck, err
}

func codeKeeperFromAPI() (codeKeeper, error) {
	ck := make(codeKeeper)
	response, err := http.Get(endpointCurrencyCodes)
	if err != nil {
		return ck, err
	}
	defer response.Body.Close()
	var codes ForeignCurrencyAPICodes
	decoder := xml.NewDecoder(response.Body)
	decoder.CharsetReader = charset.NewReaderLabel
	if err := decoder.Decode(&codes); err != nil {
		return ck, err
	}
	for _, item := range codes.Elements {
		ck[item.ISOCharCode] = Currency{NameRUS: item.Name,
			NameENG: item.EngName, Nominal: item.Nominal, APIID: item.APIID,
			ISONumCode: item.ISONumCode, ISOCharCode: item.ISOCharCode}
	}
	data, _ := json.MarshalIndent(ck, "", "    ")
	ioutil.WriteFile(basePATH+"CBR_API_codes.json", data, 0644)
	return ck, nil
}
