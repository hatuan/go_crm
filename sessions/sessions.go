package sessions

import (
	"encoding/gob"
	"erpvietnam/crm/models"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"

)

//init registers the necessary models to be saved in the session later
func init() {
	gob.Register(&models.User{})
	gob.Register(&models.Flash{})
}

// Store contains the session information for the request
var Store = sessions.NewCookieStore(
	[]byte(securecookie.GenerateRandomKey(64)), //Signing key
	[]byte(securecookie.GenerateRandomKey(32)))