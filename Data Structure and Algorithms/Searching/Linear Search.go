package Searching

import "fmt"

/*
	This example Demonstrates Linear Search in a data list
	1. Define a linear search function
	2. This array takes an array of data and a key
	3. Loop through the data list
	4. repeat the process again for string 2
	5. print the result of the 2 strings
*/

func linearSearch(datalist []int, key int) bool {
	for _, item := range datalist {
		if item == key {
			return true
		}
	}
	return false
}

func main() {
	items := []int{95,78,46,58,45,86,99,251,320}
	fmt.Println(linearSearch(items,58))
}
