package main

import (
    "fmt"
)

func main() {

    x := [5]float64 { 98, 93, 77, 82, 83}

    var total float64 = 0

    for i := 0; i < len(x); i++ {
        total += x[i]
    }
    fmt.Println(total / float64(len(x)) )

    y := x[0:2]

    for i := 0; i < len(y); i++ {
        fmt.Printf("y[%d] = %g\n", i, y[i])
    }

    var stringToVal map[string]int = make(map[string]int)
    stringToVal["one"] = 1
    stringToVal["two"] = 2
    stringToVal["three"] = 3
    stringToVal["four"] = 4

    val, ok := stringToVal["blah"]
    fmt.Println(val, ok)

    if val, ok := stringToVal["two"]; ok {
        fmt.Println(val, ok)
    }

    //var stringToVal2 map[string]int = make(map[string]int)
    var stringToVal2 = map[string]int {
        "one" : 1,
        "two" : 2,
        "three" : 3,
        "four" : 4,
    }

    for numStr, numInt := range stringToVal2 {
        fmt.Println(numStr, numInt)
    }
}

