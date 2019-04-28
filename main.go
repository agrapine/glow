package main

import (
  "./gql"
  "fmt"
)

func main() {

  go gql.Serve(8080)

  _, _ = fmt.Scanln()
  fmt.Println("finished")
}
