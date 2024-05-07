package myoss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
)

type (
	Client interface {
		PutObject(objectKey string, reader io.Reader, options ...oss.Option) (filename string, err error)
	}
)
