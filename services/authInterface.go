package services

// AuthInterface :
type AuthInterface interface {
	// Authenticate :
	Authenticate(username, password string) (*AuthResult, error)
	// ValidateToken :
	ValidateToken(tokenString string) error
}
