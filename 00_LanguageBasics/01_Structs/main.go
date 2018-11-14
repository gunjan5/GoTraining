package main

import "fmt"

// user represents a user in the system.
type user struct {
	name  string
	email string
	age   int
}

func main() {

	// Declare variable of type user and init using a struct literal.
	ivan := user{
		name:  "I-vaan",
		email: "iheartjavascript@gmail.com",
		age:   18,
	}

	// Display the field values.
	fmt.Println("Name", ivan.name)
	fmt.Println("Email", ivan.email)

	fmt.Printf("Type of Ivan: %T\n", ivan)

	// Declare a variable using an anonymous struct.
	gunjan := struct {
		name  string
		email string
		age   int
	}{
		name:  "Gunjan",
		email: "ihateprius@meep.org",
		age:   99,
	}

	// Display the field values.
	fmt.Println("Name", gunjan.name)
	fmt.Println("Email", gunjan.email)
	fmt.Println("Age", gunjan.age)
}
