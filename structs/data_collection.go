package main

import (
	"fmt"

	"github.com/slashpai/go-refresher/structs/contact"
)

func main() {
	person := &contact.Person{}
	address := &contact.Address{}
	address.AddAddress("8/4305", "M.G Road", "Bangalore", "India", 560001)
	person.AddPerson("Jacob", "George", 33, "System Analyst", *address)
	fmt.Println(person.GetPerson())
}
