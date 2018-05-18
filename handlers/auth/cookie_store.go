package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/volatiletech/authboss"
)

var cookieStore *securecookie.SecureCookie

// CookieStore holds the http information related to cookies
type CookieStore struct {
	w http.ResponseWriter
	r *http.Request
}

// NewCookieStore creates a new CookieStore
func NewCookieStore(w http.ResponseWriter, r *http.Request) authboss.ClientStorer {
	return &CookieStore{w, r}
}

// Get fetches a cookie from the store
func (s CookieStore) Get(key string) (string, bool) {
	cookie, err := s.r.Cookie(key)
	if err != nil {
		return "", false
	}

	var value string
	err = cookieStore.Decode(key, cookie.Value, &value)
	if err != nil {
		return "", false
	}

	return value, true
}

// Put put a new cookie into the store, or update an existing one
func (s CookieStore) Put(key, value string) {
	encoded, err := cookieStore.Encode(key, value)
	if err != nil {
		fmt.Println(err)
	}

	cookie := &http.Cookie{
		Expires: time.Now().UTC().AddDate(1, 0, 0),
		Name:    key,
		Value:   encoded,
		Path:    "/",
	}
	http.SetCookie(s.w, cookie)
}

// Del removes a cookie from the store
func (s CookieStore) Del(key string) {
	cookie := &http.Cookie{
		MaxAge: -1,
		Name:   key,
		Path:   "/",
	}
	http.SetCookie(s.w, cookie)
}
