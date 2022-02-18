package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	request, err := http.NewRequest("GET", "ws://localhost.com:80/", nil)
	if err != nil {
		log.Fatal(err)
	}

	httpClient := &http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 4096) // any non zero value will do, try '1'.
	for {
		n, err := response.Body.Read(buf)
		if n == 0 && err != nil { // simplified
			break
		}

		fmt.Printf("%s", buf[:n]) // no need to convert to string here
	}
	fmt.Println()
}
