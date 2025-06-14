package main

import "fmt"

func main(){
	number:=10;
	if number > 0 {
		fmt.Printf(" %d is a positive number", number);
	} else if number < 0 {
		fmt.Printf(" %d is a negetive number", number);
	} else{
		fmt.Printf("zero");
	}
}