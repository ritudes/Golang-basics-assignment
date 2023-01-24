package main

import "fmt"

func main() {

	const (
		Nov = 11
		Oct = 10
		Sep = 9
	)
	fmt.Println(Nov, Oct, Sep)
	
	const (
		Sep1 = (9 + iota)
		Oct1 
		Nov1
	)
	fmt.Println(Nov1, Oct1, Sep1)

}
