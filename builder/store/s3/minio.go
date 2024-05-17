package s3

import (
	"caict.ac.cn/llm-server/common/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinio(cfg *config.Config) (*minio.Client, error) {
	return minio.New(cfg.S3.Endpoint, &minio.Options{
		Creds:        credentials.NewStaticV4(cfg.S3.AccessKeyID, cfg.S3.AccessKeySecret, ""),
		Secure:       true,
		BucketLookup: minio.BucketLookupAuto,
		Region:       cfg.S3.Region,
	})
}
