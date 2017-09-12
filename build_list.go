package vcodeapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func BuildList(username, password, app_id string) []byte {

	client := http.Client{}
	req, err := http.NewRequest("GET", "https://analysiscenter.veracode.com/api/5.0/getbuildlist.do?app_id="+app_id, nil)
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
	return data
}
