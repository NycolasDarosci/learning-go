package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"examble.com/crypto/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	if len(currency) != 3 {
		return nil, fmt.Errorf("3 characteres required: %d received", len(currency))
	}
	upCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upCurrency))

	if err != nil {
		return nil, err
	}

	var response CEXResponse

	if res.StatusCode == http.StatusOK {
		res, err := io.ReadAll(res.Body)

		if err != nil {
			return nil, err
		}
		//json := string(res)
		//fmt.Println(json)

		// parses the JSON-encoded data and stores the result in the value pointed to by v.
		// i want a function to change the content of the variable
		// pass the address of the variable
		err = json.Unmarshal(res, &response)

		if err != nil {
			return nil, err
		}
	} else {
		// creating an error
		return nil, fmt.Errorf("status code received: %v", res.StatusCode)
	}

	rate := datatypes.Rate{Currency: currency, Price: response.Bid}
	return &rate, nil
}
