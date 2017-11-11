package vcodeapi

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/brian1917/vcodeHMAC"
)

func teamInfo(credsFile string, teamID string, includeUsers bool, includeApplications bool) ([]byte, error) {

	//Convert boolean values to API-expected strings
	includeUsersStr := "yes"
	includeAppsStr := "yes"
	if includeUsers == false {
		includeUsersStr = "no"
	}
	if includeApplications == false {
		includeAppsStr = "no"
	}

	// Create HTTP client and request
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://analysiscenter.veracode.com/api/3.0/getteaminfo.do?team_id="+teamID+"&include_users="+includeUsersStr+
		"&include_applications="+includeAppsStr, nil)
	if err != nil {
		return nil, err
	}

	// Set authorization header
	authHeader, err := vcodeHMAC.GenerateAuthHeader(credsFile, req.Method, req.URL.String())
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", authHeader)

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
		return nil, errors.New("getteaminfo.do call error: " + resp.Status)
	}

	// Return data and nil error
	return data, nil
}
