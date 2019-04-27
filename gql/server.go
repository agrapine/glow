package gql

import (
  "github.com/graphql-go/graphql"
  "github.com/graphql-go/handler"
  "log"
  "net/http"
  "strconv"
)

func HelloQuery(p graphql.ResolveParams) (interface{}, error) {
  return "word", nil
}

func _UsersQuery(p graphql.ResolveParams) (interface{}, error) {
  return []string{"test", "another", "and another"}, nil
}

func ServeGql(port int16) {

  fields := graphql.Fields{
    "hello": &graphql.Field{
      Type:    graphql.String,
      Resolve: HelloQuery,
    },
    "users": &graphql.Field{
      Type:    graphql.NewList(graphql.String),
      Resolve: _UsersQuery,
    },
  }
  query := graphql.ObjectConfig{Name: "Query", Fields: fields}
  schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(query)}
  schema, err := graphql.NewSchema(schemaConfig)
  if err != nil {
    log.Fatalf("failed to create new schema, error: %v", err)
  }

  h := handler.New(&handler.Config{
    Schema:     &schema,
    Pretty:     true,
    Playground: true,
  })

  addr := ":" + strconv.Itoa(int(port))
  log.Fatal(http.ListenAndServe(addr, h))
}
