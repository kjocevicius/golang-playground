package main

import (
	"fmt"
  "encoding/json"
  "strconv"
  "math"
)

func main() {
  var value interface{} = "nan"
  stringValue, ok := value.(string)
  if ok {
    fmt.Println("It's ok")
  }
  var resultValue interface{}
  resultValue, err := strconv.ParseFloat(stringValue, 32)
  if err != nil {
    fmt.Println(err)
  }
  if math.IsNaN(resultValue.(float64)) {
    fmt.Println("It's a NaN")
    resultValue = nil
  }

  var anotherValue float64

  matrix := [][]interface{} {
    {1.1, 1.5, 1.7},
    {2.1, anotherValue, resultValue},
  }
  fmt.Println(matrix)
  // var lines []string
  result, err := json.Marshal(matrix)
  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(string(result))
}
