package main

import (
	"encoding/json"
	"fmt"
)

type course struct{
	Name string
	Price int
	plateform string
	Password string
	Tags []string

}
func main() {
	fmt.Println("welcome to json in golang")

	lcoCourses := [] course{
		{"Reactjs",299,"lco.com","abc123",[]string{"webdev","ai"}},
		{"Reactjs",299,"lco.com","abc123",[]string{"webdev","ai"}},

	}
	//package the data as json data

	finalJson,err := json.MarshalIndent(lcoCourses,"","\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n",finalJson) 

	DecodeJson()
}

func DecodeJson(){
	jsonDataFromWeb := []byte(`
	[
        {
                "Name": "Reactjs",
                "Price": 299,
                "Password": "abc123",
                "Tags": [
                        "webdev",
                        "ai"
                ]
        }
    ]
 `)
	var lcoCourses course
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("json was valid")
		json.Unmarshal(jsonDataFromWeb,&lcoCourses)
		fmt.Printf("%#v\n",lcoCourses)
	} else {
		fmt.Println("json was not valid")
	}

	//making the map and copy data 
	var myonlineData map[string] interface{}
	json.Unmarshal(jsonDataFromWeb , &myonlineData)
	fmt.Printf("%#v\n",myonlineData)

	for k,v := range myonlineData {
		fmt.Printf("key is %v and value is %v and type is %T",k,v,v)
	}
}