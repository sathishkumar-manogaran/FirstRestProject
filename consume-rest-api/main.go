package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Starting the application...")
	getResponse()
	postResponse()
}

func getResponse() {
	response, err := http.Get("https://httpbin.org/ip")

	errorMessage(err, "The HTTP GET request failed with error")

	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println("GET Response :: \n", string(data))
}

func postResponse() {
	jsonData := map[string]string{"firstname": "Satz", "lastname": "Mano"}
	jsonValue, _ := json.Marshal(jsonData)
	buffer := bytes.NewBuffer(jsonValue)
	response, err := http.Post("https://httpbin.org/post", "application/json", buffer)

	errorMessage(err, "The HTTP POST request failed with error")

	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println("POST Response :: \n", string(data))
}

func errorMessage(err error, msg string) {
	if err != nil {
		fmt.Printf(msg+"%s\n", err)
	}
}
