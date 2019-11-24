package lib

import "github.com/graphql-go/graphql"

// BaseSchema returns schema for graphql
func BaseSchema() (graphql.Schema, error) {
	busStop := graphql.NewObject(graphql.ObjectConfig{
		Name: "BusStopDataByID",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Label": &graphql.Field{
				Type: graphql.String,
			},
			"Road": &graphql.Field{
				Type: graphql.String,
			},
			"Township": &graphql.Field{
				Type: graphql.String,
			},
			"Alias": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"Geo": &graphql.Field{
				Type: graphql.NewList(graphql.Float),
			},
		},
	})

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"busstop": &graphql.Field{
				Type: busStop,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: BusStopResolver,
			},
		},
	})

	return graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})
}
