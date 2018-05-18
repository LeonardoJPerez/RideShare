package auth

import (
	"encoding/base64"
	"net/http"
	"os"

	"github.com/RideShare-Server/log"
	"github.com/RideShare-Server/stores"
	"github.com/RideShare-Server/utils"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"github.com/juju/errors"
	"github.com/justinas/nosurf"
	"github.com/volatiletech/authboss"
	// To enable the auth and lock modules, they need to be imported
	_ "github.com/volatiletech/authboss/auth"
	_ "github.com/volatiletech/authboss/lock"
	aboauth "github.com/volatiletech/authboss/oauth2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	ab = authboss.New()
)

// buildCookieStore sets up the cookieStore
func buildCookieStore() {
	// Get the CookieStoreKey from the environment variables, and decode it
	encodedCookieKey := utils.GetEnvVariable(utils.AuthCookieStoreKey, "")
	if encodedCookieKey == "" {
		log.Error(log.AuthTopic, errors.Errorf("Empty Cookie Store Key"))
	}
	cookieStoreKey, err := base64.StdEncoding.DecodeString(encodedCookieKey)
	if err != nil {
		e := errors.Annotate(err, "Cookie Store Key Error")
		log.Error(log.AuthTopic, e)
	}
	cookieStore = securecookie.New(cookieStoreKey, nil)
}

// buildSessionCookieStore sets up the sessionStore
func buildSessionCookieStore() {
	// Get the SessionStoreKey from the environment variables, and decode it
	encodedSessionKey := utils.GetEnvVariable(utils.AuthSessionStoreKey, "")
	if encodedSessionKey == "" {
		log.Error(log.AuthTopic, errors.Errorf("Empty Session Store Key"))
	}
	sessionStoreKey, err := base64.StdEncoding.DecodeString(encodedSessionKey)
	if err != nil {
		e := errors.Annotate(err, "Session Store Key Error")
		log.Error(log.AuthTopic, e)
	}
	sessionStore = sessions.NewCookieStore(sessionStoreKey)
}

// getOAuth2Providers returns a map of providers to use with AuthBoss
func getOAuth2Providers() map[string]authboss.OAuth2Provider {
	return map[string]authboss.OAuth2Provider{
		"google": authboss.OAuth2Provider{
			OAuth2Config: &oauth2.Config{
				ClientID:     utils.GetEnvVariable(utils.GoogleAuthClientID, ""),
				ClientSecret: utils.GetEnvVariable(utils.GoogleAuthClientSecret, ""),
				Scopes:       []string{`profile`, `email`},
				Endpoint:     google.Endpoint,
			},
			Callback: aboauth.Google,
		},
	}
}

// SetupAuth sets up the auth package
func SetupAuth(db *gorm.DB) *authboss.Authboss {
	// Build the Cookie Store and Session Store
	buildCookieStore()
	buildSessionCookieStore()

	googleAuthDataStore := stores.NewGoogleAuthStore(db)

	ab.Storer = googleAuthDataStore
	ab.OAuth2Storer = googleAuthDataStore
	ab.MountPath = "/auth"

	serverPort := os.Getenv("APP_PORT")
	if serverPort == "" {
		serverPort = "8888"
	}

	ab.RootURL = utils.GetEnvVariable(utils.ApplicationRootURL, "http://localhost:"+serverPort)
	ab.AuthLoginOKPath = utils.GetEnvVariable(utils.AuthLoginOkPath, "")
	ab.AuthLogoutOKPath = utils.GetEnvVariable(utils.AuthLogoutOkPath, "")
	ab.OAuth2Providers = getOAuth2Providers()

	ab.XSRFName = "csrf_token"
	ab.XSRFMaker = func(_ http.ResponseWriter, r *http.Request) string {
		return nosurf.Token(r)
	}

	ab.CookieStoreMaker = NewCookieStore
	ab.SessionStoreMaker = NewSessionStore

	if err := ab.Init(); err != nil {
		log.Error(log.AuthTopic, err)
	}
	return ab
}
