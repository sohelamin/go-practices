package main

import "fmt"

func main() {
    person := map[string]interface{} {
        "ID":    1,
        "Name":  "Sohel Amin",
        "Email": "sohelamincse@gmail.com",
    }

    fmt.Println(person["Name"])
}
