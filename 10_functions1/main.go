package main

import "fmt"

func other() func() string{
	return func() string{
		return "I am anonymous too! :)"
	}
}

func main(){
	anonymous_function := func(){//this is called a func expression beacause it has no name
		fmt.Println("Call to anonymous func!")
	}

	anonymous2 := other()

	anonymous_function()
	fmt.Printf("%T\n", anonymous_function)

	fmt.Println(anonymous2())
	fmt.Printf("%T\n", anonymous2)

}