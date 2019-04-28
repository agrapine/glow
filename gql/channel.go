package gql

import "github.com/graphql-go/graphql"

type Channel struct {
  Id          string `json:"id,omitempty"`
  Description string `json:"description"`
}

func channels() *graphql.Field {
  return &graphql.Field{
    Type:    graphql.NewList(channelType()),
    Resolve: channelsQuery,
  }
}

func channelsQuery(_ graphql.ResolveParams) (interface{}, error) {
  return []Channel{
    {Id: "1", Description: "General"},
  }, nil
}

func channelType() graphql.Type {
  return graphql.NewObject(
    graphql.ObjectConfig{
      Name: "Channel",
      Fields: graphql.Fields{
        "id":          &graphql.Field{Type: graphql.String},
        "description": &graphql.Field{Type: graphql.String},
      },
    })

}
