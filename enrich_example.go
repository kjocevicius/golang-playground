package main

import (
	"fmt"
  "C"
  "strings"
	"strconv"
)

type Row struct {
	item_id   int
	something string
}

func enrich_data(input *[]int) []Row {
	var result []Row
	var new_row Row

	for _, v := range *input {
		new_row = Row{
			item_id:   int(v),
			something: "bla",
		}
		result = append(result, new_row)
	}

	return result
}

func main() {
  string_input := "1;2;3645;48;56;8588;44"
  parsed_input := strings.Split(string_input, ";")
	var sb strings.Builder
  for _, v := range parsed_input {
    var id, _ = strconv.ParseInt(v, 10, 0)
	  fmt.Println(id)

    var row_strings = []string{
      strconv.Itoa(int(id)),
      "bla",
    }
    sb.WriteString(strings.Join(row_strings, ";"))
    sb.WriteString("|")
  }
	fmt.Println(parsed_input)
	fmt.Println(sb.String())

	fmt.Println("...")

  input := []int{1, 2, 3, 4, 5}
  result := enrich_data(&input)

	fmt.Println(result)
}
