package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	fmt.Println("wekcome to get request")
	http.HandleFunc("/post",handler)
//start server in go routine 
	go func(){
		err := http.ListenAndServe(":8000",nil)
		if err!= nil {
			panic(err)
		}
	}()
	///small delay to server start
	time.Sleep(1 * time.Second)

	//PerformGetRequest()
	PerformPostJsonRequest()
	
}

func PerformGetRequest(){
	const myurl = "http://localhost:8000/get"
	responce,err := http.Get(myurl)
	if err!= nil {
		panic(err)
	}
	defer responce.Body.Close()
	fmt.Println("response status is ",responce.StatusCode)

	result,_ := ioutil.ReadAll(responce.Body)
	fmt.Println("data from repsoncse is ",result)
	
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	//fake json payload
	requestBody := strings.NewReader(`
	{
	"coursename" : "lets go with golang" ,
	"price" : 0,
	"platefrom" : "lcoOnline.com"
	}
	`)
	responce,err := http.Post(myurl,"application/json",requestBody)
	if err!= nil {
		panic(err)
	}
	defer responce.Body.Close()
	result,_ :=ioutil.ReadAll(responce.Body)
	fmt.Println(result)
}

func PerformPostFormRequest(){
	const myurl = "http://localhost:8000/post"
	data := url.Values{}
	data.Add("firstName","hitesh")
	data.Add("lastName","choudhary")
	data.Add("email","hitesh@go.dev")

	responce,err := http.PostForm(myurl,data)
	if err!= nil {
		panic(err)
	}

	content ,_ := ioutil.ReadAll(responce.Body)
	defer responce.Body.Close()
	fmt.Println(string(content))

}
func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"Hello from Go server")
}
