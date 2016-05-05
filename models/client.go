package models

import (
	"github.com/shopspring/decimal"
)

// Client represents the client model
type Client struct {
	ClientID                    string          `json:"client_id"`
	Name                        string          `json:"name"`
	Version                     int16           `json:"version"`
	IsActivated                 bool            `json:"is_activated"`
	RecCreatedByID              string          `json:"rec_created_by_id"`
	RecCreatedByUser            User            `json:"rec_created_by_user"`
	RecCreated                  *Timestamp      `json:"rec_created"`
	RecModifiedByID             string          `json:"rec_modified_by_id"`
	RecModifiedByUser           User            `json:"rec_modified_by_user"`
	RecModified                 *Timestamp      `json:"rec_modified"`
	CultureID                   string          `json:"culture_id"`
	AmountDecimalPlaces         int16           `json:"amount_decimal_places"`
	AmountRoundingPrecision     decimal.Decimal `json:"amount_rounding_precision"`
	UnitAmountDecimalPlaces     int16           `json:"unit_amount_decimal_places"`
	UnitAmountRoundingPrecision decimal.Decimal `json:"unit_amount_rounding_precision"`
	CurrencyLCYId               string          `json:"currency_lcy_id"`
	CurrencyLCY                 *Currency       `json:"currency_lcy"`
	Organizations               []Organization  `json:"organizations"`
}
