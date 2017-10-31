package main

import(
    "fmt"
    "encoding/json"
)

type Person struct {
    ID    int      `json:"id"`
    Name  string   `json:"name"`
    Email string   `json:"email"`
}

func main() {
    jsonSlice := []string{"Ford", "BMW", "Fiat"}
    jsonSliceString, _ := json.Marshal(jsonSlice)
    fmt.Println(string(jsonSliceString))
    // ["Ford","BMW","Fiat"]

    jsonMap := map[string]string{"name": "Sohel Amin", "email": "sohelamincse@gmail.com"}
    jsonMapString, _ := json.Marshal(jsonMap)
    fmt.Println(string(jsonMapString))
    // {"email":"sohelamincse@gmail.com","name":"Sohel Amin"}

    var emptyInterface map[string]interface{}
    jsonByte := []byte(`{"id":1,"name":"Sohel Amin","email":"sohelamincse@gmail.com"}`)
    json.Unmarshal(jsonByte, &emptyInterface)
    fmt.Println(emptyInterface)
    // map[email:sohelamincse@gmail.com id:1 name:Sohel Amin]

    jsonByte1 := []byte(`{"id":1,"name":"Sohel Amin","email":"sohelamincse@gmail.com"}`)
    person := Person{}
    json.Unmarshal(jsonByte1, &person)
    fmt.Println(person.Name)
    // Sohel Amin
}
