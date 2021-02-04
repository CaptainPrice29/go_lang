package main

import (
	"fmt"
)
func digit(a int, c chan int){
	for a != 0{
		d := a%10
		c <- d
		a = a/10
	}
	close(c)
}
func main(){
	x := 45678
	ch := make(chan int)
	go digit(x,ch)
	for r := range ch{
		fmt.Println(r)
	}
}
