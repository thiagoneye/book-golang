// Study on the Go language

// Book: Introducing Go
// Chapter 5: Arrays, Slices and Maps

package main
import "fmt"

func main() {
    elements := map[string]map[string]string{
      "H" : map[string]string{
          "name": "Hydrogen",
          "state": "gas",
      },
      "He" : map[string]string{
          "name": "Helium",
          "state": "gas",
      },
      "Li" : map[string]string{
          "name": "Lithium",
          "state": "gas",
      },
      "Be" : map[string]string{
          "name": "Beryllium",
          "state": "gas",
      },
      "B" : map[string]string{
          "name": "Boron",
          "state": "gas",
      },
      "C" : map[string]string{
          "name": "Carbon",
          "state": "gas",
      },
      "N" : map[string]string{
          "name": "Nitrogen",
          "state": "gas",
      },
      "O" : map[string]string{
          "name": "Oxygen",
          "state": "gas",
      },
      "F" : map[string]string{
          "name": "Fluorine",
          "state": "gas",
      },
      "Ne" : map[string]string{
          "name": "Neon",
          "state": "gas",
      },
    }

    if el, ok := elements["Li"]; ok {
        fmt.Println(el["name"], el["state"])
    }
}
