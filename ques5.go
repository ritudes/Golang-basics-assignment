package main

import "fmt"

func product(a int,b int,c int) int {
	var mul int
	mul = a * b * c
	return mul
}

func main(){

	fmt.Println(product(3,5,7))

}
