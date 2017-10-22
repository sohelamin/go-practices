package main

import "fmt"

type Person struct {
    ID    uint
    Name  string
    Email string
}

func (p Person) GetID() uint {
    return p.ID
}

func (p Person) GetName() string {
    return p.Name
}

func (p Person) GetEmail() string {
    return p.Email
}

func (p *Person) SetID(id uint) bool {
    p.ID = id

    return true
}

func (p *Person) SetName(name string) bool {
    p.Name = name

    return true
}

func (p *Person) SetEmail(email string) bool {
    p.Email = email

    return true
}

func main() {
    person := Person{1, "Sohel Amin", "sohelamincse@gmail.com"}
    person.SetName("Md. Sohel Amin")

    fmt.Println(person.GetName())
}
