package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/afrostream/afrostream-go-lib/server"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func main() {
	var s *server.Server

	s = server.New()

	// graphql
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	s.Engine.GET("/", func(c *gin.Context) {
		query := c.Query("query")
		params := graphql.Params{Schema: schema, RequestString: query}
		r := graphql.Do(params)
		if len(r.Errors) > 0 {
			log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
		}
		rJSON, _ := json.Marshal(r)
		c.String(http.StatusOK, fmt.Sprintf("%s", rJSON))
	})
	s.Spawn(CONF_DEFAULT_PORT)
}
