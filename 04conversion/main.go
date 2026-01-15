package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to pizza app")
	fmt.Println("please rate your pizza between 1 and 5 ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating",input)
	//conversion the string input to intger
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input),64)
	if err!=nil {
		fmt.Println("error is ",err)
	}else{
		fmt.Println("Adding to 1",numRating + 1)
	}
}