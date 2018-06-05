package auth

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/volatiletech/authboss"
)

const sessionCookieName = "bikemeet"

var sessionStore *sessions.CookieStore

// SessionStore holds the http information related to the session
type SessionStore struct {
	w http.ResponseWriter
	r *http.Request
}

// NewSessionStore creates a new SessionStore
func NewSessionStore(w http.ResponseWriter, r *http.Request) authboss.ClientStorer {
	return &SessionStore{w, r}
}

// Get fetches the value from the session store with the given key
func (s SessionStore) Get(key string) (string, bool) {
	session, err := sessionStore.Get(s.r, sessionCookieName)
	if err != nil {
		fmt.Println(err)
		return "", false
	}

	strInf, ok := session.Values[key]
	if !ok {
		return "", false
	}

	str, ok := strInf.(string)
	if !ok {
		return "", false
	}

	return str, true
}

// Put puts new data into the session store, or updates existing data
func (s SessionStore) Put(key, value string) {
	session, err := sessionStore.Get(s.r, sessionCookieName)
	if err != nil {
		fmt.Println(err)
		return
	}

	session.Values[key] = value
	if err := session.Save(s.r, s.w); err != nil {
		fmt.Println(err)
	}
}

// Del removes data from the session store
func (s SessionStore) Del(key string) {
	session, err := sessionStore.Get(s.r, sessionCookieName)
	if err != nil {
		fmt.Println(err)
		return
	}

	delete(session.Values, key)
	if err := session.Save(s.r, s.w); err != nil {
		fmt.Println(err)
	}
}
