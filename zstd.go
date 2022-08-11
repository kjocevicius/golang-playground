package main

import (
  "github.com/klauspost/compress/zstd"
  "fmt"
)

var encoder, _ = zstd.NewWriter(nil)
var decoder, _ = zstd.NewReader(nil, zstd.WithDecoderConcurrency(0))

func main() {
  value := `
Some string value that can be zipped
Some string value that can that can be zipped
Some string value that can that can be zipped
Some string value string value that can that can be zipped
Some string value that can that can be zipped
Some string value string value that can that can be zipped
Some string value that can
  `
  bytes_value := []byte(value)
  bytes_value_len := len(bytes_value)
  compressed_value_bytes := Compress(bytes_value)
  compressed_value_bytes_len := len(compressed_value_bytes)
  compressed_value_string := string(compressed_value_bytes)
  compressed_value_string_bytes := []byte(compressed_value_string)
  
  decompressed_value_bytes, err := Decompress(compressed_value_string_bytes)
  if err != nil {
    fmt.Println("FAILED to decompress")
  }
  decompressed_value_string := string(decompressed_value_bytes)
  
  // fmt.Println(bytes_value)
  // fmt.Println(compressed_value_bytes)
  // fmt.Println(compressed_value_string_bytes)
  // fmt.Println(decompressed_value_bytes)
  
  fmt.Println("Decompressed bytes: ", bytes_value_len)
  fmt.Println("Compressed bytes: ", compressed_value_bytes_len)
  fmt.Println("Compressed: ", compressed_value_string)
  fmt.Println("Decompressed: ", decompressed_value_string)
}

func Compress(src []byte) []byte {
    return encoder.EncodeAll(src, make([]byte, 0, len(src)))
} 

func Decompress(src []byte) ([]byte, error) {
    return decoder.DecodeAll(src, nil)
} 