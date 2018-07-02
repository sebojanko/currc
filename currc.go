package main

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
	"fmt"
	"flag"
	"strconv"
)

var help bool
var fromDollarFlag float64
var fromEuroFlag float64
var toDollarFlag float64
var toEuroFlag float64
var flagToCurrency = map[string]string{
	"fd": "USD_HRK",
	"fe": "EUR_HRK",
	"td": "HRK_USD",
	"te": "HRK_EUR"}

// TODO export flags to txt file so that adding more is easier
// TODO working with non-compact API
func init() {
	flag.Bool("h", false, "lists all commands")
	flag.Float64Var(&fromDollarFlag, "fd", 0.0, "convert from `<dollars>` to hrk")
	flag.Float64Var(&fromEuroFlag, "fe", 0.0, "convert from `<euros>` hrk")
	flag.Float64Var(&toDollarFlag, "td", 0.0, "convert from `<hrk>` to dollars")
	flag.Float64Var(&toEuroFlag, "te", 0.0, "convert from `<hrk>` to euros")
}

func main() {
	flag.Parse()

	if help || flag.NFlag() == 0 {
		flag.PrintDefaults()
		return
	}

	flag.Visit(func(f *flag.Flag) {
		convertCurrency(f.Name, f.Value)
	})
}

func convertCurrency(name string, value flag.Value) {
	currenciesToConvert := flagToCurrency[name]

	data, err := retrieveRates(currenciesToConvert)

	amount, err := strconv.ParseFloat(value.String(), 64)
	if err != nil {
		panic(err)
	}
	multiplier := data[currenciesToConvert]["val"]
	fmt.Println(multiplier * amount)

}
func retrieveRates(currenciesToConvert string) (map[string]map[string]float64, error) {
	responseBody := retrieveRatesFromAPI(currenciesToConvert)

	data, err := unmarshalAPIRates(responseBody)
	return data, err
}

func unmarshalAPIRates(responseBody []byte) (map[string]map[string]float64, error) {
	var data map[string]map[string]float64
	err := json.Unmarshal([]byte(responseBody), &data)
	if err != nil {
		panic(err)
	}
	return data, err
}

func retrieveRatesFromAPI(currenciesToConvert string) []byte {
	resp, _ := http.Get("https://free.currencyconverterapi.com/api/v5/convert?q=" + currenciesToConvert + "&compact=y")
	responseBody, _ := ioutil.ReadAll(resp.Body)
	return responseBody
}
