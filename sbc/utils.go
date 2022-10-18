package sbc

import (
	"io/ioutil"
	"net/http"
)

func GetHttp(host string, path string) ([]byte, error) {
	resp, err := http.Get(host + path)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
