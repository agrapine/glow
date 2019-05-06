package query

import (
  "github.com/agrapine/glow/models"
  "github.com/graphql-go/graphql"
)

func Channels() *graphql.Field {
  return &graphql.Field{
    Type:    graphql.NewList(channelType()),
    Resolve: channelsQuery,
  }
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

func channelsQuery(_ graphql.ResolveParams) (interface{}, error) {
  return []models.Channel{
    {Id: "1", Description: "General"},
  }, nil
}
