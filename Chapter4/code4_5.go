// Study on the Go language

// Book: Introducing Go
// Chapter 4: Control Structures

package main
import "fmt"

func main() {
    for i:= 1; i <= 100; i++ {
        if i % 3 == 0 {
            fmt.Println(i)
        }
    }
}
