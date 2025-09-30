package basslink

import (
	"fmt"
	"mime/multipart"

	"github.com/minio/minio-go"
)

type StorageClient struct {
	Config StorageConfig
	Client *minio.Client
}

type StorageConfig struct {
	Endpoint  string `json:"endpoint"`
	ApiKey    string `json:"api_key"`
	ApiSecret string `json:"api_secret"`
	Bucket    string `json:"bucket"`
}

func NewStorageClient(config *StorageConfig) (*StorageClient, error) {
	cl, err := minio.New(config.Endpoint, config.ApiKey, config.ApiSecret, true)
	if err != nil {
		return nil, err
	}

	return &StorageClient{
		Config: *config,
		Client: cl,
	}, nil
}

func (osc *StorageClient) StorePublic(path string, file *multipart.FileHeader) (*string, error) {
	fileHandler, err := file.Open()
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = fileHandler.Close()
	}()

	_, err = osc.Client.PutObject(osc.Config.Bucket, path, fileHandler, file.Size, minio.PutObjectOptions{
		UserMetadata: map[string]string{
			"x-amz-acl": "public-read",
		},
		ContentType: file.Header.Get("Content-Type"),
	})
	if err != nil {
		return nil, err
	}

	result := fmt.Sprintf("https://%s.%s/%s", osc.Config.Bucket, osc.Config.Endpoint, path)

	return &result, nil
}
