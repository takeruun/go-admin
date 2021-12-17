package aws

import "mime/multipart"

type AwsS3 interface {
	Upload(file multipart.File, fileName string, extension string) (url string, err error)
}
