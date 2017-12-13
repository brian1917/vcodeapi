package vcodeapi

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/brian1917/vcodeHMAC"
)

// Creates a new file upload http request with optional extra params
func uploadFile(credsFile, appID, sandboxID, path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	writer.WriteField("app_id", appID)
	writer.WriteField("sandbox_id", sandboxID)

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://analysiscenter.veracode.com/api/5.0/uploadfile.do", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Set authorization header and content-type header
	authHeader, err := vcodeHMAC.GenerateAuthHeader(credsFile, req.Method, req.URL.String())
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", authHeader)

	// Make HTTP Request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Process response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.Status != "200 OK" {
		return nil, errors.New("uploadfile.do call error: " + resp.Status)
	}

	// Return data and nil error
	return data, nil
}
