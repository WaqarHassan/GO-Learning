package main

import "fmt"

func main() {
	var age float64 = -10 * -2.1

	var (
		name    string  = "Waqar"
		section float32 = 12
	)

	var without = 10
	fmt.Println("Name = ", name)
	fmt.Println("Section = ", section)
	fmt.Println("Age = ", age)
	fmt.Println("Specifying data type is optional = ", without)
	fmt.Println("User Details = ", name, age, section)

}
