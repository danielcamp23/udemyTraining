package main

import "fmt"

func pr(a int) (int, bool){
	b := a/2
	flag := a % 2 == 0
	return b, flag
}
func main(){
	/*var a string
	fmt.Println("insert you name:")
	fmt.Scan(&a)
	fmt.Printf("Hello %s\n", a)*/


	a,b := pr(11)
	fmt.Println(a, b)
}