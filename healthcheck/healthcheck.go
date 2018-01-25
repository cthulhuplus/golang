package main

import (
	"fmt"
	//        "io"
	//        "log"
	"net/http"
	//	"io/ioutil"
	//        "os"
)

func main() {
	var client http.Client
	resp, err := client.Get("http://localhost:9080/view/test")
	if err != nil {
		// err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 { // OK
		fmt.Println(resp.Status) // Print to standard out
		//    return
	}
}
