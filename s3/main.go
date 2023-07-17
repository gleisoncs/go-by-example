package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Data struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Create an instance of the Data struct and populate it with your data
	data := Data{
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}

	// Convert the Data struct to JSON format
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Create a session using your AWS credentials
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("sa-east-1"),
		// Add your AWS credentials here
	})
	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}

	// Create an instance of the S3 service
	svc := s3.New(sess)

	// Specify the bucket name and file name for your JSON file
	bucket := "acesso-rede-4789"
	key := "foi.json"

	// Upload the JSON data to S3
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(jsonData),
	})
	if err != nil {
		fmt.Println("Error uploading to S3:", err)
		return
	}

	fmt.Println("JSON file uploaded to S3 successfully.")
}
