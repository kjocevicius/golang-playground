package main

import "fmt"

func main() {
  value := uint32(1749750893)

  value64, err64 := Num64(value)
  if err64 != nil {
    fmt.Println("Failed Num64: %s", err64)
  }
  fmt.Println("Num64: ", value64)
  
  valueU64, errU64 := NumU64(value)
  if errU64 != nil {
    fmt.Println("Failed NumU64: %s", errU64)
  }
  fmt.Println("NumU64: ", valueU64)
  
}

func Num64(n interface{}) (int64, error) {
	switch n := n.(type) {
	case int:
		return int64(n), nil
	case int8:
		return int64(n), nil
	case int16:
		return int64(n), nil
	case int32:
		return int64(n), nil
	case int64:
		return int64(n), nil
  case uint:
		return int64(n), nil
	case uintptr:
		return int64(n), nil
	case uint8:
		return int64(n), nil
	case uint16:
		return int64(n), nil
	case uint32:
		return int64(n), nil
	}
	return 0, fmt.Errorf("Not int: %s", n)
}

func NumU64(n interface{}) (uint64, error) {
	switch n := n.(type) {
	case int:
		return uint64(n), nil
	case int8:
		return uint64(n), nil
	case int16:
		return uint64(n), nil
	case int32:
		return uint64(n), nil
	case int64:
		return uint64(n), nil
	case uint:
		return uint64(n), nil
	case uintptr:
		return uint64(n), nil
	case uint8:
		return uint64(n), nil
	case uint16:
		return uint64(n), nil
	case uint32:
		return uint64(n), nil
	case uint64:
		return uint64(n), nil
	}
	return 0, fmt.Errorf("Not int: %s", n)
}