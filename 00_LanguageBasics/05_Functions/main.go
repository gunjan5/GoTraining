package main

import "fmt"

// user represents a user in the system.
type user struct {
	name    string
	email   string
	active  bool
	commits []int
}

func main() {

	// Create a value of type user.
	u, err := createUser("bill", "thingy@things.io", []int{99, 54, 221, 69})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Display the value.
	fmt.Println(*u)

	u2, _ := createUser("kill", "killbill@things.io", []int{3, 5, 420})
	fmt.Println(u2)

	err = u2.deactivateUser()
	if err != nil {
		fmt.Printf("barf: %s\n", err)
	}
}

// createUser creates and returns pointers of user type values.
func createUser(name, email string, commits []int) (*user, error) {
	return &user{name, email, true, commits}, nil
}

// deactivateUser is a method to deactivate a user
func (u user) deactivateUser() error {
	u.active = false
	return nil
}

func (u *user) averageCommits() float64 {
	sum := 0
	for _, v := range u.commits {
		sum = sum + v
	}
	return float64(sum) / float64(len(u.commits))
}
