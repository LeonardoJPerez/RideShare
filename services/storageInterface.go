package services

// StorageInterface defines a contract for a Storage provider.
type StorageInterface interface {
	// CreateBucket creates a bucket with the given bucketName.
	CreateBucket(bucketName string) error
	// DeleteBucket deletes a bucket from remote storage.
	DeleteBucket(bucketName string) error
	// ListBuckets returns a list of remote Buckets that match the given bucket name.
	ListBuckets(bucketName string) ([]string, error)
	// ListObjects fetches a list of keys(element names) within a remote bucket.
	ListObjects(bucketName string) ([]string, error)
	// Validate returns True if bucketName matches an existing and accessible remote bucket. Should return False if otherwise.
	Validate(bucketName string) (bool, error)
}
