package httpapi

import "strings"

func normalizeCurrency(value string) (string, bool) {
	currency := strings.ToUpper(strings.TrimSpace(value))

	if len(currency) != 3 {
		return "", false
	}

	for i := 0; i < len(currency); i++ {
		if currency[i] < 'A' || currency[i] > 'Z' {
			return "", false
		}
	}

	return currency, true
}
