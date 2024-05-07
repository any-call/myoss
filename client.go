package myoss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"net/url"
)

type client struct {
	*oss.Client
	endpoint        string
	accessKey       string
	secretAccessKey string
	bucketName      string
}

func NewClient(endpoint, accessKey, secretAccessKey, bucketName string) (Client, error) {
	c, err := oss.New(endpoint, accessKey, secretAccessKey)
	if err != nil {
		return nil, err
	}

	return &client{
		Client:          c,
		endpoint:        endpoint,
		accessKey:       accessKey,
		secretAccessKey: secretAccessKey,
		bucketName:      bucketName,
	}, nil
}

func (self *client) PutObject(objectKey string, reader io.Reader, options ...oss.Option) (filename string, err error) {
	if self.Client == nil {
		return "", fmt.Errorf("client is nil")
	}

	bucket, err := self.Bucket(self.bucketName)
	if err != nil {
		if err = self.CreateBucket(self.bucketName); err != nil {
			return "", err
		}
		if bucket, err = self.Bucket(self.bucketName); err != nil {
			return "", err
		}
	}
	if err = bucket.PutObject(objectKey, reader, options...); err != nil {
		return "", err
	} else {
		err = bucket.SetObjectACL(objectKey, oss.ACLPublicRead)
	}

	filePath := fmt.Sprintf("%s/%s", self.endpoint, objectKey)
	url, err := url.Parse(filePath)
	if err != nil {
		return "", err
	}

	url.Host = self.bucketName + "." + url.Host
	return url.String(), err
}
