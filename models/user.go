package models

import (
	"erpvietnam/crm/log"
	"erpvietnam/crm/settings"
	"errors"

	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

// User represents the user model
type User struct {
	ID                  string        `json:"id"`
	Name                string        `json:"name" db:"name"`
	Password            string        `json:"-" db:"password"`
	Comment             string        `json:"comment"  db:"comment"`
	FullName            string        `json:"full_name" db:"full_name"`
	PasswordAnswer      string        `json:"-" db:"password_answer"`
	PasswordQuestion    string        `json:"-" db:"password_question"`
	Email               string        `json:"email"`
	CreatedDate         *Timestamp    `json:"created_date,omitempty" db:"created_date"`
	IsActivated         bool          `json:"is_activated" db:"is_activated"`
	IsLockedOut         bool          `joson:"is_locked_out" db:"is_locked_out" `
	LastLockedOutDate   *Timestamp    `json:"last_locked_out_date,omitempty" db:"last_locked_out_date"`
	LastLockedOutReason string        `json:"last_locked_out_reason" db:"last_locked_out_reason"`
	LastLoginDate       *Timestamp    `json:"last_login_date,omitempty" db:"last_login_date"`
	LastLoginIP         string        `json:"last_login_ip" db:"last_login_ip"`
	LastModifiedDate    *Timestamp    `json:"last_modified_date,omitempty" db:"last_modified_date"`
	ClientID            string        `json:"client_id" db:"client_id"`
	Client              *Client       `json:"client" db:"-"`
	OrganizationID      string        `json:"organization_id" db:"organization_id"`
	Organization        *Organization `json:"organization" db:"-"`
	CultureUIID         string        `json:"culture_ui_id" db:"culture_ui_id"`
	Roles               []Role        `json:"roles" db:"-"`
}

// ErrUsernameTaken is thrown when a user attempts to register a username that is taken.
var ErrUsernameTaken = errors.New("username already taken")

// GetUser returns the user that the given id corresponds to. If no user is found, an
// error is thrown.
func GetUser(id string) (User, error) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	u := User{}
	err = db.Get(&u, "SELECT * FROM \"user\" WHERE id=$1", id)
	if err != nil {
		return u, err
	}
	return u, nil
}

// GetUserByUsername returns the user that the given username corresponds to. If no user is found, an
// error is thrown.
func GetUserByUsername(name string) (User, error) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	u := User{}
	err = db.Get(&u, "SELECT * FROM \"user\" WHERE name=$1", name)
	// No issue if we don't find a record
	if err == sql.ErrNoRows {
		return u, nil
	} else if err == nil {
		return u, ErrUsernameTaken
	}
	return u, err
}

type PreferenceDTO struct {
	OrganizationID string
	CultureUIID    string
	WorkingDate    *Timestamp
}

func (u User) GetPreference() (PreferenceDTO, error) {
	preference := PreferenceDTO{}

	//TODO: check if u.OrganizationID == "" => load rootOrganization
	preference.OrganizationID = u.OrganizationID

	preference.WorkingDate = &Timestamp{time.Now()}

	if u.CultureUIID == "" {
		u.CultureUIID = "vi-VN"
	}
	preference.CultureUIID = u.CultureUIID

	return preference, nil
}

func (u *User) SetPreference(preference PreferenceDTO) error {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.NamedExec(`UPDATE "user" SET organization_id = :organization_id, culture_ui_id = :culture_ui_id WHERE id = :id`,
		map[string]interface{}{
			"organization_id": preference.OrganizationID,
			"culture_ui_id":   preference.CultureUIID,
			"id":              u.ID,
		})
	if err != nil {
		return err
	}
	u.OrganizationID = preference.OrganizationID
	u.CultureUIID = preference.CultureUIID

	return nil
}
