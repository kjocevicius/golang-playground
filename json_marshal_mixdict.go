package main

import (
	"fmt"
  "encoding/json"
)

func main() {
  jsonBlob := []byte(`{
    "b0d_click_itghfjdgythkr": 512,
    "buyd_cliytuyikjm": 512,
    "buyytutyutyu_eur": 512,
    "totaldfgd": false,
    "po-item_fdgdgpressions": 8,
    "porttemfdg_cks": 8,
    "po-itefes": 8,
    "portgh-itemfdgews": 8,
    "total_items_sold_365d_8d": false,
    "selrcentile_dgf365d_8d": false,
    "winsdfg0d_8d": false,
    "totems_lfdgdf65d_8d": false,
    "por64-last-a-stats": 512,
    "aa": true
  }`)
  
  var output map[string]int
  
  // var lines []string
  err := json.Unmarshal(jsonBlob, &output)
  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(output)
}
