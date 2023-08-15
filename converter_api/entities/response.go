package entities

type ConversionResult struct {
	ConvertedValue float64 `json:"valorConvertido"`
	ExchangeSymbol string  `json:"simboloMoeda"`
}
