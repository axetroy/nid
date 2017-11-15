package main

import (
  "github.com/axetroy/nid"
  "fmt"
)

func main() {
  nid := nid.New(8)
  fmt.Print(nid)
}
