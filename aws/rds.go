package aws

import (
	"fmt"

	"github.com/RideShare-Server/utils"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/rds/rdsutils"
)

func (s *Service) GetRDSConnectionString() (string, error) {
	// arn := utils.GetEnvVariable("RDS_DB_ARN")
	endpoint := utils.GetEnvVariable("RDS_DB_ENDPOINT")
	dbName := utils.GetEnvVariable("RDS_DB_NAME")
	user := utils.GetEnvVariable("RDS_DB_USER")

	// awsCreds := stscreds.NewCredentials(s.session, arn)
	awsCreds := credentials.NewEnvCredentials()
	authToken, err := rdsutils.BuildAuthToken(endpoint, s.Region, user, awsCreds)
	if err != nil {
		// TODO: Log error
		fmt.Print(err.Error())
		return "", err
	}

	fmt.Println("User: ", user)
	fmt.Println("AuthToken: ", authToken)
	fmt.Println("Endpoint: ", endpoint)
	fmt.Println("DBName: ", dbName)

	// DNS: user:password@protocol(endpoint)/dbname?<params>
	//db, err := sql.Open("mysql", dnsStr)
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=true", user, authToken, endpoint, dbName), nil
}
