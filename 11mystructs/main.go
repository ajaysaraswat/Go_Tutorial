package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	//no inheritance in golang; no super parent

	hitesh := User{"Hitesh","hitesh@gmail.com",true,18}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details are: %+v\n",hitesh)
	fmt.Printf("Name is %v and email is %v\n",hitesh.Name , hitesh.Email)
}
//first letter capitalized means public 
type User struct {
	Name string
	Email string
	Status bool
	Age int
}