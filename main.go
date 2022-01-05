package main

import (
	"fmt"
	"strings"
)

// This is a very basic program to demonstrate a few things:
// Scanln - lets you take input from the user via the console
// StringBuilder - an efficient way to concatenate strings
// Slice manipulation in order to piece together our Pig Latin translation

func main() {

	// var then variable name then variable type
	var firstname string
	var lastname string

	fmt.Println("Enter First Name: ")

	// Taking input from user
	fmt.Scanln(&firstname)

	fmt.Println("Enter Last Name: ")

	// Taking input from user
	fmt.Scanln(&lastname)

	// StringBuilder helps put together our final output
	var newPigLatinName strings.Builder
	newPigLatinName.WriteString(firstname[1:])
	newPigLatinName.WriteString(firstname[0:1])
	newPigLatinName.WriteString("ay ")
	newPigLatinName.WriteString(lastname[1:])
	newPigLatinName.WriteString(lastname[0:1])
	newPigLatinName.WriteString("ay")

	// Print function is used to
	// display output in the same line
	fmt.Print("Your Full Name in Pig Latin is: " + strings.Title(strings.ToLower(newPigLatinName.String())))

}
