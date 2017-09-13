package vcodeapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"errors"
)

func DetailedReport(username, password, build_id string) ([]byte, error) {
	var errorMsg error = nil

	client := http.Client{}
	req, err := http.NewRequest("GET", "https://analysiscenter.veracode.com/api/5.0/detailedreport.do?build_id="+build_id, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	if resp.Status != "200 OK" {
		errorMsg = errors.New("detailedreport.do call error: " + resp.Status)
	}
	return data, errorMsg
}
