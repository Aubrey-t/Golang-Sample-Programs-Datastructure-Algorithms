package Basics

import "fmt"

func main(){
	var rows int
	var temp int = 1
	fmt.Print("Enter number of rows : ")
	fmt.Scan(&rows)

	for i := 0; i < rows; i++ {

		for j := 1; j <= rows-i ; j++ {
			fmt.Print(" ")
		}

		for k := 0; k <= i; k++ {

			if (k==0 || i==0) {
				temp = 1
			}else{
				temp = temp*(i-k+1)/k
			}

			fmt.Printf(" %d",temp)
		}
		fmt.Println("")

	}

}

/*
Output of Program :

Enter number of rows : 7
        1
       1 1
      1 2 1
     1 3 3 1
    1 4 6 4 1
   1 5 10 10 5 1
  1 6 15 20 15 6 1

*/
