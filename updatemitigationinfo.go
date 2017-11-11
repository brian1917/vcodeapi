package vcodeapi

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/brian1917/vcodeHMAC"
)

func updateMitigationInfo(credsFile, buildID, action, comment, flawList string) ([]byte, error) {

	// Create HTTP form
	form := url.Values{}
	form.Add("build_id", buildID)
	form.Add("action", action)
	form.Add("comment", comment)
	form.Add("flaw_id_list", flawList)

	// Create HTTP client and request
	client := http.Client{}
	req, err := http.NewRequest("POST", "https://analysiscenter.veracode.com/api/updatemitigationinfo.do",
		strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	// Set authorization header and content-type header
	authHeader, err := vcodeHMAC.GenerateAuthHeader(credsFile, req.Method, req.URL.String())
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", authHeader)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make HTTP Request
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
		return nil, errors.New("updatemitigationinfo.do call error: " + resp.Status)
	}

	// Return data and nil error
	return data, nil
}
