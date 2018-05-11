package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

const defaultRegion = "us-east-1"

// Service :
type BaseService struct {
	session *session.Session
	Region  string
}

// NewSession generates a new AWS Session.
func (s *BaseService) newSession() (*session.Session, error) {
	region := aws.String(s.Region)
	config := &aws.Config{
		Region: region,
		CredentialsChainVerboseErrors: aws.Bool(true),
	}

	return session.NewSession(config)
}
