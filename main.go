package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func Inputln() string {
	reader := bufio.NewScanner(os.Stdin)
	if reader.Scan() {
		return reader.Text()
	}
	if err := reader.Err(); err != nil {
		fmt.Println(err)
		return ""
	}
	return ""
}

type Student struct {
	Name string
	Age  int
}

var students []Student

func loadStudents() {
	file, err := os.Open("data.json")
	if err != nil {
		if os.IsNotExist(err) {
			students = []Student{}
			return
		}
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&students)
	if err != nil {
		panic(err)
	}
}

func saveStudents() {
	file, err := os.Create("data.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(students)
	if err != nil {
		panic(err)
	}
}

func addStudent() {
	fmt.Println("enter student name:")
	name := Inputln()
	fmt.Println("enter student age:")
	age, err := strconv.Atoi(Inputln())
	if err != nil {
		panic(err)
	}

	for _, s := range students {
		if s.Name == name {
			fmt.Println("student already exists.")
			return
		}
	}

	students = append(students, Student{Name: name, Age: age})
	saveStudents()
	fmt.Printf("added student %s aged %d to the system.\n", name, age)
}

func deleteStudent() {
	fmt.Println("name of student to delete:")
	name := Inputln()

	for i, s := range students {
		if s.Name == name {
			students = append(students[:i], students[i+1:]...)
			saveStudents()
			fmt.Printf("deleted student %s from the system.\n", name)
			return
		}
	}

	fmt.Println("student not found.")
}

func listStudents() {
	if len(students) == 0 {
		fmt.Println("no students...")
		return
	}

	for _, s := range students {
		fmt.Printf("name: %s, age: %d\n", s.Name, s.Age)
	}
}

func main() {
	loadStudents()

	fmt.Println("==== students ====")
	fmt.Println("1. add")
	fmt.Println("2. delete")
	fmt.Println("3. list")
	fmt.Println("4. quit")

	for {
		fmt.Println("your choice:")
		choice, err := strconv.Atoi(Inputln())
		if err != nil {
			fmt.Println("invalid input")
			continue
		}

		switch choice {
		case 1:
			addStudent()
		case 2:
			deleteStudent()
		case 3:
			listStudents()
		case 4:
			os.Exit(0)
		default:
			fmt.Println("invalid choice")
		}
	}
}
