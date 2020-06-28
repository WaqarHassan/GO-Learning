package main

import "fmt"

func main() {
	var age float64 = -10 * -2.1

	var (
		name    string  = "Waqar"
		section float32 = 12
	)

	var isMale bool = true

	var without = 10
	fmt.Println("Name = ", name)
	fmt.Println("Section = ", section)
	fmt.Println("Age = ", age)
	fmt.Println("Is Male = ", isMale)
	fmt.Println("Specifying data type is optional = ", without)
	fmt.Println("User Details = ", name, age, "\nSection: ", section)

	const PI = 3.14159
	fmt.Println("Value of PI = ", PI)

	fmt.Println("PI = %i \n", PI)
	fmt.Printf("%0.2f", PI)

}
