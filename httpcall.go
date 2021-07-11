package main

import (
	"net/http"
	"os"
)

//we need net pkg to http call

type webInfo struct {
	url string
}

func getRequest(web webInfo) (response *http.Response, exc error) {
	res, err := http.Get(web.url)
	if err != nil {
		os.Exit(1)
	}
	return res, err
}
