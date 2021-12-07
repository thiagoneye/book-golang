// Study on the Go language

// Book: Introducing Go
// Chapter 3: Variables

// Code for Temperature Conversion

package main
import "fmt"

func main() {
    fmt.Println("Enter a temperature in Fahrenheit: ")
    var input float64
    fmt.Scanf("%f", &input)

    output := (input - 32)*5/9
    fmt.Println("The temperature in Celsius is:")
    fmt.Println(output)
}
