package main

import "fmt"

const LoginTokin string = "ghghgh"// first word as capital ,Public


func main() {
	var userName string = "Ajay"
	fmt.Println(userName)
	fmt.Printf("variables is of type : %T \n", userName)

	var isLogin bool = false
	fmt.Println(isLogin)
	fmt.Printf("variables is of type : %T \n", isLogin)

	var smallVal uint8 = 255 //256 gives error as uint8 can hold values from 0 to 255
	fmt.Println(smallVal)
	fmt.Printf("variables is of type : %T \n", smallVal)

	var smallFloat float32 = 255.555 //float32 store 5 values after decimal and moreprecison in float64
	fmt.Println(smallFloat)
	fmt.Printf("variables is of type : %T \n",smallFloat)

	//default values and some aliases
	var anotherVariable int16
	fmt.Println(anotherVariable) //always initialized with 0 

	//implicit type ---> string(decided by value assigned but it does not allowed to change)
	var website = "example.com"
	fmt.Println(website)

	//no var style -> we used it inside the function only not outside the function
	noOfUsers := 250
	fmt.Println(noOfUsers)

	//const 
	fmt.Println(LoginTokin)
}