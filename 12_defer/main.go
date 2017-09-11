package main

import "fmt"

func hello(){
	fmt.Print("Hello ")
}

func world(){
	fmt.Println("world!")
}

func main(){
	/*hello()
	defer world()*///works the same, defer applies on the function it is inside of
	defer world()
	hello()
}