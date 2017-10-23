package main

import (
    "fmt"
    "time"
)

func SayHi() {
    fmt.Println("Goroutines Function")
}

func main() {
    go SayHi()
    time.Sleep(time.Second) // Delay the main function to show the goroutines has been executed
    fmt.Println("Main Function")
}
