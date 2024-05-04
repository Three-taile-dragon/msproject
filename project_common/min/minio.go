package min

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"strconv"
)

type MinioClient struct {
	c *minio.Client
}

func New(endpoint, accessKeyID, secretAccessKey string, useSSL bool) (*MinioClient, error) {
	// 初始化 Minio Client
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	return &MinioClient{
		c: minioClient,
	}, err
}

func (c *MinioClient) Put(
	ctx context.Context,
	bucketName string,
	fileName string,
	data []byte,
	fileSize int64,
	ContentType string,
) (minio.UploadInfo, error) {
	object, err := c.c.PutObject(ctx, bucketName, fileName, bytes.NewBuffer(data), fileSize, minio.PutObjectOptions{
		ContentType: ContentType,
	})
	return object, err
}

func (c *MinioClient) Compose(
	ctx context.Context,
	bucketName string,
	fileName string,
	TotalChunks int,
) (minio.UploadInfo, error) {
	dst := minio.CopyDestOptions{
		Bucket: bucketName,
		Object: fileName,
	}

	var srcs []minio.CopySrcOptions
	for i := 1; i < TotalChunks; i++ {
		// 分片编号
		formatInt := strconv.FormatInt(int64(i), 10)
		src := minio.CopySrcOptions{
			Bucket: bucketName,
			Object: fileName + "_" + formatInt,
		}
		srcs = append(srcs, src)
	}

	object, err := c.c.ComposeObject(ctx, dst, srcs...)
	return object, err
}

func (c *MinioClient) Del(
	ctx context.Context,
	bucketName string,
	fileName string,
	forceDelete bool,
) error {
	err := c.c.RemoveObject(ctx, bucketName, fileName, minio.RemoveObjectOptions{
		ForceDelete: forceDelete,
	})
	return err
}

func (c *MinioClient) DelS(
	ctx context.Context,
	bucketName string,
	filePrefix string,
	governanceBypass bool,
) error {
	objectsCh := make(chan minio.ObjectInfo)

	go func() {
		defer close(objectsCh)
		// List all objects from a bucket-name with a matching prefix.
		opts := minio.ListObjectsOptions{Prefix: filePrefix, Recursive: true}
		for object := range c.c.ListObjects(context.Background(), bucketName, opts) {
			if object.Err != nil {
				log.Println(object.Err)
			}
			objectsCh <- object
		}
	}()

	errorCh := c.c.RemoveObjects(ctx, bucketName, objectsCh, minio.RemoveObjectsOptions{
		GovernanceBypass: governanceBypass,
	})
	for e := range errorCh {
		log.Println("Failed to remove " + e.ObjectName + ", error: " + e.Err.Error())
		return e.Err
	}
	return nil
}

// TODO 后续可实现 秒传 功能

func (c *MinioClient) Get(
	ctx context.Context,
	bucket string,
	filename string) bool {
	object, err := c.c.GetObject(ctx, bucket, filename, minio.GetObjectOptions{})
	if err != nil {
		log.Println(err)
		return false
	}
	stat, err := object.Stat()
	if err != nil {
		log.Println(err)
		return false
	}
	return stat.Key != ""
}
