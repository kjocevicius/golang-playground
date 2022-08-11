package main

import (
	"fmt"
)

func main() {
	MagicOfUnderlyingArray()
	AppendingExample()
}

func MagicOfUnderlyingArray() {
	fmt.Println("=== MagicOfUnderlyingArray")

	a := make([]interface{}, 0, 3)
	fmt.Printf("Capacity: %d\n", cap(a))
	b := append(a, 1, 2, 3)
	c := append(a, 4, 3, 2)
	d := append(a, 5, 4, 3, 2) // this is too long to fit into underlying array, so new one is allocated
	c[0] = 9
	fmt.Printf("a=%#v \nb=%#v \nc=%#v \nd=%#v \n", a, b, c, d)
	fmt.Printf("\nSame underlying array (b, c): %s", alias(b, c))
	fmt.Printf("\nDifferent underlying array (b, d): %s", alias(b, d))
	fmt.Println("")
}

func AppendingExample() {
	fmt.Println("=== AppendingExample")

	rowLen := 4
	data := [][]interface{}{
		makeRow(rowLen, []interface{}{0, 1, 2}),
		makeRow(rowLen, []interface{}{0, 1, 2}),
		makeRow(rowLen, []interface{}{0, 1, 2}),
	}

	fmt.Printf("Capacity: %d\n", cap(data[0]))
	newData := []interface{}{4.2, 4.3, 4.4}
	newTable := appendColumn(&data, &newData)
	fmt.Printf("Same underlying array: %s\n", alias(data[0], (*newTable)[0]))
	fmt.Printf("Capacity: %d\n", cap((*newTable)[0]))

	// Why underlying array remains the same?
	// I would expect it to change, however it seems to just increase (double) capacity???
	fmt.Printf("Capacity: %d\n", cap(data[0]))
	newTable = appendColumn(&data, &newData)
	fmt.Printf("Same underlying array: %s\n", alias(data[0], (*newTable)[0]))
	fmt.Printf("Capacity: %d\n", cap((*newTable)[0]))

	fmt.Println(data)
}

func appendColumn(data *[][]interface{}, column *[]interface{}) *[][]interface{} {
	for idx, val := range *column {
		(*data)[idx] = append((*data)[idx], val)
	}
	return data
}

func makeRow(capacity int, data []interface{}) []interface{} {
	row := make([]interface{}, 0, capacity)
	fmt.Printf("Capacity initial: %d\n", cap(row))
	row = append(row, data...)
	fmt.Printf("Capacity initialized: %d\n", cap(row))
	return row
}

func alias(x, y []interface{}) bool {
	return cap(x) > 0 && cap(y) > 0 && &x[0:cap(x)][cap(x)-1] == &y[0:cap(y)][cap(y)-1]
}
