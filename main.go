package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	sess, err := session.NewSession()
	if err != nil {
		log.Fatalln(err.Error())
	}
	svc := s3.New(sess)

	obj := &s3.PutObjectInput{
		Body:                 aws.ReadSeekCloser(strings.NewReader("contents")),
		Bucket:               aws.String(os.Getenv("BUCKET_NAME")),
		Key:                  aws.String("test"),
		ServerSideEncryption: aws.String("AES256"),
	}
	result, err := svc.PutObject(obj)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
