package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"bytes"
)

func main() {
	client := &http.Client{}

	color := "red"
	breed := "pug"
	url := fmt.Sprintf("http://localhost:8087/dog?color=%v&breed=%v", color, breed)
	reqBody := []byte("hello")
	req, _ := http.NewRequest("GET", url, bytes.NewBuffer(reqBody))
	req.Header.Set("key", "value")

	fmt.Println(req)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
}
