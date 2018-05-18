package services

// AuthInterface :
type AuthInterface interface {
	// Authenticate :
	Authenticate(username, password string) (*AuthResult, error)

	ConfirmChangePassword(username, password, code string) (*AuthResult, error)

	ChangePassword(username string) (*AuthResult, error)

	CreateUser(username, password string) (*AuthResult, error)

	// ValidateToken :
	ValidateToken(tokenString string) error
}
