package vcodeapi

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

func sandboxList(username, password, appID string) ([]byte, error) {
	var errorMsg error = nil

	client := http.Client{}
	req, err := http.NewRequest("GET", "https://analysiscenter.veracode.com/api/5.0/getsandboxlist.do?app_id="+appID, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.Status != "200 OK" {
		errorMsg = errors.New("getsandboxlist.do call error: " + resp.Status)
	}
	return data, errorMsg

}
