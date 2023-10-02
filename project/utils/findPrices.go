package utils

import "strings"

func FindPriceIndexes(slice []string, was, now string) (oldPrice, newPrice string) {
	for i, s := range slice {
		if strings.ToLower(s) == strings.ToLower(was) {
			oldPrice = slice[i+1]
			removeCurrencySign("$", &oldPrice)
		}
		if strings.ToLower(s) == strings.ToLower(now) {
			newPrice = slice[i+1]
			removeCurrencySign("$", &newPrice)
		}
	}

	return
}

func removeCurrencySign(currency string, price *string) {
	if strings.Contains(*price, currency) {
		_, clearPrice, _ := strings.Cut(*price, "$")
		price = &clearPrice
	}
}
