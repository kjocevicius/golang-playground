package main

import "fmt"

func main() {
  var arr = make([]string, 10)
  fmt.Println(arr)

  arr[0] = "labas"
  arr[7] = "krabas"
  // arr[10] = "krabas" //out of range
  fmt.Println(arr)

  arr = []string {
    "krabas",
    "labas",
    "namas",
  }
  mapOfArr := make(map[string][]string, len(arr))

  for _, v := range arr {
    mapOfArr[v] = make([]string, 10)
  }

  mapOfArr["krabas"][1] = "lazda"
  fmt.Println(mapOfArr["krabas"][1])
}
