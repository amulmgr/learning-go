package main

//Example based off of : https://www.golang-book.com/books/intro/8#section1

import (
	"fmt"
)

//Calling a function with the argument (x) does not change the argument outside of the function.
//The variable is COPIED to the function
func changeValue_NO_Pointer(x int) {
	x = 0
}

//Pointers reference a location in memory where a value is stored
func changeValue_WITH_Pointer(x *int) {
	//By using the pointer we can modify the variable value
	*x = 0
}

func main() {
	x := 5
	//Calling this function will NOT change the value of x
	changeValue_NO_Pointer(x)
	//Should print 5
	fmt.Println(x)

	//The & operator returns a pointer to a variable
	//Calling this function with the address of x (&x). This function expects a pointer
	//to an int will change the value of x ouside of the function
	changeValue_WITH_Pointer(&x)

	//Should print 0
	fmt.Println(x)

	//You can also get a pointer to a variable using the new function
	px := new(int)

	//Set value pointed to to 7
	*px = 7
	//Should print 7
	fmt.Println(*px)

	//Change the value pointed to by px
	changeValue_WITH_Pointer(px)

	//Should print 0
	fmt.Println(*px)

}
