package main

import "fmt"

func main(){
	records := make([][]string, 0, 10)
	record1 := []string{"Daniel", "Campillo", "24",}
	record2 := []string{"Luisa", "Villa", "27",}
	record3 := []string{"John", "Doe", "96",}
	records = append(records, record1)
	records = append(records, record2)
	records = append(records, record3)

	fmt.Println(records)

	otherway := new([100]int)[0:50]//Creates an array of 100 and takes 50
	//otherway := make([]int, 50, 100) equals to this
	fmt.Println(otherway)

	var myslice1 = []int{}
	fmt.Println("myslice1==nil",myslice1==nil) 

	var myslice2 []int
	fmt.Println("myslice2==nil",myslice2==nil) 	
}
