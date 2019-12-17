package service

import (
	"bytes"
	"errors"
	log "github.com/sirupsen/logrus"
	"io"
	"khazen/config"
	"mime/multipart"
	"net/http"
	"os"
)

func UploadFile(fileKey string, fileName string, url string, title string) (err error) {
	log.Debugf("Start uploading file to %s",url)

	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(fileKey, fileName)
	if err != nil {
		return
	}
	_, err = io.Copy(part, file)

	_ = writer.WriteField("title", title)
	err = writer.Close()
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	if config.Config.UseSakkuUploadFileService{
		req.Header.Set("service", config.Config.SakkuUploadFile.Service)
		req.Header.Set("service-key", config.Config.SakkuUploadFile.ServiceKey)
	}

	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		return
	}

	if rsp.StatusCode != http.StatusOK {
		err = errors.New(rsp.Status)
	}
	return
}
