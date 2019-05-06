package main

import (
  "github.com/agrapine/glow/query"
  "github.com/agrapine/glow/server"
  "github.com/graphql-go/graphql"
  "log"
  "os"
  "os/signal"
  "syscall"
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

  waitForTerminate()
}

func waitForTerminate() {
  osSignal := make(chan os.Signal, 1)
  signal.Notify(osSignal, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

  // Block until we receive our signal.
  <-osSignal

  log.Println("shutting down")
  os.Exit(0)
}