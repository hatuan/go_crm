package models

import (
	"erpvietnam/crm/log"
	"erpvietnam/crm/settings"
	"fmt"
	"strconv"
	"strings"

	"database/sql"
	"database/sql/driver"

	"github.com/jmoiron/sqlx"
)

type Flash struct {
	Type    string
	Message string
}

type Response struct {
	ReturnStatus     bool
	ReturnMessage    []string
	ValidationErrors map[string]interface{}
	TotalPages       int
	TotalRows        int
	PageSize         int
	IsAuthenticated  bool
	Data             map[string]interface{}
}

type Token struct {
	TransactionalInformation
	Token string //`json:"token"` //must use lowercase
}

type LoginDTO struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type TransactionalInformation struct {
	ReturnStatus     bool
	ReturnMessage    []string
	ValidationErrors map[string]interface{}
	TotalPages       int
	TotalRows        int
	PageSize         int
	IsAuthenticated  bool
}

type InquiryInformationDTO struct {
	CurrentPageNumber int
	PageSize          int
	SearchCondition   string
	SortDirection     string
	SortExpression    string
}

type InfiniteScrollingInformation struct {
	After          string
	FetchSize      string
	SortDirection  string
	SortExpression string
}

type ApplicationMenuDTO struct {
	MenuID                 string `json:"menu_id"`
	Description            string `json:"description"`
	Route                  string `json:"route"`
	Module                 string `json:"module"`
	MenuOrder              int    `json:"menu_order"`
	RequiresAuthentication bool   `json:"requires_authentication"`
}

//ApplicationModelDTO user for return from  controllers.InitializeApplication
type ApplicationModelDTO struct {
	TransactionalInformation
	MenuItems []ApplicationMenuDTO `json:"menu_items"`
}

const EmptyUUID = "00000000-0000-0000-0000-000000000000"

func CheckUnique(table, ID, code, orgID string) (bool, error) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if ID == "" {
		ID = EmptyUUID
	}
	strSQL := fmt.Sprintf("SELECT id FROM %s WHERE code = $1 AND id <> $2 AND organization_id = $3", table)
	log.Info(strSQL)

	var otherID string
	err = db.Get(&otherID, strSQL, code, ID, orgID)

	if err != nil && err == sql.ErrNoRows {
		return true, nil
	} else if err != nil {
		log.Fatal(err)
		return false, err
	}
	return false, nil
}

// InterfaceArray is a type implementing the sql/driver/value interface
// This is due to the native driver not supporting arrays...
type InterfaceArray []interface{}

// Value returns the driver compatible value
func (a InterfaceArray) Value() (driver.Value, error) {
	var strs []string
	for _, i := range a {
		if str, ok := i.(string); ok {
			strs = append(strs, q(str))
		} else {
			strs = append(strs, "''")
		}
	}
	return "{" + strings.Join(strs, ",") + "}", nil
}

// Int64Array is a type implementing the sql/driver/value interface
// This is due to the native driver not supporting arrays...
type Int64Array []int64

// Value returns the driver compatible value
func (a Int64Array) Value() (driver.Value, error) {
	var strs []string
	for _, i := range a {
		strs = append(strs, strconv.FormatInt(i, 10))
	}
	return "{" + strings.Join(strs, ",") + "}", nil
}

// StringArray is a type implementing the sql/driver/value interface
// This is due to the native driver not supporting arrays...
type StringArray []string

// Value returns the driver compatible value
func (a StringArray) Value() (driver.Value, error) {
	var strs []string
	for _, i := range a {
		strs = append(strs, q(i))
	}
	return "{" + strings.Join(strs, ",") + "}", nil
}

// q
func q(s string) string {
	re := strings.NewReplacer("'", "''")
	return "'" + re.Replace(s) + "'"
}
