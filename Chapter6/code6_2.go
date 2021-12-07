// Study on the Go language

// Book: Introducing Go
// Chapter 6: Functions

// Variadic Functions

package main
import "fmt"

func add(args ...int) int {
    total := 0
    for _, v := range args {
        total += v
    }
    return total
}

func main() {
    fmt.Println(add(1,2,3))
}
