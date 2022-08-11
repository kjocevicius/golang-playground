package main

import "fmt"

func main() {
  var primary = int64(123456)
  var secondary = int64(789012)
  var shifted = secondary << 32
  var sum = shifted + primary
  fmt.Println(shifted)
  fmt.Println(sum)


  var primaryU = uint64(primary)
  var secondaryU = uint64(secondary)
  var shiftedU = secondaryU << 32
  var sumU = shiftedU + primaryU
  fmt.Println(shiftedU)
  fmt.Println(sumU)
}
