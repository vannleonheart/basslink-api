package common

import (
	"errors"
	"fmt"
	"mime/multipart"
)

func (s *Service) handleUploadFile(path string, form *multipart.Form) (*[]string, error) {
	result := make([]string, 0)

	for _, files := range form.File {
		if len(files) == 0 {
			return nil, errors.New("no file uploaded")
		}

		for _, file := range files {
			fileName := fmt.Sprintf("%s/%s", path, file.Filename)
			url, err := s.App.Storage.StorePublic(fileName, file)
			if err != nil {
				return nil, err
			}

			if url != nil {
				result = append(result, *url)
			}
		}
	}

	return &result, nil
}
