package main

import "fmt"

func main() {
	fmt.Println("welcome to control statements")
	loginCount := 20
	var result string
	if loginCount < 10 {
		result = "regular user"
	} else if loginCount > 10 {
		result = "watch out user"
	} else{
		result = "exactly 10 users"
	}

	fmt.Println(result)

	if num:=3; num<5 {
		fmt.Println("num is less than 5 ")
	}else{
		fmt.Println("num is not less than 5 ")
	}
}