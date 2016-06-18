package models

type Currency struct {
	ID                string       `json:"id"`
	Code              string       `json:"code"`
	Name              string       `json:"name"`
	RecCreatedByID    string       `json:"rec_created_by_id"`
	RecCreatedByUser  User         `json:"rec_created_by_user" db:"-"`
	RecCreated        *Timestamp   `json:"rec_created"`
	RecModifiedByID   string       `json:"rec_modified_by_id"`
	RecModifiedByUser User         `json:"rec_modified_by_user" db:"-"`
	RecModified       *Timestamp   `json:"rec_modified"`
	Status            int8         `json:"status"`
	Version           int16        `json:"version"`
	ClientID          string       `json:"client_id"`
	OrganizationID    string       `json:"organization_id"`
	Organization      Organization `json:"organization" db:"-"`
}
