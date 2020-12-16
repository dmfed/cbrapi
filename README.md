# Central Bank of Russia API implementation

Package **cbrapi** partially implements API of Central Bank of Russia. Implemented functions allow to quote currency rates at specified date or range of dates.
See [website of Central Bank of Russia](https://www.cbr.ru/development/) for details.

Usage is quite simple (see below sample code). The only two points that need to be mentioned are as follows. 

1. Methods which accept dates have signatures with **interface{}** type. These require either time.Time object or date in form of string formatted strictly as
"DD/MM/YYYY". Supplying string formatted in a different way would result in a error.

2. ExchangeRate object implements Stringer, so if you care to only receive the float (the rate itself), use relevant field and do not Print/Sprint etc.

See [package documentation at pkg.go.dev](https://pkg.go.dev/github.com/dmfed/cbrapi) for details.

To use the library: **go get github.com/dmfed/cbrapi**

And here comes a working example: 

```go
import "https://github.com/dmfed/cbrapi/"

package main

import (
	"fmt"
	"time"

	"github.com/dmfed/cbrapi"
)

func main() {
	usd, err := cbrapi.New("USD")
	if err != nil {
		fmt.Println(err)
		return
	}
	usdrate, err := usd.RateAtDate(time.Now()) // returns current exchange rate USD/RUB
	fmt.Println(usdrate, err)
	eurrate, _ := cbrapi.QuoteAtDate("EUR", "01/09/2020") // returns exchange rate of EUR at September 1st 2020
	fmt.Println(eurrate)
	usdrange, _ := usd.RateAtRangeDates("01/09/2020", "04/09/2020")
	for _, item := range usdrange {
		fmt.Println(item)
	}
}
```
