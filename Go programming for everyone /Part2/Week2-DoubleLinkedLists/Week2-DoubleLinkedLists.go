/*
 * Doubly Linked List Implementation for Palindrome Checking
 * 
 * This program implements a doubly linked list data structure that stores
 * characters and provides functionality to check if a sequence of characters
 * forms a palindrome. The doubly linked list allows bidirectional traversal,
 * making palindrome checking efficient by comparing characters from both ends.
 * 
 * Features:
 * - Add/delete elements from front and rear
 * - Insert/delete at specific positions
 * - Find elements by value
 * - Palindrome checking with optional case-insensitive comparison
 * - Complete test suite demonstrating all functions
 */

package main

import (
	"fmt"
	"strings"
)

// Node represents a single node in the doubly linked list
type Node struct {
	data rune  // Store character as rune for Unicode support
	prev *Node // Pointer to previous node
	next *Node // Pointer to next node
}

// DoublyLinkedList represents the doubly linked list structure
type DoublyLinkedList struct {
	head *Node // Pointer to first node
	rear *Node // Pointer to last node
	size int   // Number of elements in the list
}

// addToFront adds an element to the front of the list
func (dll *DoublyLinkedList) addToFront(val rune) {
	newNode := &Node{data: val, prev: nil, next: dll.head}

	if dll.head != nil {
		dll.head.prev = newNode
	} else {
		// List was empty, so rear should also point to new node
		dll.rear = newNode
	}

	dll.head = newNode
	dll.size++
}

// addToRear adds an element to the rear of the list
func (dll *DoublyLinkedList) addToRear(val rune) {
	newNode := &Node{data: val, prev: dll.rear, next: nil}

	if dll.rear != nil {
		dll.rear.next = newNode
	} else {
		// List was empty, so head should also point to new node
		dll.head = newNode
	}

	dll.rear = newNode
	dll.size++
}

// deleteFront removes the front element from the list
func (dll *DoublyLinkedList) deleteFront() bool {
	if dll.head == nil {
		return false // List is empty
	}

	if dll.head.next != nil {
		dll.head = dll.head.next
		dll.head.prev = nil
	} else {
		// Only one element in list
		dll.head = nil
		dll.rear = nil
	}

	dll.size--
	return true
}

// deleteRear removes the rear element from the list
func (dll *DoublyLinkedList) deleteRear() bool {
	if dll.rear == nil {
		return false // List is empty
	}

	if dll.rear.prev != nil {
		dll.rear = dll.rear.prev
		dll.rear.next = nil
	} else {
		// Only one element in list
		dll.head = nil
		dll.rear = nil
	}

	dll.size--
	return true
}

// findV returns pointer to first element with the given value, or nil if not found
func (dll *DoublyLinkedList) findV(val rune) *Node {
	current := dll.head

	for current != nil {
		if current.data == val {
			return current
		}
		current = current.next
	}

	return nil
}

// delete removes the first element with the given value
func (dll *DoublyLinkedList) delete(val rune) bool {
	node := dll.findV(val)
	if node == nil {
		return false // Value not found
	}

	// Update links
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		// Deleting head
		dll.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		// Deleting rear
		dll.rear = node.prev
	}

	dll.size--
	return true
}

// isEmpty returns true if the list is empty
func (dll *DoublyLinkedList) isEmpty() bool {
	return dll.head == nil
}

// findLength returns the number of elements in the list
func (dll *DoublyLinkedList) findLength() int {
	return dll.size
}

// insertPosition inserts an element after the ith position (0-indexed)
func (dll *DoublyLinkedList) insertPosition(i int, val rune) bool {
	if i < 0 || i >= dll.size {
		return false // Invalid position
	}

	if i == dll.size-1 {
		// Insert at rear
		dll.addToRear(val)
		return true
	}

	// Find the ith node
	current := dll.head
	for pos := 0; pos < i; pos++ {
		current = current.next
	}

	// Insert after current
	newNode := &Node{data: val, prev: current, next: current.next}
	if current.next != nil {
		current.next.prev = newNode
	}
	current.next = newNode
	dll.size++
	return true
}

// deletePosition deletes the element at the ith position (0-indexed)
func (dll *DoublyLinkedList) deletePosition(i int) bool {
	if i < 0 || i >= dll.size {
		return false // Invalid position
	}

	if i == 0 {
		return dll.deleteFront()
	}

	if i == dll.size-1 {
		return dll.deleteRear()
	}

	// Find the ith node
	current := dll.head
	for pos := 0; pos < i; pos++ {
		current = current.next
	}

	// Delete current
	current.prev.next = current.next
	current.next.prev = current.prev
	dll.size--
	return true
}

