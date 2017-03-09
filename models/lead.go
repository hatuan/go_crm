package models

type Lead struct {
	ID                *int64     `db:"id"  json:",string"`
	NumberSequenceID  int64      `db:"number_sequence_id"`
	No                string     `db:"no"`
	Description       string     `db:"description"`
	ContactType       int8       `db:"contact_type"`
	ContactName       string     `db:"contact_name"`
	ContactPersonName string     `db:"contact_person_name"`
	Street            string     `db:"street"`
	CityID            string     `db:"city_id"`
	CountyID          int64      `db:"county_id"`
	CountryRegionID   int64      `db:"country_region_id"`
	StateID           int64      `db:"state_id"`
	ZipPostalID       int64      `db:"zip_postal_id"`
	Phone             string     `db:"phone"`
	PhoneExtension    string     `db:"phone_extension"`
	Mobile            string     `db:"mobile"`
	Sms               string     `db:"sms"`
	Telex             string     `db:"telex"`
	Fax               string     `db:"fax"`
	Email             string     `db:"email"`
	URL               string     `db:"url"`
	Pager             string     `db:"pager"`
	Latitude          float32    `db:"latitude"`
	Longtude          float32    `db:"longtude"`
	Timezone          string     `db:"timezone"`
	AddressMasterID   int64      `db:"address_master_id"`
	DateOpen          *Timestamp `db:"date_open"`
	DateClose         *Timestamp `db:"date_close"`
	UserOwnerID       int64      `db:"user_owner_id"`
	UserOpenByID      int64      `db:"user_open_by_id"`
	UserCloseByID     int64      `db:"user_close_by_id"`
	Priority          int8       `db:"priority"`
	SaleUnitID        int64      `db:"sale_unit_id"`
	SourceTypeID      int64      `db:"source_type_id"`
	SourceID          int64      `db:"source_id"`
	RecCreatedByID    int64      `db:"rec_created_by" json:",string"`
	RecCreatedByUser  string     `db:"rec_created_by_user"`
	RecCreated        *Timestamp `db:"rec_created_at"`
	RecModifiedByID   int64      `db:"rec_modified_by" json:",string"`
	RecModifiedByUser string     `db:"rec_modified_by_user"`
	RecModified       *Timestamp `db:"rec_modified_at"`
	Status            int8       `db:"status"`
	Version           int16      `db:"version"`
	ClientID          int64      `db:"client_id" json:",string"`
	OrganizationID    int64      `db:"organization_id" json:",string"`
	Organization      string     `db:"organization"`
}
