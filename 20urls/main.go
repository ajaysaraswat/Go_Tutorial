package main

import (
	"fmt"
	"net/url"
)
const mystring string = "https://www.google.com/search?q=golang+tutorials"

func main() {
	fmt.Println("welcome to the urls in golang")
	fmt.Println(mystring)

	//parsing the url
	result,err := url.Parse(mystring)
	if err!=nil {
		panic(err)
	}
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query() //store in map , key-value pair 
	
	fmt.Println(qparams["q"])

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "google.com",
		Path: "/search",
		RawQuery: "q=golang+tutorials",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}