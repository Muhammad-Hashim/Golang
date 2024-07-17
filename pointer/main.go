package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Hello, Golang!")
	// mynumber :=45
	// var ptr = &mynumber
	//  fmt.Println("Address of mynumber:", *ptr)
	var newry = []string{"Hello, Golang"}
	newry = append(newry, "one", "two", "three")
	fmt.Println("newry:", newry)

	var mymap = make([]int, 2)
	mymap[0] = 10
	mymap[1] = 20
	fmt.Println("mymap:", mymap[1])

}
