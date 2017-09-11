package main

import "fmt" 

func changeme(myslice []string){
	myslice[0] = "Daniel"
	myslice[1] = "Campillo"
	fmt.Println("from changeme: ", myslice)
}

func changeme2(mymap map[string]int){
	mymap["Daniel"] = 25
	mymap["Luisa"] = 27
	fmt.Println("from changeme2: ", mymap)
}

func changeme3(myslice []string){
	myslice = append(myslice, "Daniel")
	myslice = append(myslice, "Campillo")
	fmt.Println("From changeme3: ", myslice)
}

func changeme4(myslice []string)[]string{
	myslice = append(myslice, "Daniel")
	myslice = append(myslice, "Campillo")
	fmt.Println("From changeme4: ", myslice)
	return myslice
}

func changeme5(myslice []string){
	myslice = append(myslice, "Daniel")
	myslice = append(myslice, "Campillo")
	fmt.Println("From changeme5: ", myslice)
}

func main(){
	//slice := make([]string, 1, 10) //initializes the slice with the first element to zero 	
	slice1 := make([]string, 2, 10) //with 1 theres is a chash
	fmt.Println(slice1)
	changeme(slice1)
	fmt.Println("from main: ", slice1)

	slice2 := make([]string, 0, 10) //with 1 theres is a chash
	fmt.Println(slice2)
	changeme3(slice2)
	fmt.Println("from main: ", slice2)

	slice3 := make([]string, 0, 10) //with 1 theres is a chash
	fmt.Println(slice3)
	slice3 = changeme4(slice3)
	fmt.Println("from main: ", slice3)

	slice4 := make([]string, 0, 10) //with 1 theres is a chash
	fmt.Println(slice4)
	changeme5(slice4)
	fmt.Println("from main 1: ", slice4)
	fmt.Println("from main 2: ", slice4[:cap(slice4)])

	mymap := make(map[string]int)
	fmt.Println(mymap)
	changeme2(mymap)
}
