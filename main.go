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
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	sess, err := session.NewSession()
	if err != nil {
		log.Fatalln(err.Error())
	}
	svc := s3.New(sess)

	inp := &s3.ListBucketsInput{}
	out, err := svc.ListBuckets(inp)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, b := range out.Buckets {
		log.Println(*b.Name)
	}

	obj := &s3.PutObjectInput{
		Body:                 aws.ReadSeekCloser(strings.NewReader("contents")),
		Bucket:               aws.String(os.Getenv("BUCKET_NAME")),
		Key:                  aws.String("test"),
		//ServerSideEncryption: aws.String("AES256"),
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
