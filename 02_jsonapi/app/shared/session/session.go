package session

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	// Store : responsible for storing cookies (I presume - as the name implies)
	Store *sessions.CookieStore
	// Name : sessions name
	Name string
)

// Session stores session level information
type Session struct {
	Options   sessions.Options `json:"Options"`
	Name      string           `json:"Name"`
	SecretKey string           `json:"SecretKey"`
}

// Configure the session cookie store
func Configure(s Session) {
	Store = sessions.NewCookieStore([]byte(s.SecretKey))
	Store.Options = &s.Options
	Name = s.Name
}

// Instance returns a new session (never an error though)
func Instance(r *http.Request) *sessions.Session {
	session, _ := Store.Get(r, Name)
	return session
}

// Empty deletes all the current session values
func Empty(sess *sessions.Session) {
	// Clear out all stored values in the cookie
	for k := range sess.Values {
		delete(sess.Values, k)
	}
}
