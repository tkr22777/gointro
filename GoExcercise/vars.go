package main
import (
    "fmt"
    "strconv"
)

func main() {
    var x string = "So tired of hello world"
    fmt.Println(x)

    var title string
    title = "The denerate case of Mustafiz"
    fmt.Println(title)

    var aString string
    aString = "first"
    fmt.Println(aString)
    aString = "second"
    fmt.Println(aString)

    var y string = "bye"
    var z string = "world"

    fmt.Println( y == z )

    newString := "My damn new string in go lang"
    fmt.Println(newString)

    newNumber := 234234
    fmt.Println("This is a new number in go:" + strconv.Itoa(newNumber))

    i := 1
    for i < 10 {
        if i % 2 == 0 {
            fmt.Println(i, " even")
        } else {
            fmt.Println(i, " odd")
        }
        i = i + 1
    }
}
