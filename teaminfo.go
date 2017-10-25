package vcodeapi

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/brian1917/vcodeHMAC"
)

func teamInfo(credsFile string, teamID string, includeUsers bool, includeApplications bool) ([]byte, error) {
	var errorMsg error

	//CONVERT BOOL TO API ACCEPTED STRING
	includeUsersStr := "yes"
	includeAppsStr := "yes"
	if includeUsers == false {
		includeUsersStr = "no"
	}
	if includeApplications == false {
		includeAppsStr = "no"
	}

	client := http.Client{}
	req, err := http.NewRequest("GET", "https://analysiscenter.veracode.com/api/3.0/getteaminfo.do?team_id="+teamID+"&include_users="+includeUsersStr+
		"&include_applications="+includeAppsStr, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", vcodeHMAC.GenerateAuthHeader(credsFile, req.Method, req.URL.String()))

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.Status != "200 OK" {
		errorMsg = errors.New("getteaminfo.do call error: " + resp.Status)
	}
	return data, errorMsg
}
