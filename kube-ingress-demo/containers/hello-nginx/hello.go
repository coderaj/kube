package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"net/http"
)

const hi_service = "nginx-service"

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello from service - Hello\n\n")
	fmt.Fprintf(w, "Calling Service - Hi\nGET http://"+hi_service+"/hi\n\n")

	req, err := http.NewRequest("GET", "http://"+hi_service+"/hi", nil)
	if err != nil {
		fmt.Printf("%s", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("%s", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Fprintf(w, string(body))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
