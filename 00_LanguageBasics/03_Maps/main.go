package main

import "fmt"

func main() {

	// Declare and make a map of integer type values.
	departments := make(map[string]int)

	// Initialize some data into the map.
	departments["IT"] = 20
	departments["Marketing"] = 15
	departments["Executives"] = 5
	departments["Sales"] = 50
	departments["Security"] = 8

	// Display each key/value pair.
	for key, value := range departments {
		fmt.Printf("Dept: %s People: %d\n", key, value)
	}

	// Display map length.
	fmt.Printf("Len: %d\n", len(departments))

	// Check if a key exists in the map.
	if _, ok := departments["Engineering"]; !ok {
		fmt.Println("OMG, how is this company even functioning without engineers!")
	}

	// Add a KV pair.
	departments["Engineering"] = 42

	// Delete a KV pair.
	delete(departments, "Marketing")

	fmt.Println(departments)

}
