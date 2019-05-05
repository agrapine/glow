package server

import (
  "context"
  "fmt"
  "github.com/graphql-go/graphql"
  "github.com/graphql-go/handler"
  "log"
  "net/http"
)

type configKeyType int
const configKey configKeyType = 42

type Config struct {
  Port   int16
  Schema graphql.Schema
}

func (config *Config) Serve() {
  gqlHandler := handler.New(&handler.Config{
    Schema:     &config.Schema,
    Pretty:     true,
    Playground: true,
  })

  h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    ctx := context.WithValue(r.Context(), configKey, config)
    gqlHandler.ContextHandler(ctx, w, r)
  })

  addr := fmt.Sprintf(":%d", config.Port)
  log.Printf("listening on %s\n", addr)
  log.Fatal(http.ListenAndServe(addr, h))
}

func Of(ctx context.Context) *Config {
  return ctx.Value(configKey).(*Config)
}
