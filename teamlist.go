package vcodeapi

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/brian1917/vcodeHMAC"
)

func teamList(credsFile string) ([]byte, error) {
	var errorMsg error

	client := http.Client{}
	req, err := http.NewRequest("GET", "https://analysiscenter.veracode.com/api/3.0/getteamlist.do", nil)
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
		errorMsg = errors.New("getteamlist.do call error: " + resp.Status)
	}
	return data, errorMsg
}
