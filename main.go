package main

import "fmt"

func multiply(a int, b int) int {
	return a * b
}

func main() {
	fmt.Println("Hello")
	/* capital "P" to print for all functions in the packages*/
	fmt.Println(multiply(3, 4))

	var c int = 2 
	var d int = 5 
	fmt.Println(c, "X", d, "= ", c * d)
	arrays() // everything can be seperated into different functions but must be called in the main function to be printed/executed 


	// constant - variable that doesnt change (const _ int)
}
func arrays() {
// array - sequence of specific length, same type inside square brackets */
	// square brackets tell you how many the array will hold
	// curly brackets what they are

													/* IF YOU DECLARE IT YOU MUST USE IT */
	numbers := [4]int {1,2,3,4}
	  fmt.Println(numbers)
	fruit := [4]string {"lemon","raspberry","passion fruit", "cherry"}
	//var fruit [4]string = [4]string{"lemon", "fruit", "passion fruit", "cherry"}
	//var numbers [4]int  = [4]int {1,2,3,4}                /*unspecified value will return 0000*/
	animals := [...]string {"dog", "hippo","cat", "hamster"}
	  fmt.Println(animals)

	/* unspecified number of things in array [...] will automatically determin the amount*/

	//index points position in the array starting at 0
	// to print part of the array
	fmt.Println(fruit[0:2]) /* will include before last number (0:2) will print 0,1 */
	fmt.Println(fruit[1] + ", " + fruit[2])
	//slices - array with unspecified length





	arr := [4]int {3,7,11,17}
	var total int = arr[0] + arr[1] + arr[2] + arr[3]
		fmt.Println("The sum of", arr[0:4],  "is", (total))
  }

