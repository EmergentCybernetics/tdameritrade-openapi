/*
 * TD Ameritrade API
 *
 * TD Ameritrade API
 *
 * API version: 3.0.1
 * Contact: austin.millan@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CashAccountPositions struct for CashAccountPositions
type CashAccountPositions struct {
	ShortQuantity float32 `json:"shortQuantity,omitempty"`
	AveragePrice float32 `json:"averagePrice,omitempty"`
	CurrentDayProfitLoss float32 `json:"currentDayProfitLoss,omitempty"`
	CurrentDayProfitLossPercentage float32 `json:"currentDayProfitLossPercentage,omitempty"`
	LongQuantity float32 `json:"longQuantity,omitempty"`
	SettledLongQuantity float32 `json:"settledLongQuantity,omitempty"`
	SettledShortQuantity float32 `json:"settledShortQuantity,omitempty"`
	AgedQuantity float32 `json:"agedQuantity,omitempty"`
	Instrument string `json:"instrument,omitempty"`
	MarketValue float32 `json:"marketValue,omitempty"`
}