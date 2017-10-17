package vcodeapi

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/brian1917/vcodeHMAC"
)

func buildList(credsFile, appID string) ([]byte, error) {
	var errorMsg error

	client := http.Client{}
	req, err := http.NewRequest("GET", "https://analysiscenter.veracode.com/api/5.0/getbuildlist.do?app_id="+appID, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", vcodeHMAC.GenerateAuthHeader(credsFile, req.Method, req.URL.String()))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	if resp.Status != "200 OK" {
		errorMsg = errors.New("getbuildlist.do call error: " + resp.Status)
	}
	return data, errorMsg
}
