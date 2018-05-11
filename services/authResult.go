package services

// AuthResult :
type AuthResult struct {
	// The access token.
	AccessToken string `type:"string"`

	// The expiration period of the authentication result.
	ExpiresIn int64 `type:"integer"`

	// The ID token.
	IDToken string `type:"string"`

	// The refresh token.
	RefreshToken string `type:"string"`

	// The token type.
	TokenType string `type:"string"`

	// The Challenge type
	Challenge string `type:"string"`
}
