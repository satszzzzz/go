package main

import (
    "fmt"
    "errors"
    "time"
    "sync"
)

func div(n int, d int) (int, int, error) {
    var err error
    if d == 0 {
        err = errors.New("Cannot divide by zero")
        return -1, -1, err
    }
    var result int = n / d
    remainder := n % d
    return result, remainder, nil
}

var dbData = []string{"db1", "db2", "db3", "db4", "db5"}
var results = []string{}

type person struct{
    age uint8
    isMale bool
}

func main() {
    var1 := "Hi"
    var var2 string = "World"
    res, rem, err := div(4, 2)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(var1 + "\n" + var2 + " " + fmt.Sprintf("%d", res))
    fmt.Println("Remainder:", rem)
    	
    var x []int32 = []int32{4,5,6}
    x[1]=69
    fmt.Println(x[1:2], x[0], "+", x)
    y := append(x, 7)
    fmt.Println(y, len(y), cap(y))

    var a []int32 = make([]int32, 3, 8)
    fmt.Println(a, len(a), cap(a))
    // map[string]uint8 -> key : value -> [key]value dtypes

    var mymap = make(map[int8]string)
    
    // Method 1: Add key-value pairs after initialization
    mymap[1] = "one"
    mymap[2] = "two"
    mymap[3] = "three"
    
    // Method 2: Initialize with values inline
    colors := map[string]string{
        "red":   "#FF0000",
        "green": "#00FF00",
        "blue":  "#0000FF",
    }
    
    fmt.Println("Map 1:", mymap)
    fmt.Println("Map 2:", colors)
    fmt.Println("Value at key 2:", mymap[2])

    for k,v := range mymap {
        fmt.Println("Key: ", k, "Value: ", v)
    }

    for k,v := range colors {
        fmt.Println("Key: ", k, "Value: ", v)
    }

    var temp1 = make(map[int8]int8)
    temp1[0]=8
    temp1[1]=8

    temp2 := map[string]string{
        "hehe":"haha",
        "hoho": "hihi",
    }

    fmt.Println(temp1, temp2)

    for ele := range temp1 {
        fmt.Println(ele, temp1[ele])
    }

    var i int=0

    for i<10 {
        fmt.Print(i, " ")
        i+=1
    }

    fmt.Println()

    for j:=0;j<10;j++ {
        fmt.Print(j," ")
    }
    fmt.Println()

    var satwik person = person{24, true}
    fmt.Println(satwik, "\nAGe: ", satwik.age, "Male: ", satwik.isMale)

    var amlan person = person{age: 24, isMale: true}
    fmt.Println(satwik ==  amlan)

    var z int32
    var p *int32
    z=6
    p=&z
    z=4
    fmt.Printf("Value at p: %d, Value of z: %d", *p, z)

    t0 := time.Now()
	var wg = sync.WaitGroup{}

	for i:=0; i<len(dbData); i++ {
		wg.Add(1)
		go dbCall(i, &wg)
	}
	wg.Wait()
	fmt.Println("Results: ", results)
    fmt.Println("Time taken: ", time.Since(t0))



}

func makeCHannel() {
    var c = make(chan int)
    c <- 1 // chan keyword to create a channel for holfing int, assigning 1 to channel c
    var i = <- c // channel c is empty
}

var m = sync.Mutex{}

func dbCall(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	delay:=2000
    m.Lock()
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println("Database called is ", dbData[i])
	results = append(results, dbData[i])
    m.Unlock()
}
