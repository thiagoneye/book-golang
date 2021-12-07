// Study on the Go language

// Book: Introducing Go
// Chapter 2: Types

package main

import "fmt"

func main() {
  fmt.Println("1 + 1 =", 1.0+1.0) // Operation with Numeric Type (Addition)
  fmt.Println(len("Hello, world")) // Operation with String (Length)
  fmt.Println("Hello, world"[1]) // Operation with String (Indexing)
  fmt.Println("Helo, " + "World") // Operation with String (Concatenation)
  fmt.Println(false && false) // Operation with Boolean (And)
  fmt.Println(true || false) // Operation with Boolean (Or)
  fmt.Println(!true) // Operation with Boolean (Or)
}
