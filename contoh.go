package main

import "fmt"

func main() {
	var fruits1 = []string{"apple", "banana", "mango", "durian", "pineapple"}
	fruits1 = append(fruits1[:3], "rambutan")
	fmt.Printf("%#v\n", fruits1)
}
