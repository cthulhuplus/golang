package main

import (
//        "fmt"
//        "io"
//        "log"
        "net/http"
//	"io/ioutil"
//        "os"
)

func main() {
  var client http.Client
  resp, err := client.Get("https://localhost:3680")
  if err != nil {
    // err
  }
  defer resp.Body.Close()

  if resp.StatusCode == 200 { // OK
//    fmt.Println(resp.Status)  // Print to standard out
    return
  }
}
