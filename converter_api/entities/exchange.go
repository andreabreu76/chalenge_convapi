package entities

import "time"

type Exchange struct {
	ID             uint
	Amount         float64
	FromCurrency   string
	ToCurrency     string
	Rate           float64
	ConvertedValue float64
	Date           time.Time
}
