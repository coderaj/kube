package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello from service - hi\n")
	out, _ := exec.Command("ping", "www.google.com", "-c 5", "-i 3", "-w 10").Output()
	if strings.Contains(string(out), "Destination Host Unreachable") {
		fmt.Fprintf(w, "www.google.com is down\n")
	} else {
		fmt.Fprintf(w, "www.google.com is reachable \n")
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
