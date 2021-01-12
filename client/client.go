package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	for {
		resp, err := http.Get("http://localhost:3000/")
		if err != nil {
			fmt.Printf( "failed to reach server: %v\n", err)
			os.Exit(1)
		}

		body, err := ioutil.ReadAll(resp.Body)
		fmt.Printf("Response from server: %v\n", string(body))
		resp.Body.Close()
		if err != nil {
			fmt.Printf( "error: %v\n", err)
			os.Exit(1)
		}

		time.Sleep(10 * time.Second)
	}
}