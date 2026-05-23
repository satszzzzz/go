package main

import (
	"fmt"
)

func makeCHannel() {
    var c = make(chan int)
    
    go process(c)  // Run process in goroutine to send values
    
    for value := range c {  // Receive values from channel
        fmt.Println(value)
    }
}

func process(c chan int){
	defer close(c)
	for i:=1;i<=4;i++{
		c <- i
	}
}

func main() {
	makeCHannel()
}