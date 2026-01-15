package main

import "fmt"

func main() {
	var fruitList = []string{"apple","tomato","peach"}
	fmt.Printf("type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "mango","pomagrate")
	fmt.Printf("fruitList is %v\n", fruitList)
	fruitList = append(fruitList[1:]) // print from index 1 
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:3]) // print from index 1  to 3 (non inclusive)
	fmt.Println(fruitList)

	highscroes := make([]int,4)
	highscroes[0] = 234
	highscroes[1] = 945
	highscroes[2] = 465
	highscroes[3] = 867
	//highscroes[4] = 777 // this will throw error as the length of slice is 4
	highscroes = append(highscroes, 555,666,321)
	fmt.Println(highscroes)

	//how to remove a value from slice based on index
	var courses = []string{"reactjs","javascript","swift","python","ruby"}
	fmt.Println("courses before removing ", courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("courses after removing ", courses)
}

