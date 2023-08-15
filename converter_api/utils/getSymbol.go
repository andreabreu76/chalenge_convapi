package utils

func GetCurrencySymbol(currencyCode string) string {
	switch currencyCode {
	case "USD":
		return "$"
	case "BRL":
		return "R$"
	case "EUR":
		return "€"
	case "BTC":
		return "₿" // U+20BF (Bitcoin don't have an official symbol, I try this one)
	default:
		return ""
	}
}
