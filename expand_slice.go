package main

import "fmt"

func main() {
  data := []interface{} { 0, 1, 2, 3}
  desiredLen := 10

  expandedSlice := append(data, make([]interface{}, desiredLen - len(data))...)
  
  fmt.Println(len(expandedSlice))
  fmt.Println(expandedSlice)
}
