package query

import (
  "github.com/agrapine/glow/models"
  "fmt"
  "github.com/graphql-go/graphql"
)

func Users() *graphql.Field {
  return &graphql.Field{
    Type:    graphql.NewList(userType()),
    Resolve: usersQuery,
  }
}

func userType() graphql.Type {
  return graphql.NewObject(graphql.ObjectConfig{
    Name: "User",
    Fields: graphql.Fields{
      "id":    &graphql.Field{Type: graphql.String, Description: "id of User"},
      "name":  &graphql.Field{Type: graphql.String, Description: "name of User"},
      "email": &graphql.Field{Type: graphql.String, Description: "email of User"},
    },
  })
}

func usersQuery(p graphql.ResolveParams) (interface{}, error) {
  return []models.User{
    {Id: "1", Name: "krisu", Email: "kristiina@agrapine.com"},
    {Id: "2", Name: fmt.Sprintf("alex is ..."), Email: "alexandru@agrapine.com"},
  }, nil
}
