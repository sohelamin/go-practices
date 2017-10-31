package main

import (
    "fmt"
    "os"
    "io/ioutil"
)

func main() {
    // Create a file
    file, err := os.Create("test.txt")

    if err != nil {
        // handle the error
    }

    defer file.Close()

    file.WriteString("Test Line")

    // Open & read the file
    content, err := ioutil.ReadFile("test.txt")

    if err != nil {
        // handle the error
    }

    fmt.Println(string(content))
}
