package basslink

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/vannleonheart/goutil"
	"github.com/vannleonheart/telegram-api-go"
	"os"
	"strings"
	"time"
)

const (
	FileStream     = "file"
	HttpStream     = "http"
	TelegramStream = "telegram"
)

type Stream struct {
	Type   string                 `json:"type"`
	Enable bool                   `json:"enable"`
	Config map[string]interface{} `json:"config"`
}

func WriteToFile(config map[string]interface{}, data interface{}) error {
	if data == nil {
		return errors.New("data is nil")
	}

	iPath, exist := config["path"]
	if !exist {
		return errors.New("file path is not exist")
	}

	strPath := iPath.(string)
	if len(strPath) <= 0 {
		return errors.New("file path is empty")
	}

	iFilename, exist := config["filename"]
	if !exist {
		return errors.New("filename is not exist")
	}

	strFilename := iFilename.(string)
	if len(strFilename) <= 0 {
		return errors.New("filename is empty")
	}

	strExtension := "txt"
	iExtension, exist := config["extension"]
	if exist {
		strIExtension := iExtension.(string)
		if len(strIExtension) > 0 {
			strExtension = strings.TrimLeft(strIExtension, ".")
		}
	}

	strRotation := "daily"
	iRotation, exist := config["rotation"]
	if !exist {
		strIRotation := iRotation.(string)
		if len(strIRotation) > 0 {
			strRotation = strIRotation
		}
	}

	switch strings.ToLower(strRotation) {
	case "hourly":
		strFilename = fmt.Sprintf("%s-%s", strFilename, time.Now().Format("2006-01-02-15"))
	case "daily":
		strFilename = fmt.Sprintf("%s-%s", strFilename, time.Now().Format("2006-01-02"))
	case "monthly":
		strFilename = fmt.Sprintf("%s-%s", strFilename, time.Now().Format("2006-01"))
	case "yearly":
		strFilename = fmt.Sprintf("%s-%s", strFilename, time.Now().Format("2006"))
	}

	strFilename = fmt.Sprintf("%s/%s.%s", strPath, strFilename, strExtension)

	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	fl, err := os.OpenFile(strFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer func() {
		_ = fl.Close()
	}()

	_, err = fl.WriteString(fmt.Sprintf("%s\n", string(byteData)))

	return err
}

func SendToHttp(config map[string]interface{}, data interface{}, headers *map[string]string) error {
	if data == nil {
		return errors.New("data is nil")
	}

	iUrl, exist := config["url"]
	if !exist {
		return errors.New("url is not exist")
	}

	strUrl := iUrl.(string)
	if len(strUrl) <= 0 {
		return errors.New("url is empty")
	}

	requestHeaders := map[string]string{}
	if headers != nil {
		requestHeaders = *headers
	}

	requestHeaders["Content-Type"] = "application/json"

	if _, err := goutil.SendHttpPost(strUrl, data, &requestHeaders, nil, nil); err != nil {
		return err
	}

	return nil
}

func PublishToTelegram(config map[string]interface{}, data interface{}) error {
	if data == nil {
		return errors.New("data is nil")
	}

	iBaseUrl, exist := config["base_url"]
	if !exist {
		return errors.New("base url is not exist")
	}

	iToken, exist := config["token"]
	if !exist {
		return errors.New("token is not exist")
	}

	iTo, exist := config["to"]
	if !exist {
		return errors.New("recipient is not exist")
	}

	recipients := iTo.([]interface{})
	if len(recipients) <= 0 {
		return errors.New("no recipient")
	}

	telegramClient := telegram.New(&telegram.Config{
		BaseUrl: iBaseUrl.(string),
		Token:   iToken.(string),
	})

	byMsg, err := json.Marshal(data)
	if err != nil {
		return err
	}

	for _, to := range recipients {
		if _, err := telegramClient.SendMessage(to.(string), string(byMsg), nil); err != nil {
			return err
		}
	}

	return nil
}
