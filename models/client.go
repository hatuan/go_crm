package models

import (
	"github.com/shopspring/decimal"
	"database/sql"
	"erpvietnam/crm/log"

	"erpvietnam/crm/settings"
	"errors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Client represents the client model
type Client struct {
	ClientID                    string          `json:"client_id"`
	Name                        string          `json:"name"`
	Version                     int16           `json:"version"`
	IsActivated                 bool            `json:"is_activated"`
	RecCreatedByID              string          `json:"rec_created_by_id"`
	RecCreatedByUser            User            `json:"rec_created_by_user" db:"-"`
	RecCreated                  *Timestamp      `json:"rec_created"`
	RecModifiedByID             string          `json:"rec_modified_by_id"`
	RecModifiedByUser           User            `json:"rec_modified_by_user" db:"-"`
	RecModified                 *Timestamp      `json:"rec_modified"`
	CultureID                   string          `json:"culture_id"`
	AmountDecimalPlaces         int16           `json:"amount_decimal_places"`
	AmountRoundingPrecision     decimal.Decimal `json:"amount_rounding_precision"`
	UnitAmountDecimalPlaces     int16           `json:"unit_amount_decimal_places"`
	UnitAmountRoundingPrecision decimal.Decimal `json:"unit_amount_rounding_precision"`
	CurrencyLCYId               string          `json:"currency_lcy_id"`
	CurrencyLCY                 *Currency       `json:"currency_lcy" db:"-"`
	Organizations               []Organization  `json:"organizations" db:"-"`
}

// ErrOrganizationsIsEmpty is thrown when do not found any Organization.
var ErrOrganizationsIsEmpty = errors.New("Organizations is empty")


func (c * Client) Organizations() ([]Organization, error) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Fatal(err)
		return []Organization{}, err
	}
	defer db.Close()

	organizations := []Organization{}
	err = db.Get(&organizations, "SELECT * FROM organization WHERE client_id = $1", c.ClientID)
	if err != nil && err == sql.ErrNoRows {
		log.Error(err)
		return organizations, ErrOrganizationsIsEmpty
	} else if err != nil {
		log.Error(err)
		return organizations, err
	}
	c.Organizations = organizations

	return organizations, nil
}
