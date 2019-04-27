package gql

import (
  "../domain"
  "github.com/graphql-go/graphql"
  "github.com/graphql-go/handler"
  "log"
  "net/http"
  "strconv"
)

func _HelloQuery(p graphql.ResolveParams) (interface{}, error) {
  return "word", nil
}

func _UsersQuery(p graphql.ResolveParams) (interface{}, error) {
  return []domain.User{
    {Id: "1", Name: "alex", Email: "alex@home.com"},
  }, nil
}

func _UserType() graphql.Type {
  return graphql.NewObject(graphql.ObjectConfig{
    Name: "User",
    Fields: graphql.Fields{
      "id":    &graphql.Field{Type: graphql.String},
      "name":  &graphql.Field{Type: graphql.String},
      "email": &graphql.Field{Type: graphql.String},
    },
  })
}

func ServeGql(port int16) {

  fields := graphql.Fields{
    "hello": &graphql.Field{
      Type:    graphql.String,
      Resolve: _HelloQuery,
    },
    "users": &graphql.Field{
      Type:    graphql.NewList(_UserType()),
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
