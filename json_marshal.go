package main

import (
	"fmt"
  "encoding/json"
)

func main() {
  matrix := [][]float64 {
    {1.1, 1.5, 1.7},
    {2.1, 2.5, 2.7},
  }
  fmt.Println(matrix)
  // var lines []string
  result, err := json.Marshal(matrix)
  if err != nil {
    fmt.Println(err)
  }

  fmt.Println( string(result))
}
