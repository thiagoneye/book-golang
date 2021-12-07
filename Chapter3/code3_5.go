// Study on the Go language

// Book: Introducing Go
// Chapter 3: Variables

// Code for Double the Input Value

package main
import "fmt"

func main() {
    fmt.Println("Enter a number: ")
    var input float64
    fmt.Scanf("%f", &input)

    output := 2*input
    fmt.Println("The output is ")
    fmt.Println(output)
}
