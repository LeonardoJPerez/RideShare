package aws

import (
	"fmt"

	"github.com/RideShare-Server/utils"
)

func (s *Service) GetDns() (string, error) {
	endpoint := utils.GetEnvVariable("RDS_DB_ENDPOINT")
	dbName := utils.GetEnvVariable("RDS_DB_NAME")
	user := utils.GetEnvVariable("RDS_DB_USER")
	pass := utils.GetEnvVariable("RDS_DB_PASSWRD")

	// // awsCreds := stscreds.NewCredentials(s.session, arn)
	// awsCreds := credentials.NewEnvCredentials()
	// authToken, err := rdsutils.BuildAuthToken(endpoint, s.Region, user, awsCreds)
	// if err != nil {
	// 	// TODO: Log error
	// 	fmt.Print(err.Error())
	// 	return "", err
	// }

	// DNS: user:password@protocol(endpoint)/dbname?<params>
	//db, err := sql.Open("mysql", dnsStr)
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pass, endpoint, dbName), nil
}
