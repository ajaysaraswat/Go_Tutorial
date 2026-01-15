package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")
	languages := make(map[string]string)
	languages["Js"] = "Javascript"
	languages["Ts"] = "typescript"
	languages["ks"] = "kavascript"

	fmt.Println(languages)
	delete(languages,"ks") //delete from the map 
	fmt.Println(languages)

	//for each loop
	for key, value := range languages{
		fmt.Println(key,value)
	}
}