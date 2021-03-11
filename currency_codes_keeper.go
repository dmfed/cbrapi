package cbrapi

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"golang.org/x/net/html/charset"
)

var (
	basePATH       = os.TempDir()
	keeperFilename = "cbrapi_codes.json"
	keeperTMP      = filepath.Join(basePATH, keeperFilename)
)

// globalVarAPICodes points to codeKeeper type initialized on start
// This enables other parts of program to find API codes, names etc. of
// currencies known to the Central Bank of Russia API
var globalVarAPICodes codeKeeper

type codeKeeper map[string]Currency

// initAPI looks for a local copy of CBR_API_codes.json and tries to load it
// into memory. If successful it points global variable APICodes to usable codeKeeper struct.
// If it fails to read local file then it tries to fetch codes from cbr.ru API
// If that also fails it returns an error indicating either that out request to API failed,
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
	data, err := ioutil.ReadFile(keeperTMP)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &ck)
	return ck, err
}

func codeKeeperFromAPI() (codeKeeper, error) {
	ck := make(codeKeeper)
	response, err := http.Get(EndpointCurrencyCodes)
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
	ioutil.WriteFile(keeperTMP, data, 0644)
	return ck, nil
}
