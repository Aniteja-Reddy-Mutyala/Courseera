package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// main function reads employee salary data from a file and computes
// the minimum, maximum, and average salaries along with employee names
func main() {
	// Open the salary data file
	file, err := os.Open("employees.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Map to store employee names as keys and salaries as values
	// Key format: "FirstName.LastName"
	employeeSalaries := make(map[string]int)

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		
		// Split the line into parts (firstName, lastName, salary)
		parts := strings.Fields(line)
		
		// Validate that we have exactly 3 parts
		if len(parts) != 3 {
			fmt.Println("Invalid line format:", line)
			continue
		}

		firstName := parts[0]
		lastName := parts[1]
		salaryStr := parts[2]

		// Convert salary string to integer
		salary, err := strconv.Atoi(salaryStr)
		if err != nil {
			fmt.Println("Error parsing salary for", firstName, lastName, ":", err)
			continue
		}

		// Create employee key by concatenating first and last name with a dot
		employeeKey := firstName + "." + lastName
		employeeSalaries[employeeKey] = salary
	}

	// Check for any errors during file reading
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Ensure we have data to process
	if len(employeeSalaries) == 0 {
		fmt.Println("No employee data found in file")
		return
	}

	// Initialize variables for finding min, max, and calculating average
	var minSalary, maxSalary, totalSalary int
	var employeeWithMinSalary, employeeWithMaxSalary string
	isFirstEmployee := true

	// Iterate through the map to find minimum, maximum, and sum of salaries
	for employeeName, salary := range employeeSalaries {
		// Initialize min and max with the first employee's salary
		if isFirstEmployee {
			minSalary = salary
			maxSalary = salary
			employeeWithMinSalary = employeeName
			employeeWithMaxSalary = employeeName
			isFirstEmployee = false
		}

		// Check if current salary is smaller than minimum
		if salary < minSalary {
			minSalary = salary
			employeeWithMinSalary = employeeName
		}

		// Check if current salary is larger than maximum
		if salary > maxSalary {
			maxSalary = salary
			employeeWithMaxSalary = employeeName
		}

		// Add to total for average calculation
		totalSalary += salary
	}

	// Calculate average salary
	numberOfEmployees := len(employeeSalaries)
	averageSalary := float64(totalSalary) / float64(numberOfEmployees)

	// Display the results
	fmt.Println("=== Company Salary Analysis ===")
	fmt.Printf("Total Employees: %d\n\n", numberOfEmployees)
	
	fmt.Printf("Employee with Lowest Salary: %s\n", 
		strings.Replace(employeeWithMinSalary, ".", " ", 1))
	fmt.Printf("Salary: $%d\n\n", minSalary)
	
	fmt.Printf("Employee with Highest Salary: %s\n", 
		strings.Replace(employeeWithMaxSalary, ".", " ", 1))
	fmt.Printf("Salary: $%d\n\n", maxSalary)
	
	fmt.Printf("Average Salary: $%.2f\n", averageSalary)
}