// isPalindrome checks if the list forms a palindrome
// caseInsensitive flag determines if 'a' should equal 'A'
func (dll *DoublyLinkedList) isPalindrome(caseInsensitive bool) bool {
	if dll.isEmpty() {
		return true
	}

	front := dll.head
	back := dll.rear

	// March from both ends towards the middle
	for front != back && front.prev != back {
		frontChar := front.data
		backChar := back.data

		// Convert to lowercase if case-insensitive
		if caseInsensitive {
			frontLower := strings.ToLower(string(frontChar))
			backLower := strings.ToLower(string(backChar))
			if len(frontLower) > 0 {
				frontChar = rune(frontLower[0])
			}
			if len(backLower) > 0 {
				backChar = rune(backLower[0])
			}
		}

		if frontChar != backChar {
			return false
		}

		front = front.next
		back = back.prev
	}

	return true
}

// printList prints the list from front to rear
func (dll *DoublyLinkedList) printList() {
	current := dll.head
	fmt.Print("List: ")
	for current != nil {
		fmt.Printf("%c ", current.data)
		current = current.next
	}
	fmt.Println()
}

// printReverse prints the list from rear to front
func (dll *DoublyLinkedList) printReverse() {
	current := dll.rear
	fmt.Print("Reverse: ")
	for current != nil {
		fmt.Printf("%c ", current.data)
		current = current.prev
	}
	fmt.Println()
}

// loadString loads a string into the doubly linked list
func (dll *DoublyLinkedList) loadString(str string) {
	for _, char := range str {
		dll.addToRear(char)
	}
}

// testAddToFront tests the addToFront function
func testAddToFront() {
	fmt.Println("=== Testing addToFront ===")
	dll := DoublyLinkedList{}
	dll.addToFront('C')
	dll.addToFront('B')
	dll.addToFront('A')
	dll.printList() // Should print: A B C
	fmt.Printf("Length: %d\n\n", dll.findLength())
}

// testAddToRear tests the addToRear function
func testAddToRear() {
	fmt.Println("=== Testing addToRear ===")
	dll := DoublyLinkedList{}
	dll.addToRear('A')
	dll.addToRear('B')
	dll.addToRear('C')
	dll.printList() // Should print: A B C
	fmt.Printf("Length: %d\n\n", dll.findLength())
}

// testDeleteFront tests the deleteFront function
func testDeleteFront() {
	fmt.Println("=== Testing deleteFront ===")
	dll := DoublyLinkedList{}
	dll.loadString("ABC")
	fmt.Print("Before: ")
	dll.printList()
	dll.deleteFront()
	fmt.Print("After deleting front: ")
	dll.printList() // Should print: B C
	fmt.Printf("Length: %d\n\n", dll.findLength())
}

// testDeleteRear tests the deleteRear function
func testDeleteRear() {
	fmt.Println("=== Testing deleteRear ===")
	dll := DoublyLinkedList{}
	dll.loadString("ABC")
	fmt.Print("Before: ")
	dll.printList()
	dll.deleteRear()
	fmt.Print("After deleting rear: ")
	dll.printList() // Should print: A B
	fmt.Printf("Length: %d\n\n", dll.findLength())
}

// testFindV tests the findV function
func testFindV() {
	fmt.Println("=== Testing findV ===")
	dll := DoublyLinkedList{}
	dll.loadString("ABCDE")
	node := dll.findV('C')
	if node != nil {
		fmt.Printf("Found 'C' in the list\n")
	} else {
		fmt.Printf("'C' not found\n")
	}
	node = dll.findV('Z')
	if node != nil {
		fmt.Printf("Found 'Z' in the list\n")
	} else {
		fmt.Printf("'Z' not found\n")
	}
	fmt.Println()
}

// testDelete tests the delete function
func testDelete() {
	fmt.Println("=== Testing delete ===")
	dll := DoublyLinkedList{}
	dll.loadString("ABCDE")
	fmt.Print("Before: ")
	dll.printList()
	dll.delete('C')
	fmt.Print("After deleting 'C': ")
	dll.printList() // Should print: A B D E
	fmt.Printf("Length: %d\n\n", dll.findLength())
}

