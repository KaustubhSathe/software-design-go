package main

import "fmt"

func main() {
	fmt.Println(helloWorld("Kaustubh"))
}


func helloWorld(name string) string {
	return fmt.Sprintf("Hi, %v. Welcome!", name)	
}