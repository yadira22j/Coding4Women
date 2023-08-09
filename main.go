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
	slice()
	maps()
	flowControl()

					// FOR LOOP:
	var accountBalance int = 2308
	fmt.Println("you have ", accountBalance, "in your account." ) 	
	account(accountBalance)
	
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
	
	arr := [4]int {3,7,11,17}
	var total int = arr[0] + arr[1] + arr[2] + arr[3]
		fmt.Println("The sum of", arr[0:4],  "is", (total))
  }
func slice(){
    //slices - array with unspecified length
  mySlice := []int {1,2,3,4,5}
  mySlice = append(mySlice, 6)
  fmt.Println(mySlice)
   //append to add to an array/slice
  mySlice = append(mySlice[0:2], mySlice[3:5]...) 
  // to remove a number from the slice
  fmt.Println(mySlice)
  
/* slices - are used more often because you want to be able to add or remove items  ar
arrays will be a specific number so if you change code will bring error more often*/
}

func maps (){
	
// MAPS - unordered collection of key- value pairs
//keys are the same and values are the same (keys all strings) (values all ints)
//map essentially a dictionary
//map[grapes:8 pies:3 teabags:20] //will print in random order (shorthand way of creating a map)

  var myMap = make(map[int]string) //creating a map
  fmt.Println(myMap) // will print empty

  var myMap2 map[int]string  // identify key and value key=string value = int
  myMap2 = map[int]string{3:"trees", 5: "lemons"}
  myMap2 [22] = "limes" // to append(add) to maps
  fmt.Println(myMap2)
  fmt.Println(myMap2[22]) // to print just 1 item of the map

}

func flowControl (){
/*uses booleans - true and false statements
 if it is true  outcome 1
 if it is false outcome 2 */

// if statement - will be execute if expression is true
// can add multiple expresions if multiple statements are true
//  if testExpression1 ||
//  || or && and
/*acountaBalance = 100
accountCredit = 200
if (accountBalance + accountCredit > 0{
	return true*/

// else would be creating the other path

/* ELSE IF - to check multiple if statements
if , else if, else if, else*/

/* for loops - execute multiple times until execution is false
counter set at 0, 
 i = 0
1 < 10
i++ - i plus 1   - increment by one
} 
*/
  sum := 0 
  for i:=0; i<10; i++ {
	  sum = sum + i }

    fmt.Println(sum)

	// for range loop - index and value
}

func account(funds int) { // same with bool for true or false statements
	if funds > 100000000 {
		fmt.Println("you're so rich!")
	} else if funds>= 100{
		fmt.Println("you can just about afford to survive the day!")
	} else {
		fmt.Println("peak!")
	}
	// account balance declared as variable in func main
	// account(accountBalance) in func main
	//fmt.Println("you have ", accountBalance, "in your account." )
}

