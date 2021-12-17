package infrastructure

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB struct {
		Production struct {
			Host     string
			Username string
			Password string
			DBName   string
		}
		Test struct {
			Host     string
			Username string
			Password string
			DBName   string
		}
	}
	Routing struct {
		Port string
	}
	AWS struct {
		S3 struct {
			Region          string
			Bucket          string
			AccessKeyID     string
			SecretAccessKey string
			Endpoint        string
		}
	}
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした。")
	}

	c := new(Config)

	c.DB.Production.Host = os.Getenv("DB_HOST")
	c.DB.Production.Username = os.Getenv("DB_USER")
	c.DB.Production.Password = os.Getenv("DB_PASSWORD")
	c.DB.Production.DBName = os.Getenv("DB_NAME")

	c.AWS.S3.Region = "ap-northeast-1"
	c.AWS.S3.Bucket = os.Getenv("S3_IMAGE_BUCKET")
	c.AWS.S3.AccessKeyID = os.Getenv("AWS_ACCESS_KEY_ID")
	c.AWS.S3.SecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	c.AWS.S3.Endpoint = os.Getenv("S3_ENDPOINT")

	c.DB.Test.Host = "localhost"
	c.DB.Test.Username = "username"
	c.DB.Test.Password = "password"
	c.DB.Test.DBName = "db_name_test"

	c.Routing.Port = ":3000"

	return c
}
