package contact

import (
	"fmt"
)

// Address
type Address struct {
	houseNo string
	street  string
	city    string
	country string
	zipcode int64
}

// Person
type Person struct {
	firstName  string
	lastName   string
	age        int32
	occupation string
	address    Address
}

// AddAddress
func (address *Address) AddAddress(houseNo, street, city, country string, zipcode int64) {
	address.houseNo = houseNo
	address.street = street
	address.city = city
	address.country = country
	address.zipcode = zipcode
}

// GetAddress
func (address *Address) GetAddress() string {
	return fmt.Sprintf("%s %s %s %s %d", address.houseNo, address.street, address.city, address.country, address.zipcode)
}

// AddPerson
func (person *Person) AddPerson(firstName, lastName string, age int32, occupation string, address Address) {
	person.firstName = firstName
	person.lastName = lastName
	person.age = age
	person.occupation = occupation
	person.address = address
}

// GetPerson
func (person *Person) GetPerson() string {
	return fmt.Sprintf("%s %s %d %s %v", person.firstName, person.lastName, person.age, person.occupation, person.address.GetAddress())
}
