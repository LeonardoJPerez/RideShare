package aws

import (
	"errors"
	"fmt"

	"github.com/RideShare-Server/log"
	"github.com/RideShare-Server/services"
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

type CognitoService struct {
	BaseService
}

// NewCognitoService :
func NewCognitoService() CognitoService {
	service := CognitoService{}
	service.Region = defaultRegion
	session, err := service.newSession()
	if err != nil {
		// TODO: Log error.
	}

	service.session = session

	return service
}

// CreateUser :
func (s CognitoService) CreateUser(username, password string) (*services.AuthResult, error) {
	svc := cognitoidentityprovider.New(s.session, &aws.Config{Region: aws.String(s.Region)})
	params := &cognitoidentityprovider.AdminCreateUserInput{
		TemporaryPassword: aws.String(password),
		Username:          aws.String(username),
		UserPoolId:        aws.String(userPoolID),
	}

	resp, err := svc.AdminCreateUser(params)
	if err != nil {

	}
	if resp.User == nil {
		return nil, errors.New("Could not authenticate user")
	}

	return nil, err
}

func (s CognitoService) ChangePassword(username string) (*services.AuthResult, error) {
	svc := cognitoidentityprovider.New(s.session, &aws.Config{Region: aws.String(s.Region)})

	params := &cognitoidentityprovider.ForgotPasswordInput{
		ClientId: aws.String(clientID),
		Username: aws.String(username),
	}

	resp, err := svc.ForgotPassword(params)
	utils.PrettyPrint(resp)

	return nil, err
}

func (s CognitoService) ConfirmChangePassword(username, password, code string) (*services.AuthResult, error) {
	svc := cognitoidentityprovider.New(s.session, &aws.Config{Region: aws.String(s.Region)})
	params := &cognitoidentityprovider.ConfirmForgotPasswordInput{
		ClientId:         aws.String(clientID),
		ConfirmationCode: aws.String(code),
		Password:         aws.String(password),
		Username:         aws.String(username),
	}

	resp, err := svc.ConfirmForgotPassword(params)
	utils.PrettyPrint(resp)

	return nil, err
}

// Authenticate :
func (s CognitoService) Authenticate(username, password string) (*services.AuthResult, error) {
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
		return &services.AuthResult{
			Challenge: ch,
		}, fmt.Errorf("Challenge required. [%s]", ch)
	}

	if resp.AuthenticationResult == nil {
		return nil, errors.New("Could not authenticate user")
	}

	result := &services.AuthResult{
		AccessToken: aws.StringValue(resp.AuthenticationResult.AccessToken),
		ExpiresIn:   aws.Int64Value(resp.AuthenticationResult.ExpiresIn),
		IDToken:     aws.StringValue(resp.AuthenticationResult.IdToken),
		// NewDeviceMetadata: resp.AuthenticationResult.NewDeviceMetadata,
		RefreshToken: aws.StringValue(resp.AuthenticationResult.RefreshToken),
	}

	return result, err
}

// ValidateToken :
func (s CognitoService) ValidateToken(tokenString string) error {
	jwk := s.getJWT()
	token, err := validateToken(tokenString, s.Region, userPoolID, jwk)
	if err != nil || !token.Valid {
		return fmt.Errorf("token is not valid\n%v", err)
	}

	return nil
}

// Private Methods
func (s CognitoService) challengeResponse(ch *cognitoidentityprovider.AdminInitiateAuthOutput) error {
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

func (s CognitoService) getJWT() map[string]JWKKey {
	jwkURL := fmt.Sprintf(jwkURLFormat, s.Region, userPoolID)
	return getJWK(jwkURL)
}

func (s CognitoService) initiateAuth(username, password string) (*cognitoidentityprovider.AdminInitiateAuthOutput, error) {
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
