package aws

import "mime/multipart"

type AwsS3Repository struct {
	AwsS3 AwsS3
}

func (awsS3 *AwsS3Repository) Upload(file multipart.File, fileName string, extension string) (url string, err error) {
	return awsS3.AwsS3.Upload(file, fileName, extension)
}
