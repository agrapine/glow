package main

import (
  "./query"
  "./server"
  "fmt"
  "github.com/graphql-go/graphql"
  "log"
)

func main() {
  fields :=
    graphql.Fields{
      "users":    query.Users(),
      "channels": query.Channels(),
      "messages": query.Messages(),
    }

  schema, e := graphql.NewSchema(graphql.SchemaConfig{
    Query: graphql.NewObject(graphql.ObjectConfig{
      Name:   "Query",
      Fields: fields})})
  if e != nil {
    log.Fatal(e)
  }

  config := &server.Config{
    Port:   8080,
    Schema: schema,
  }
  go config.Serve()

  _, _ = fmt.Scanln()
  fmt.Println("finished")
}
