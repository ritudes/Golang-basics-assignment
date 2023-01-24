package main

import "fmt"

func main() {

	var age int
	fmt.Scan(&age)
	
	if age > 10 && age < 20 {
                fmt.Println("Younger")
        } else if age > 20 && age < 30 {
                fmt.Println("Adulthood")
        }else if age > 30 && age < 60 {
                fmt.Println("Getting wiser")
        }else if age > 60 {
                fmt.Println("Getting older")
        }  
}

