package tigris

import (
    "context"
    "fmt"
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/s3"
	"strings"
)

func Client(ctx context.Context) (*s3.Client, error) {
    // 1. Create an aws.Config instance
    cfg, err := config.LoadDefaultConfig(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed to load Tigris config: %w", err)
    }

    // 2. Create an Amazon S3 client using the AWS Config instance created, "cfg"
    return s3.NewFromConfig(cfg, func(o *s3.Options){
        o.BaseEndpoint = aws.String("https://fly.storage.tigris.dev")
        o.Region = "auto"
    }), nil
}

func UploadText( ctx context.Context ) (string, error){

	conn, err := Client( ctx )

	if err==nil{
		_, err = conn.PutObject(ctx, &s3.PutObjectInput{
			Bucket: aws.String("green-fog-5197"),
			Key:    aws.String("sample.txt"),
			Body:    strings.NewReader("testing content"),
		})
	}

	if err!=nil{
		return err.Error(), err
	}else{
		return "success",nil
	}
}