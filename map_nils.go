package main

import "fmt"

func main() {
  someMap1 := map[int]int64 {
    7: 66,
    6: 2,
  }
  
  someMap2 := map[int64]int {
    0: 1,
    6: 7,
  }

  someArray := []int {
    6,
    7,
  }

  result := someMap1[someArray[someMap2[0]]]

  fmt.Println(result)
}
