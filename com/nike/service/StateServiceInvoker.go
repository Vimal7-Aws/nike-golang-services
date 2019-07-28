package service

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func InvokeStateService(state string) string {
	_ = state
	url := "http://localhost:7070/states"

	statesClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
	}

	res, getErr := statesClient.Do(req)
	if getErr != nil {
		log.Println(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Println(readErr)
	}

	return string(body)

}
