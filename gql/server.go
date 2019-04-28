package gql

import (
  "fmt"
  "github.com/graphql-go/graphql"
  "github.com/graphql-go/handler"
  "log"
  "net/http"
)

func ServeGql(port int16) {

  fields :=
    graphql.Fields{
      "hello": hello(),
      "users": users(),
    }

  schema, e := graphql.NewSchema(graphql.SchemaConfig{
    Query: graphql.NewObject(graphql.ObjectConfig{
      Name:   "Query",
      Fields: fields})})
  if e != nil {
    log.Fatal(e)
  }

  h := handler.New(&handler.Config{
    Schema:     &schema,
    Pretty:     true,
    Playground: true,
  })

  addr := fmt.Sprintf(":%d", port)
  log.Fatal(http.ListenAndServe(addr, h))
}
