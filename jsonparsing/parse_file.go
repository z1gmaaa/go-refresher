package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

// Users struct which contains
// an array of users
type Users struct {
	Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

// Social struct which contains a
// list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	// Open our jsonFile
	jsonFile := "users.json"

	// read our opened json as a byte array.
	byteValue, err := os.ReadFile(jsonFile)

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// we initialize our Users array
	var users Users

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println(users)
	// we iterate through every user within our users array and
	// print out the user Type, their name, and their social profile urls
	// as just an example

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
		fmt.Println("Twitter Handle: " + users.Users[i].Social.Twitter)
	}

	fmt.Println("\nTabular format\n----------")
	// Using third party lib to format
	data := [][]string{}
	for i := 0; i < len(users.Users); i++ {
		data = append(data, []string{users.Users[i].Type, strconv.Itoa(users.Users[i].Age), users.Users[i].Name, users.Users[i].Social.Facebook, users.Users[i].Social.Twitter})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Type", "Age", "Name", "Facebook Url", "Twitter Handle"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}
