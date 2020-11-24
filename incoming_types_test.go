package cbrcurrency

import (
	"encoding/xml"
	"fmt"
	"testing"
)

// Let's test that sample response from the API unmarshals
// into our structs
func Test_ForeignCurrencyAPICodesItem_Type_Unmarshalls(t *testing.T) {
	input := `<Item ID="R01010">
	<Name>Австралийский доллар</Name>
	<EngName>Australian Dollar</EngName>
	<Nominal>1</Nominal>
	<ParentCode>R01010 </ParentCode>
	<ISO_Num_Code>36</ISO_Num_Code>
	<ISO_Char_Code>AUD</ISO_Char_Code>
	</Item>`
	tobytes := []byte(input)
	var item ForeignCurrencyAPICodesItem
	err := xml.Unmarshal(tobytes, &item)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Sample input for type %T fails to unmarshall into our struct\n", item)
		t.Fail()
	}
	if item.APIID != "R01010" ||
		item.Name != "Австралийский доллар" ||
		item.EngName != "Australian Dollar" ||
		item.Nominal != int(1) ||
		item.ISONumCode != int(36) ||
		item.ISOCharCode != "AUD" {
		fmt.Printf("Sample input for type %T fails to unmarshall into our struct\n", item)
		t.Fail()
	}
}

func Test_ForeignCurrencyAPICodes_Type_Unmarshalls(t *testing.T) {
	input := `<Valuta name="Foreign Currency Market Lib">
		<Item ID="R01010">
		<Name>Австралийский доллар</Name>
		<EngName>Australian Dollar</EngName>
		<Nominal>1</Nominal>
		<ParentCode>R01010 </ParentCode>
		<ISO_Num_Code>36</ISO_Num_Code>
		<ISO_Char_Code>AUD</ISO_Char_Code>
		</Item>
		<Item ID="R01015">
		<Name>Австрийский шиллинг</Name>
		<EngName>Austrian Shilling</EngName>
		<Nominal>1000</Nominal>
		<ParentCode>R01015 </ParentCode>
		<ISO_Num_Code>40</ISO_Num_Code>
		<ISO_Char_Code>ATS</ISO_Char_Code>
		</Item>
		<Item ID="R01020A">
		<Name>Азербайджанский манат</Name>
		<EngName>Azerbaijan Manat</EngName>
		<Nominal>1</Nominal>
		<ParentCode>R01020 </ParentCode>
		<ISO_Num_Code>944</ISO_Num_Code>
		<ISO_Char_Code>AZN</ISO_Char_Code>
		</Item>
		<Item ID="R01035">
		<Name>Фунт стерлингов Соединенного королевства</Name>
		<EngName>British Pound Sterling</EngName>
		<Nominal>1</Nominal>
		<ParentCode>R01035 </ParentCode>
		<ISO_Num_Code>826</ISO_Num_Code>
		<ISO_Char_Code>GBP</ISO_Char_Code>
		</Item>
		</Valuta>`
	tobytes := []byte(input)
	var currencies ForeignCurrencyAPICodes
	err := xml.Unmarshal(tobytes, &currencies)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Sample input for type %T fails to unmarshall into our struct\n", currencies)
		t.Fail()
	}
	if len(currencies.Elements) != 4 ||
		currencies.Elements[3].APIID != "R01035" {
		fmt.Printf("Sample input for type %T fails to unmarshall into our struct\n", currencies)
		t.Fail()
	}
}

func Test_ResponseDailyElement_Type_Unmarshalls(t *testing.T) {
	input := `<Valute ID="R01035">
	<NumCode>826</NumCode>
	<CharCode>GBP</CharCode>
	<Nominal>1</Nominal>
	<Name>Фунт стерлингов Соединенного королевства</Name>
	<Value>43,8254</Value>
	</Valute>`
	tobytes := []byte(input)
	var element ResponseDailyElement
	err := xml.Unmarshal(tobytes, &element)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Sample input for type %T fails to unmarshall into our struct\n", element)
		t.Fail()
	}
	if element.APIID != "R01035" ||
		element.NumericCode != int(826) ||
		element.CharacterCode != "GBP" ||
		element.Nominal != int(1) ||
		element.Name != "Фунт стерлингов Соединенного королевства" ||
		element.Value != "43,8254" {
		fmt.Printf("Sample input for type %T fails to unmarshall into our struct\n", element)
		t.Fail()
	}
}

