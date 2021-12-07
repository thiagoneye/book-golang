// Study on the Go language

// Book: Introducing Go
// Chapter 6: Functions

// Closure

package main
import "fmt"

func main() {
    add := func(x, y int) int {
        return x + y
    }
    fmt.Println(add(1,1))
}
