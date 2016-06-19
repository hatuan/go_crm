package models

import (
	"database/sql"
	"erpvietnam/crm/log"

	"github.com/shopspring/decimal"

	"erpvietnam/crm/settings"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Client represents the client model
type Client struct {
	ClientID                    string          `db:"id"`
	Name                        string          `db:"name"`
	Version                     int16           `db:"version"`
	IsActivated                 bool            `db:"is_activated"`
	RecCreatedBy                string          `db:"rec_created_by"`
	RecCreatedByUser            *User           `db:"-"`
	RecCreatedAt                *time.Time      `db:"rec_created_at"`
	RecModifiedBy               string          `db:"rec_modified_by"`
	RecModifiedByUser           *User           `db:"-"`
	RecModifiedAt               *time.Time      `db:"rec_modified_at"`
	CultureID                   string          `db:"culture_id"`
	AmountDecimalPlaces         int16           `db:"amount_decimal_places"`
	AmountRoundingPrecision     decimal.Decimal `db:"amount_rounding_precision"`
	UnitAmountDecimalPlaces     int16           `db:"unit-amount_decimal_places"`
	UnitAmountRoundingPrecision decimal.Decimal `db:"unit-amount_rounding_precision"`
	CurrencyLCYId               string          `db:"currency_lcy_id"`
	CurrencyLCY                 *Currency       `db:"-"`
	Organizations               []Organization  `db:"-"`
}

// ErrOrganizationsIsEmpty is thrown when do not found any Organization.
var ErrOrganizationsIsEmpty = errors.New("Organizations is empty")

// ErrClientNotFound is thrown when do not found any Client.
var ErrClientNotFound = errors.New("Client not found")

func (c *Client) GetOrganizations() ([]Organization, error) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Fatal(err)
		return []Organization{}, err
	}
	defer db.Close()

	organizations := []Organization{}
	err = db.Select(&organizations, "SELECT * FROM organization WHERE client_id = $1", c.ClientID)
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

func (c *Client) Get(id string) error {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.QueryRowx("SELECT *  FROM client WHERE id=$1", id).StructScan(c)
	if err == sql.ErrNoRows {
		return ErrClientNotFound
	} else if err != nil {
		return err
	}

	return nil
}
