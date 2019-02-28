package net

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

func PostJSON(url string, json []byte, timeout int) (bs []byte, e error)  {
	client := &http.Client{
		Timeout:time.Duration(timeout) * time.Second,
	}

	resp, e := client.Post(url, "application/json", bytes.NewBuffer(json))

	if e == nil {
		bs, e = ioutil.ReadAll(resp.Body)

		if e == nil {
			e = resp.Body.Close()
		}
	}

	return
}