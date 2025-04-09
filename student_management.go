package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	ID      int
	Name    string
	Age     int
	Grade   float64
	Course  string
}

var students []Student
var nextID = 1

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nStudent Management System")
		fmt.Println("1. Add Student")
		fmt.Println("2. View All Students")
		fmt.Println("3. Update Student")
		fmt.Println("4. Delete Student")
		fmt.Println("5. Search Student")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")

		input, _ := reader.ReadString('\n')
		choice, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			addStudent(reader)
		case 2:
			viewAllStudents()
		case 3:
			updateStudent(reader)
		case 4:
			deleteStudent(reader)
		case 5:
			searchStudent(reader)
		case 6:
			fmt.Println("Exiting program...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func addStudent(reader *bufio.Reader) {
	fmt.Println("\nAdd New Student")

	fmt.Print("Enter student name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter student age: ")
	ageInput, _ := reader.ReadString('\n')
	age, err := strconv.Atoi(strings.TrimSpace(ageInput))
	if err != nil {
		fmt.Println("Invalid age. Please enter a number.")
		return
	}

	fmt.Print("Enter student grade: ")
	gradeInput, _ := reader.ReadString('\n')
	grade, err := strconv.ParseFloat(strings.TrimSpace(gradeInput), 64)
	if err != nil {
		fmt.Println("Invalid grade. Please enter a number.")
		return
	}

	fmt.Print("Enter student course: ")
	course, _ := reader.ReadString('\n')
	course = strings.TrimSpace(course)

	student := Student{
		ID:     nextID,
		Name:   name,
		Age:    age,
		Grade:  grade,
		Course: course,
	}

	students = append(students, student)
	nextID++
	fmt.Println("Student added successfully!")
}

func viewAllStudents() {
	fmt.Println("\nAll Students")
	if len(students) == 0 {
		fmt.Println("No students found.")
		return
	}

	for _, student := range students {
		fmt.Printf("ID: %d, Name: %s, Age: %d, Grade: %.2f, Course: %s\n", 
			student.ID, student.Name, student.Age, student.Grade, student.Course)
	}
}

func updateStudent(reader *bufio.Reader) {
	fmt.Print("\nEnter student ID to update: ")
	idInput, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idInput))
	if err != nil {
		fmt.Println("Invalid ID. Please enter a number.")
		return
	}

	var found *Student
	for i := range students {
		if students[i].ID == id {
			found = &students[i]
			break
		}
	}

	if found == nil {
		fmt.Println("Student not found.")
		return
	}

	fmt.Printf("Updating student %s (leave blank to keep current value):\n", found.Name)

	fmt.Print("Enter new name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if name != "" {
		found.Name = name
	}

	fmt.Print("Enter new age: ")
	ageInput, _ := reader.ReadString('\n')
	ageInput = strings.TrimSpace(ageInput)
	if ageInput != "" {
		age, err := strconv.Atoi(ageInput)
		if err != nil {
			fmt.Println("Invalid age. Age not updated.")
		} else {
			found.Age = age
		}
	}

	fmt.Print("Enter new grade: ")
	gradeInput, _ := reader.ReadString('\n')
	gradeInput = strings.TrimSpace(gradeInput)
	if gradeInput != "" {
		grade, err := strconv.ParseFloat(gradeInput, 64)
		if err != nil {
			fmt.Println("Invalid grade. Grade not updated.")
		} else {
			found.Grade = grade
		}
	}

	fmt.Print("Enter new course: ")
	course, _ := reader.ReadString('\n')
	course = strings.TrimSpace(course)
	if course != "" {
		found.Course = course
	}

	fmt.Println("Student updated successfully!")
}

func deleteStudent(reader *bufio.Reader) {
	fmt.Print("\nEnter student ID to delete: ")
	idInput, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idInput))
	if err != nil {
		fmt.Println("Invalid ID. Please enter a number.")
		return
	}

	for i, student := range students {
		if student.ID == id {
			students = append(students[:i], students[i+1:]...)
			fmt.Println("Student deleted successfully!")
			return
		}
	}

	fmt.Println("Student not found.")
}

func searchStudent(reader *bufio.Reader) {
	fmt.Println("\nSearch Student")
	fmt.Println("1. Search by ID")
	fmt.Println("2. Search by Name")
	fmt.Print("Enter your choice: ")

	input, _ := reader.ReadString('\n')
	choice, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}

	switch choice {
	case 1:
		fmt.Print("Enter student ID: ")
		idInput, _ := reader.ReadString('\n')
		id, err := strconv.Atoi(strings.TrimSpace(idInput))
		if err != nil {
			fmt.Println("Invalid ID. Please enter a number.")
			return
		}

		for _, student := range students {
			if student.ID == id {
				fmt.Printf("Found student: ID: %d, Name: %s, Age: %d, Grade: %.2f, Course: %s\n", 
					student.ID, student.Name, student.Age, student.Grade, student.Course)
				return
			}
		}
		fmt.Println("Student not found.")

	case 2:
		fmt.Print("Enter student name: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		var found []Student
		for _, student := range students {
			if strings.Contains(strings.ToLower(student.Name), strings.ToLower(name)) {
				found = append(found, student)
			}
		}

		if len(found) == 0 {
			fmt.Println("No students found with that name.")
			return
		}

		fmt.Println("Found students:")
		for _, student := range found {
			fmt.Printf("ID: %d, Name: %s, Age: %d, Grade: %.2f, Course: %s\n", 
				student.ID, student.Name, student.Age, student.Grade, student.Course)
		}

	default:
		fmt.Println("Invalid choice.")
	}
}