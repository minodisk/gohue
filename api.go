package gohue

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Fetch(method, url string, obj interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	json.Unmarshal(body, obj)
	return nil
}
