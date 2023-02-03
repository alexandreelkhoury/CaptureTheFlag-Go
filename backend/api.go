package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func callFirstAPI() string {
	re := regexp.MustCompile(`The secret key is: (.*)`)
	for i := 3000; i <= 4000; i++ {
		url := fmt.Sprintf("http://34.77.36.161:%d", i)
		resp, err := http.Get(url)
		if err == nil {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			fmt.Println("Received response on port", i)
			match := re.FindStringSubmatch(string(body))
			if len(match) > 1 {
				fmt.Println("Secret key:", match[1])
				//secondAPI(match[1])
				return match[1]

			}
		}
	}
	return "api not found"
}

func secondAPI(secretKey string) {
	url := "http://34.77.36.161:3941"
	data := []byte(fmt.Sprintf("secretKey=%s", secretKey))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response body:", string(body))
}

func thirdAPI() {
	url := "http://34.77.36.161:3610"
	data := []byte("finalKey=8116fdd3f12b6d7c4b136cbdaa3360a57eb4eb676ae63294450ee1f4f34b36f3")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response body:", string(body))
}
