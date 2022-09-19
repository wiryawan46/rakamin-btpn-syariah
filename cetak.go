package main

import "fmt"

func main() {
	// Cara 1
	userList := print("User 1", "User 2", "User 3")
	fmt.Printf("%v", userList)

	// Cara 2
	//userLists := []string{"User 1", "User 2", "User 3"}
	//userPrint := print(userLists...)
	//fmt.Printf("%v", userPrint)
}

// Variadic parameter
func print(names ...string) []map[string]string {
	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("user%d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}

	return result
}
