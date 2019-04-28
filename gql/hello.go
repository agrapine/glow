package gql

import "github.com/graphql-go/graphql"

func hello() *graphql.Field {
  return &graphql.Field{
    Type:    graphql.String,
    Resolve: helloQuery,
  }
}

func helloQuery(_ graphql.ResolveParams) (interface{}, error) {
  return "world", nil
}