package aws

import (
	"errors"
	"fmt"

	"github.com/RideShare-Server/log"
	"github.com/RideShare-Server/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

var (
	clientID     = ""
	userPoolID   = ""
	jwkURLFormat = "https://cognito-idp.%v.amazonaws.com/%v/.well-known/jwks.json"
)

type (
	AuthResult struct {
		// The access token.
		AccessToken string `type:"string"`

		// The expiration period of the authentication result.
		ExpiresIn int64 `type:"integer"`

		// The ID token.
		IDToken string `type:"string"`

		// The new device metadata from an authentication result.
		NewDeviceMetadata *cognitoidentityprovider.NewDeviceMetadataType `type:"structure"`

		// The refresh token.
		RefreshToken string `type:"string"`

		// The token type.
		TokenType string `type:"string"`
	}
)

func init() {
	clientID = utils.GetEnvVariable("COGNITO_CLIENT_ID")
	userPoolID = utils.GetEnvVariable("COGNITO_USER_POOL_ID")

	if userPoolID == "" {
		log.Error(log.AuthTopic, errors.New("UserPoolID is missing"))

	}
	if clientID == "" {
		log.Error(log.AuthTopic, errors.New("ClientID is missing"))
	}
}

// Authenticate DONE!
func (s *Service) Authenticate(username, password string) (*AuthResult, error) {
	resp, err := s.initiateAuth(username, password)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cognitoidentityprovider.ErrCodeUserNotFoundException:
				return nil, errors.New("user not found")
			default:
				return nil, aerr
			}
		}
	}

	ch := aws.StringValue(resp.ChallengeName)
	if ch != "" {
		return nil, fmt.Errorf("Challenge required. [%s]", ch)
	}

	if resp.AuthenticationResult == nil {
		return nil, errors.New("Could not authenticate user")
	}

	result := &AuthResult{
		AccessToken:       aws.StringValue(resp.AuthenticationResult.AccessToken),
		ExpiresIn:         aws.Int64Value(resp.AuthenticationResult.ExpiresIn),
		IDToken:           aws.StringValue(resp.AuthenticationResult.IdToken),
		NewDeviceMetadata: resp.AuthenticationResult.NewDeviceMetadata,
		RefreshToken:      aws.StringValue(resp.AuthenticationResult.RefreshToken),
	}

	return result, err
}

// ValidateToken Done!
func (s *Service) ValidateToken(tokenString string) error {
	jwk := s.getJWT()
	token, err := validateToken(tokenString, s.Region, userPoolID, jwk)
	if err != nil || !token.Valid {
		return fmt.Errorf("token is not valid\n%v", err)
	}

	return nil
}

func (s *Service) challengeResponse(ch *cognitoidentityprovider.AdminInitiateAuthOutput) error {
	svc := cognitoidentityprovider.New(s.session, &aws.Config{Region: aws.String(s.Region)})

	challengeResp := map[string]*string{
		"USERNAME":     aws.String(""),
		"NEW_PASSWORD": aws.String(""),
	}

	params := &cognitoidentityprovider.RespondToAuthChallengeInput{
		ChallengeName:      ch.ChallengeName,
		ChallengeResponses: challengeResp,
		Session:            ch.Session,
		ClientId:           aws.String(clientID),
	}

	_, err := svc.RespondToAuthChallenge(params)

	return err
}

func (s *Service) getJWT() map[string]JWKKey {
	jwkURL := fmt.Sprintf(jwkURLFormat, s.Region, userPoolID)
	return getJWK(jwkURL)
}

func (s *Service) initiateAuth(username, password string) (*cognitoidentityprovider.AdminInitiateAuthOutput, error) {
	svc := cognitoidentityprovider.New(s.session, &aws.Config{Region: aws.String(s.Region)})
	params := &cognitoidentityprovider.AdminInitiateAuthInput{
		AuthFlow: aws.String("ADMIN_NO_SRP_AUTH"),
		AuthParameters: map[string]*string{
			"USERNAME": aws.String(username),
			"PASSWORD": aws.String(password),
		},
		ClientId:   aws.String(clientID),
		UserPoolId: aws.String(userPoolID),
	}

	return svc.AdminInitiateAuth(params)
}
