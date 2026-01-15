package main

import "fmt"

func main() {
	fmt.Println("functions in golang")
	greet()
	result := adder(3,5)
	fmt.Println("result is ",result)
	fmt.Println(proAdder(2,3,4,5))
}
func greet(){
	fmt.Println("greeting from greet function")
}
func adder(a int,b int) int{
	return a+b;
	
}
func proAdder(values ...int)int {
	total := 0
	 for _,v := range values {
		total += v
	 }
	 return total
}