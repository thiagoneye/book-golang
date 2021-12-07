// Study on the Go language

// Book: Introducing Go
// Chapter 3: Variables

// Code for Length Conversion

package main
import "fmt"

func main() {
    fmt.Println("Enter a length in feet: ")
    var input float64
    fmt.Scanf("%f", &input)

    output := input*0.3048
    fmt.Println("The length in meters is:")
    fmt.Println(output)
}
