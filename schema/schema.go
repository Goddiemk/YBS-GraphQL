package schema

import (
	"../resolvers"
	"github.com/graphql-go/graphql"
)

var Schema graphql.Schema

func init() {
	busInfo := graphql.NewObject(graphql.ObjectConfig{
		Name: "BusInfo",
		Fields: graphql.Fields{
			"BusName": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"Start": &graphql.Field{
				Type: graphql.String,
			},
			"End": &graphql.Field{
				Type: graphql.String,
			},
			"Route": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"BusStops": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	})

	busStop := graphql.NewObject(graphql.ObjectConfig{
		Name: "BusStopDataById",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
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

	busLine := graphql.NewObject(graphql.ObjectConfig{
		Name: "BusLine",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},

			"Lines": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	})

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"busdata": &graphql.Field{
				Type: busInfo,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolvers.BusInfoResolv,
			},
			"busline": &graphql.Field{
				Type: graphql.NewList(busLine),
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.BusLineResolv,
			},
			"busstop": &graphql.Field{
				Type: busStop,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolvers.BusStopResolv,
			},
		},
	})

	Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})
}
