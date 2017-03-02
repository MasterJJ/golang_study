package main

import "fmt"
import "math"

func main() {
    var a float64 = 10.0

        for i := 0; i< 10; i++ {
            a = a - 0.1
        }

    fmt.Println(a) // 9.00000000000

        const epsilon = 1e-14
        fmt.Println(math.Abs(a-9.0) <= epsilon)


}
