// 조건
// 4년으로 나누어 떨어짐
    // 그 중, 100년으로 나누어 떨어짐 -> 평년
    // 그 중, 100년으로 나누어 떨어지지않음 -> 윤년
        // 그 중, 400년으로 나누어 떨어짐 -> 윤년

package main

import "fmt"

func main() {
    var year int
    
    fmt.Scan(&year)
    
    if year % 4 == 0 {
        if year % 100 == 0 {
            if year % 400 == 0 {
                fmt.Print(1)
            }else {
                fmt.Print(0)
            }
        }else {
            fmt.Print(1)
        }
    }else {
        fmt.Print(0)
    }
}