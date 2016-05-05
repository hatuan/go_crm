package models

type Role struct {
	ID             string       `json:"id"`
	Name           string       `json:"name"`
	ClientID       string       `json:"client_id"`
	Client         Client       `json:"client"`
	OrganizationID string       `json:"organization_id"`
	Organization   Organization `json:"organization"`
}
