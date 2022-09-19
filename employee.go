package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func main() {
	var jsonString = `{"full_name" : "User 1", "email" : "user1@mail.com", "age" : 23}`
	//jsonString := `{"full_name" : "User 1", "email" : "user1@mail.com", "age" : 23}`

	var result Employee
	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	fmt.Println("Full name", result.FullName)
}
