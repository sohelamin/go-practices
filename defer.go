package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {
    resp, err := http.Get("https://api.github.com/")

    if err != nil {
        // Handle the error
    }

    // Close once the reading has been done
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    fmt.Println(string(body))
}