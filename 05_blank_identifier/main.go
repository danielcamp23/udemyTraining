package main

import "fmt"

const(
	j = iota
	k = iota
	l = iota
)

const(
	m = iota
	n
)
const p string = "esta es mi constante"
//var a string
func myfunc() (string, string){
	return "hola", "daniel"
}

func change(a *int, b int){
	*a += b
}

func main(){
	//a := "esto es a"
	//b := "estos es b"
	//var a string
	z := 10
	//var b string
	//a, _ = myfunc()
	/*fmt.Println(a)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(p)*/
	change(&z, 3)
	fmt.Printf("now z: %d\n", z)
}

// := cuando creo e inicializo.   = cuando ya declaré anteriormente
//var tales inicializa a zero
