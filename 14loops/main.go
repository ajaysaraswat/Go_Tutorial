package main

import "fmt"


func main() {
	fmt.Println("welcome to loops in golang")
	days := []string{"sunday","moday","tuesday","wedbesday"}
	fmt.Println(days)
// 	for d := 0; d<len(days);d++ {
// 	fmt.Println(days[d])
//  }

	// for i:= range days { // i return the index not value
	// 	fmt.Println(days[i])
	// }

	for index,day := range days {
		fmt.Printf("index is %v and value is %v\n",index,day)
	}

	j := 1
	for j<10 {
		fmt.Println("value of j is ",j);
		j++;
	}
}	