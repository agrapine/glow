package gql

import "github.com/graphql-go/graphql"

type User struct {
  Id    string `json:"id,omitempty"`
  Name  string `json:"name"`
  Email string `json:"email"`
}

func users() *graphql.Field {
  return &graphql.Field{
    Type:    graphql.NewList(userType()),
    Resolve: usersQuery,
  }
}

func usersQuery(_ graphql.ResolveParams) (interface{}, error) {
  return []User{
    {Id: "1", Name: "krisu", Email: "kristiina@agrapine.com"},
    {Id: "2", Name: "alex", Email: "alexandru@agrapine.com"},
  }, nil
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
