package main

import (
  "fmt"
)

func main() {
  // Shorthand variable declaration, can only be used instead of a func
  a := 10
  b := "golang"
  c := 4.17
  d := true

  // Variable declaration, assignment = initialization, initialization doesn't need to be done at the same time that it is assigned

  var e string = "e variable assignment"

  fmt.Println("Shorthand Variable Declaration:")
  fmt.Println("a := 10 \t", a)
  fmt.Println("b := \"golang\" \t", b)
  fmt.Println("c := 4.17 \t", c)
  fmt.Println("d := true \t", d)
  fmt.Println("\nVariable Assignment/Initialisation:")
  fmt.Println("var e string = \"evariable assignment\" \t ", e)
}