func Test_ResponseDaily_Type_Unmarshalls(t *testing.T) {
	input := `<ValCurs Date="02.03.2002" name="Foreign Currency Market">
	<Valute ID="R01010">
	<NumCode>036</NumCode>
	<CharCode>AUD</CharCode>
	<Nominal>1</Nominal>
	<Name>Австралийский доллар</Name>
	<Value>16,0102</Value>
	</Valute>
	<Valute ID="R01035">
	<NumCode>826</NumCode>
	<CharCode>GBP</CharCode>
	<Nominal>1</Nominal>
	<Name>Фунт стерлингов Соединенного королевства</Name>
	<Value>43,8254</Value>
	</Valute>
	<Valute ID="R01090">
	<NumCode>974</NumCode>
	<CharCode>BYR</CharCode>
	<Nominal>1000</Nominal>
	<Name>Белорусских рублей</Name>
	<Value>18,4290</Value>
	</Valute>
	</ValCurs>`
	tobytes := []byte(input)
	var response ResponseDaily
	err := xml.Unmarshal(tobytes, &response)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Sample input for type %T fails to unmarshall into our struct\n", response)
		t.Fail()
	}
	if len(response.Elements) != 3 ||
		response.Date != "02.03.2002" ||
		response.Elements[2].Name != "Белорусских рублей" ||
		response.Elements[1].CharacterCode != "GBP" {
		fmt.Printf("Sample input for type %T fails to unmarshall into our struct\n", response)
		t.Fail()
	}
}

func Test_ResponseRangeElement_Type_Unmarshalls(t *testing.T) {
	input := `<Record Date="06.03.2001" Id="R01235">
	<Nominal>1</Nominal>
	<Value>28,6600</Value>
	</Record>`
	tobytes := []byte(input)
	var element ResponseRangeElement
	err := xml.Unmarshal(tobytes, &element)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Sample input for type %T fails to unmarshall into our struct\n", element)
		t.Fail()
	}
	if element.Date != "06.03.2001" ||
		element.APIID != "R01235" ||
		element.Nominal != int(1) ||
		element.Value != "28,6600" {
		fmt.Printf("Sample input for type %T fails to unmarshall into our struct\n", element)
		t.Fail()
	}
}

func Test_ResponseRange_Type_Unmarshalls(t *testing.T) {
	input := `<ValCurs ID="R01235" DateRange1="02.03.2001" DateRange2="14.03.2001" name="Foreign Currency Market Dynamic">
	<Record Date="02.03.2001" Id="R01235">
	<Nominal>1</Nominal>
	<Value>28,6200</Value>
	</Record>
	<Record Date="03.03.2001" Id="R01235">
	<Nominal>1</Nominal>
	<Value>28,6500</Value>
	</Record>
	<Record Date="06.03.2001" Id="R01235">
	<Nominal>1</Nominal>
	<Value>28,6600</Value>
	</Record>
	</ValCurs>`
	tobytes := []byte(input)
	var respRange ResponseRange
	err := xml.Unmarshal(tobytes, &respRange)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Sample input for type %T fails to unmarshall into our struct\n", respRange)
		t.Fail()
	}
	if respRange.APIID != "R01235" ||
		respRange.DateStart != "02.03.2001" ||
		respRange.DateEnd != "14.03.2001" ||
		len(respRange.Elements) != 3 ||
		respRange.Elements[2].Nominal != int(1) ||
		respRange.Elements[1].Value != "28,6500" ||
		respRange.Elements[0].Date != "02.03.2001" ||
		respRange.Elements[0].APIID != "R01235" {
		fmt.Printf("Sample input for type %T fails to unmarshall into our struct\n", respRange)
		t.Fail()
	}
}
