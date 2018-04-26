package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

const defaultRegion = "us-east-1"

// Service :
type Service struct {
	session *session.Session
	Region  string
}

// New :
func New() *Service {
	service := new(Service)
	service.Region = defaultRegion
	session, err := service.newSession()
	if err != nil {
		// TODO: Log error.
		return nil
	}

	service.session = session

	return service
}

// NewSession generates a new AWS Session.
func (s *Service) newSession() (*session.Session, error) {
	region := aws.String(s.Region)
	config := &aws.Config{
		Region: region,
		CredentialsChainVerboseErrors: aws.Bool(true),
	}

	return session.NewSession(config)
}
