package Searching

import "fmt"

// Binary Search implementation

/*
	1.Binary Search consistently reduce the amount of data to be searched and thereby increasing efficiency
	2.Define the 2 variables, low and high(arrays start from 0 hence the -1 on variable high)
	3.As long as we are in the range of the array of data
	4.add the 2 variables and divide by 2 to get the median
	5.if the median is less than the target value(needle), increase the starting point by 1 otherwise set high to median minus 1
	6.Note that this algorithm works only for sorted data
*/

func binarySearch(needle int, haystack []int) bool {

	low := 0
	high := len(haystack) - 1 // len equals to 8

	for low <= high {
		median := (low + high) / 2

		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}

		if low == len(haystack) || haystack[low] != needle {
			return false
		}
	}
	return true
}

func main(){
	items := []int{1,2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(63, items))
}