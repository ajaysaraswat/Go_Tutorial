package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.google.com"
func main() {
	fmt.Println("welcome to web requests")
	responce, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("response status: %T", responce)
	defer responce.Body.Close() //always close the body after reading it 
    databytes,err := ioutil.ReadAll(responce.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println("data is ",content)

}