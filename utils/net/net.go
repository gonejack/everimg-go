package net

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

func PostJSON(url string, json []byte, timeout int) (bs []byte, err error)  {
	client := &http.Client{
		Timeout:time.Duration(timeout) * time.Second,
	}

	resp, err := client.Post(url, "application/json", bytes.NewBuffer(json))

	if err == nil {
		bs, err = ioutil.ReadAll(resp.Body)

		if err == nil {
			err = resp.Body.Close()
		}
	}

	return
}