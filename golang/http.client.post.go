package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	// LHST ...
	LHST = "http://127.0.0.1:3000/"
)

func httpDoPost() {

	// default client, http.DefaultClient
	// custom client
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(25 * time.Second)
				c, err := net.DialTimeout(netw, addr, time.Second*20)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
		},
	}

	// body
	form := url.Values{}
	form.Set("username", "gao")
	b := strings.NewReader(form.Encode()) // for string
	//b = bytes.NewReader([]byte(form.Encode())) // for bytes
	req, err := http.NewRequest(http.MethodPost, LHST+"answer", b)
	// header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Connection", "Keep-Alive")
	req.Header.Set("Content-Type", "application/json")
	// call
	res, err := client.Do(req)
	//res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
