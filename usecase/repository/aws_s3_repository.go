package repository

import "mime/multipart"

type AwsS3Repository interface {
	Upload(file multipart.File, fileName string, extension string) (url string, err error)
}
