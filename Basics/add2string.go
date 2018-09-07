/*
	This example takes user input and add the 2 strings
	1. print enter message string
	2. allocate a place for the first input
	3. take the input string and place it in the variable
	4. repeat the process again for string 2
	5. print the result of the 2 strings
*/

package Basics

import "fmt"

func main(){
	fmt.Println("Enter First String")
	var first string
	fmt.Scanln(&first)
	fmt.Print("Enter Second String")
	var second string
	fmt.Scanln(&second)
	fmt.Print(first + second)
}