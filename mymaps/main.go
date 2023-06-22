package main

import (
	"fmt"

)

func main() {

	fmt.Println("Maps in goLang")

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["jdk"] = "java"

	fmt.Println("List of all languages: ",languages)
	fmt.Println("JS: ",languages["js"])

	delete(languages, "jdk")
	fmt.Println("New list of langueges are : " , languages)


	// comma okay 
	if _,ok := languages["py"]; ok{
		fmt.Printf("%v value exists \n",ok)
	}else {
		fmt.Println("languages doesn't exist")
	}

	//loops in Maps

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}