package basslink

import (
	"io"

	"github.com/minio/minio-go"
)

type FileStoreClient struct {
	Config FileStoreConfig
}

type FileStoreConfig struct {
	Endpoint  string `json:"endpoint"`
	ApiKey    string `json:"api_key"`
	ApiSecret string `json:"api_secret"`
	Bucket    string `json:"bucket"`
}

func NewFileStoreClient(config *FileStoreConfig) *FileStoreClient {
	return &FileStoreClient{
		Config: *config,
	}
}

func (f *FileStoreClient) UploadFile(fileName string, fileContent io.Reader, fileSize int64, publicRead bool) error {
	cl, err := minio.New(f.Config.Endpoint, f.Config.ApiKey, f.Config.ApiSecret, true)
	if err != nil {
		return err
	}
	userMetadata := map[string]string{}
	if publicRead {
		userMetadata["x-amz-acl"] = "public-read"
	}
	_, err = cl.PutObject(f.Config.Bucket, fileName, fileContent, fileSize, minio.PutObjectOptions{
		UserMetadata: userMetadata,
	})
	if err != nil {
		return err
	}
	return nil
}
