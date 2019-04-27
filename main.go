package main

import (
  "./gql"
  "fmt"
)

func main() {
  go gql.ServeGql(8080)
  
  _, _ = fmt.Scanln()
  fmt.Println("finished")
}
