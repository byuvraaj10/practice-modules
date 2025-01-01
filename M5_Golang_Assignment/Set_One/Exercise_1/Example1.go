package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Department constants
const (
	HR      = "HR"
	IT      = "IT"
	FINANCE = "FINANCE"
	SALES   = "SALES"
)

type Employee struct {
	ID         int
	Name       string
	Age        int
	Department string
}

// Error definitions
var (
	ErrInvalidID        = fmt.Errorf("ID must be a positive unique number")
	ErrInvalidName      = fmt.Errorf("Name cannot be empty")
	ErrInvalidAge       = fmt.Errorf("Age must be greater than 18")
	ErrInvalidDepartment = fmt.Errorf("Department must be one of HR, IT, FINANCE, SALES")
)

// Utility function for validation
func validateEmployee(emp Employee, employees []Employee) error {
	if emp.ID <= 0 || containsID(emp.ID, employees) {
		return ErrInvalidID
	}
	if strings.TrimSpace(emp.Name) == "" {
		return ErrInvalidName
	}
	if emp.Age <= 18 {
		return ErrInvalidAge
	}
	validDepartments := map[string]bool{HR: true, IT: true, FINANCE: true, SALES: true}
	if !validDepartments[strings.ToUpper(emp.Department)] {
		return ErrInvalidDepartment
	}
	return nil
}

func containsID(id int, employees []Employee) bool {
	for _, e := range employees {
		if e.ID == id {
			return true
		}
	}
	return false
}

func promptInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func addEmployee(employees *[]Employee) {
	fmt.Println("\n=== Add New Employee ===")

	id, _ := strconv.Atoi(promptInput("Enter ID: "))
	name := promptInput("Enter Name: ")
	age, _ := strconv.Atoi(promptInput("Enter Age: "))
	dept := promptInput("Enter Department (HR/IT/FINANCE/SALES): ")

	emp := Employee{
		ID:         id,
		Name:       name,
		Age:        age,
		Department: strings.ToUpper(dept),
	}

	if err := validateEmployee(emp, *employees); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	*employees = append(*employees, emp)
	fmt.Println("Employee added successfully!")
}

func searchEmployee(employees []Employee) {
	fmt.Println("\n=== Search Employee ===")
	query := strings.ToLower(promptInput("Enter ID or Name to search: "))

	for _, emp := range employees {
		if strconv.Itoa(emp.ID) == query || strings.Contains(strings.ToLower(emp.Name), query) {
			printEmployee(emp)
			return
		}
	}
	fmt.Println("No matching employee found.")
}

func listEmployeesByDepartment(employees []Employee) {
	fmt.Println("\n=== List Employees by Department ===")
	dept := strings.ToUpper(promptInput("Enter Department (HR/IT/FINANCE/SALES): "))

	found := false
	for _, emp := range employees {
		if emp.Department == dept {
			printEmployee(emp)
			found = true
		}
	}
	if !found {
		fmt.Printf("No employees found in %s department.\n", dept)
	}
}

func printEmployee(emp Employee) {
	fmt.Printf("\nID: %d\nName: %s\nAge: %d\nDepartment: %s\n", emp.ID, emp.Name, emp.Age, emp.Department)
}

func countEmployeesByDepartment(employees []Employee) {
	fmt.Println("\n=== Employee Count by Department ===")
	count := map[string]int{}
	for _, emp := range employees {
		count[emp.Department]++
	}
	fmt.Printf("HR: %d, IT: %d, FINANCE: %d, SALES: %d\n", count[HR], count[IT], count[FINANCE], count[SALES])
}

func mainMenu() {
	fmt.Println("\n=== Employee Management System ===")
	fmt.Println("1. Add Employee")
	fmt.Println("2. Search Employee")
	fmt.Println("3. List Employees by Department")
	fmt.Println("4. Count Employees by Department")
	fmt.Println("5. Exit")
}

func main() {
	employees := []Employee{}
	for {
		mainMenu()
		choice := promptInput("Enter your choice: ")
		switch choice {
		case "1":
			addEmployee(&employees)
		case "2":
			searchEmployee(employees)
		case "3":
			listEmployeesByDepartment(employees)
		case "4":
			countEmployeesByDepartment(employees)
		case "5":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
