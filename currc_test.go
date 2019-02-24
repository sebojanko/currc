package main

import (
	"strconv"
	"testing"
)

var currency string = "USD_HRK"
var flag_t string = "fd"
var amount string = "10"

func TestRetrieveRatesFromAPI(t *testing.T) {
	a := retrieveRatesFromAPI(currency)
	if a == nil {
		t.Error("Response is empty.")
	}
}

func TestUnmarshalAPIRates(t *testing.T) {
	response := retrieveRatesFromAPI(currency)
	data, err := unmarshalAPIRates(response)
	if err != nil {
		t.Error(err)
	}
	if data == nil {
		t.Error("Data is empty.")
	}
	if data[currency] == nil {
		t.Error("Multiplier is nonexistent.")
	}
	if data[currency]["val"] == 0 {
		t.Error("Multiplies is zero.")
	}
}

func TestConvertCurrency(t *testing.T) {
	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		panic(err)
	}
	data, err := unmarshalAPIRates(retrieveRatesFromAPI(currency))
	if err != nil {
		t.Error(err)
	}
	calculatedAmount := amountFloat * data[currency]["val"]

	convertedAmount := convertCurrency(flag_t, amount)
	if convertedAmount == 0 {
		t.Error("Converted amount is zero.")
	}

	if convertedAmount != calculatedAmount {
		t.Errorf("Converted amount is wrong, expected %f, got %f", calculatedAmount, convertedAmount)
	}

}
