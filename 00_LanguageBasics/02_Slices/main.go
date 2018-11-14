package main

import "fmt"

func main() {

	// Declare a nil slice of integers.
	var numbers []int

	// Append numbers to the slice.
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i*10)
	}

	// Display each value.
	for _, number := range numbers {
		fmt.Println(number)
	}

	// Declare a slice of strings.
	names := []string{"A", "B", "C", "D", "E", "F", "U"}

	// Display each index position and name.
	for i, name := range names {
		fmt.Printf("Index: %d  Name: %s\n", i, name)
	}

	fmt.Println("Print a sub-slice:")

	// Take a slice of index 1 and 2.
	slice := names[1:3]

	// Display the value of the new slice.
	for i, name := range slice {
		fmt.Printf("Index: %d  Name: %s\n", i, name)
	}

	diffNums := make([]int, 3, 5)
	fmt.Printf("Numbers: %v\n", diffNums)
	fmt.Printf("Size of diffNums: %d\n", len(diffNums))
	fmt.Printf("Capacity of diffNums: %d\n", cap(diffNums))
}
