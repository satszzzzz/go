package main

import (
    "fmt"
    "errors"
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
    
}
