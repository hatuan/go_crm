package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"erpvietnam/crm/log"
)

var environments = map[string]string{
	"production":    "settings/prod.json",
	"preproduction": "settings/pre.json",
	"tests":         "settings/tests.json",
}

type DatabaseInfo struct {
	DriverName string `json:"driver_name"`
	User       string `json:"user"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Host       string `json:"host"`
	Port       string `json:"port"`
}

type settings struct {
	ListenURL          string       `json:"listen_url"`
	PrivateKeyPath     string       `json:"privatekey_path"`
	PublicKeyPath      string       `json:"publickey_path"`
	JWTExpirationDelta int          `json:"jwt_expiration_delta"`
	Database           DatabaseInfo `json:"database"`
}

func (s *settings) GetDbConn() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		s.Database.User, s.Database.Password, s.Database.Name, s.Database.Host, s.Database.Port)
}

// Settings contains the initialized configuration struct
var Settings settings
var env = "preproduction"

func init() {
	env = os.Getenv("GO_ENV")
	if env == "" {
		fmt.Println("Warning: Setting preproduction environment due to lack of GO_ENV value")
		env = "preproduction"
	}
	LoadSettingsByEnv(env)
}

func LoadSettingsByEnv(env string) {
	content, err := ioutil.ReadFile(environments[env])
	if err != nil {
		log.Panic("Error while reading config file", err)
		panic(err)
	}
	jsonErr := json.Unmarshal(content, &Settings)
	if jsonErr != nil {
		log.Panic("Error while parsing config file", jsonErr)
		panic(err)
	}
}
