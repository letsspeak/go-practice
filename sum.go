package main

import (
  "fmt"
)

func sum(a []int) int {
  s := 0
  for _,v := range a {
//    fmt.Printf("i = %d, v = %d\n", i, v)
    fmt.Printf("v = %d\n", v)
    s += v
  }
  return s
}

func main() {
  s := sum([]int{1,2,3,4})
  fmt.Printf("%d\n", s)
}
