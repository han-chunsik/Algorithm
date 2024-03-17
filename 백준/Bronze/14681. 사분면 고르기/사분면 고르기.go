package main

import "fmt"

func main() {
    var a,b int
    
    fmt.Scanln(&a)
    fmt.Scanln(&b)
    
    if 0 < a {
        if 0 < b {
            fmt.Print(1)
        }else {
            fmt.Print(4)
        }
    }else {
        if 0 < b {
            fmt.Print(2)
        }else{
            fmt.Print(3)
        }
    }
}