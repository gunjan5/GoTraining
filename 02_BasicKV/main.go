package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting my KV store...")
	storage := map[string]string{}
	var operation string

	for {
		fmt.Println("\nEnter the operation [ GET | PUT | DEL | PRINT | EXIT ]:")
		fmt.Scanf("%s", &operation)

		switch operation {
		case "GET":
			fmt.Println("Enter the key:")
			var k string
			fmt.Scanf("%s", &k)
			_, ok := storage[k]
			if !ok {
				fmt.Printf("Key (%s) doesn't exist in the datastore\n", k)
				continue
			}
			fmt.Printf("GET: %s = %s\n", k, storage[k])
		case "PUT":
			fmt.Println("Enter the key:")
			var k, v string
			fmt.Scanf("%s", &k)
			fmt.Println("Enter the value:")
			fmt.Scanf("%s", &v)
			storage[k] = v
			fmt.Printf("Successfully PUT: %s = %s\n", k, v)
		case "DEL":
			fmt.Println("Enter the key:")
			var k string
			fmt.Scanf("%s", &k)
			_, ok := storage[k]
			if !ok {
				fmt.Printf("Key (%s) doesn't exist in the datastore\n", k)
				continue
			}
			fmt.Printf("DEL: %s = %s\n", k, storage[k])
			delete(storage, k)
		case "PRINT":
			for k, v := range storage {
				fmt.Printf("%s : %s\n", k, v)
			}
		case "EXIT":
			fmt.Println("Exiting. Bye o/")
			os.Exit(0)
		default:
			fmt.Println("invalid operation")
			os.Exit(1)
		}
	}

}