// testInsertPosition tests the insertPosition function
func testInsertPosition() {
	fmt.Println("=== Testing insertPosition ===")
	dll := DoublyLinkedList{}
	dll.loadString("ACE")
	fmt.Print("Before: ")
	dll.printList()
	dll.insertPosition(0, 'B') // Insert 'B' after position 0
	fmt.Print("After inserting 'B' at position 0: ")
	dll.printList() // Should print: A B C E
	dll.insertPosition(2, 'D') // Insert 'D' after position 2
	fmt.Print("After inserting 'D' at position 2: ")
	dll.printList() // Should print: A B C D E
	fmt.Printf("Length: %d\n\n", dll.findLength())
}

// testDeletePosition tests the deletePosition function
func testDeletePosition() {
	fmt.Println("=== Testing deletePosition ===")
	dll := DoublyLinkedList{}
	dll.loadString("ABCDE")
	fmt.Print("Before: ")
	dll.printList()
	dll.deletePosition(2) // Delete position 2 (C)
	fmt.Print("After deleting position 2: ")
	dll.printList() // Should print: A B D E
	fmt.Printf("Length: %d\n\n", dll.findLength())
}

// testPalindromes tests palindrome detection with various strings
func testPalindromes() {
	fmt.Println("=== Testing Palindrome Detection ===")

	testCases := []struct {
		str              string
		caseInsensitive  bool
		expectedResult   bool
	}{
		{"otto", false, true},
		{"OtTo", false, false},
		{"OtTo", true, true},
		{"racecar", false, true},
		{"RaceCar", true, true},
		{"hello", false, false},
		{"A", false, true},
		{"AA", false, true},
		{"AB", false, false},
		{"ABA", false, true},
		{"Madam", true, true},
		{"MadaM", false, true}, // M-a-d-a-M is a palindrome
		{"MadDam", false, false}, // This is NOT a palindrome
	}

	for _, tc := range testCases {
		dll := DoublyLinkedList{}
		dll.loadString(tc.str)
		result := dll.isPalindrome(tc.caseInsensitive)
		status := "✓"
		if result != tc.expectedResult {
			status = "✗"
		}
		fmt.Printf("%s \"%s\" (case-insensitive: %v) -> %v (expected: %v)\n",
			status, tc.str, tc.caseInsensitive, result, tc.expectedResult)
	}
	fmt.Println()
}

// testIsEmpty tests the isEmpty function
func testIsEmpty() {
	fmt.Println("=== Testing isEmpty ===")
	dll := DoublyLinkedList{}
	fmt.Printf("Empty list isEmpty(): %v\n", dll.isEmpty())
	dll.addToFront('A')
	fmt.Printf("List with one element isEmpty(): %v\n", dll.isEmpty())
	dll.deleteFront()
	fmt.Printf("After deleting last element isEmpty(): %v\n\n", dll.isEmpty())
}

// testBidirectionalTraversal tests forward and backward traversal
func testBidirectionalTraversal() {
	fmt.Println("=== Testing Bidirectional Traversal ===")
	dll := DoublyLinkedList{}
	dll.loadString("ABCDE")
	dll.printList()
	dll.printReverse()
	fmt.Println()
}

func main() {
	fmt.Println("Doubly Linked List Implementation - Comprehensive Testing\n")
	fmt.Println("=" + strings.Repeat("=", 60))
	fmt.Println()

	// Run all tests
	testAddToFront()
	testAddToRear()
	testDeleteFront()
	testDeleteRear()
	testFindV()
	testDelete()
	testInsertPosition()
	testDeletePosition()
	testIsEmpty()
	testBidirectionalTraversal()
	testPalindromes()

	// Interactive palindrome checker
	fmt.Println("=== Interactive Palindrome Checker ===")
	fmt.Println("Enter strings to check for palindromes (or 'quit' to exit):")

	for {
		var input string
		var caseInsensitiveStr string

		fmt.Print("\nEnter string: ")
		fmt.Scanln(&input)

		if input == "quit" {
			break
		}

		fmt.Print("Case insensitive comparison? (yes/no): ")
		fmt.Scanln(&caseInsensitiveStr)

		caseInsensitive := strings.ToLower(caseInsensitiveStr) == "yes"

		dll := DoublyLinkedList{}
		dll.loadString(input)

		if dll.isPalindrome(caseInsensitive) {
			fmt.Printf("✓ \"%s\" IS a palindrome", input)
		} else {
			fmt.Printf("✗ \"%s\" is NOT a palindrome", input)
		}

		if caseInsensitive {
			fmt.Println(" (case-insensitive)")
		} else {
			fmt.Println(" (case-sensitive)")
		}
	}

	fmt.Println("\nThank you for using the Doubly Linked List Palindrome Checker!")
}