package aws

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Service struct {
	BaseService
}

// NewS3Service :
func NewS3Service() *S3Service {
	service := new(S3Service)
	service.Region = defaultRegion
	session, err := service.newSession()
	if err != nil {
		// TODO: Log error.
		return nil
	}

	service.session = session

	return service
}

// CloudName prepends the given string
func (s *S3Service) CloudName(env string) string {
	if strings.HasPrefix(env, "viro-") {
		return env
	}

	return fmt.Sprintf("viro-%s", strings.ToLower(env))
}

// CreateBucket creates a bucket with the given bucketName on S3
func (s *S3Service) CreateBucket(bucketName string) error {
	bucketName = s.CloudName(bucketName)

	// Build the bucket name with the `viro` prefix.
	input := &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	}

	// Create the actual bucket.
	svc := s3.New(s.session)
	_, err := svc.CreateBucket(input)
	if err != nil {
		return fmt.Errorf("failed to open file %q, %v", bucketName, err)
	}
	return nil
}

// DeleteBucket deletes a bucket from S3 storage.
func (s *S3Service) DeleteBucket(bucketName string) error {
	bucketName = s.CloudName(bucketName)
	svc := s3.New(s.session)
	input := &s3.DeleteBucketInput{
		Bucket: aws.String(bucketName),
	}

	_, err := svc.DeleteBucket(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			return aerr
		}
		return err
	}

	return nil
}

// ListBuckets returns a list of S3 Buckets that match the given bucketName
func (s *S3Service) ListBuckets(bucketName string) ([]string, error) {
	bucketName = s.CloudName(bucketName)
	input := &s3.ListBucketsInput{}

	svc := s3.New(s.session)
	output, err := svc.ListBuckets(input)
	if err != nil {
		return nil, fmt.Errorf("failed to open file, %v", err)
	}

	var matchingBuckets []string
	for _, bucket := range output.Buckets {
		cloudName := strings.ToLower(*bucket.Name)
		if strings.Contains(cloudName, bucketName) {
			matchingBuckets = append(matchingBuckets, cloudName)
		}
	}
	return matchingBuckets, nil
}

// ListObjects fetches a list of keys within a remote bucket.
func (s *S3Service) ListObjects(bucketName string) ([]string, error) {
	bucketName = s.CloudName(bucketName)
	svc := s3.New(s.session)
	input := &s3.ListObjectsInput{
		Bucket:  aws.String(bucketName),
		MaxKeys: aws.Int64(1000),
	}

	result, err := svc.ListObjects(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		}
		return nil, err
	}

	objects := []string{}
	for _, o := range result.Contents {
		objects = append(objects, *o.Key)
	}

	return objects, nil
}

// Validate returns True if bucketName matches an exiting and accessible bucket in AWS.
// Returns False if otherwise.
func (s *S3Service) Validate(bucketName string) (bool, error) {
	buckets, err := s.ListBuckets(s.CloudName(bucketName))
	return len(buckets) > 0, err
}
