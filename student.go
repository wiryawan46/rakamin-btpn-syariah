package main

import (
	"fmt"
	"os"
)

type Student struct {
	ID       string
	FullName string
}

func main() {
	students := []Student{
		{
			ID:       "1",
			FullName: "Student 1",
		},
		{
			ID:       "2",
			FullName: "Student 2",
		},
	}

	idStudent := os.Args[1]
	ShowStudent(students, idStudent)
}

func ShowStudent(students []Student, id string) {
	for i := range students {
		if students[i].ID == id {
			fmt.Printf("Studenet ID %s \n Student Name %s", students[i].ID, students[i].FullName)
		}
	}
}
