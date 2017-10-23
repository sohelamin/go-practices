package main

import "fmt"

func Receiver(c chan string) {
    message := <-c
    fmt.Println(message)
}

func main() {
    c := make(chan string)
    go Receiver(c)
    c <- "Hello Channel Function"

    fmt.Println("Main Function")
}
