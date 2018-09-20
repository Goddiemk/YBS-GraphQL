package main

import (
	"./schema"
	"github.com/graphql-go/graphql"
	"reflect"
	"testing"
)

type T struct {
	Query    string
	Schema   graphql.Schema
	Expected interface{}
}

var Tests = []T{}

func init() {
	Tests = []T{
		{
			Query: `
			query{
				busdata(id:1){
					Start
				}
			}
		`,
			Schema: schema.Schema,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"busdata": map[string]interface{}{
						"Start": "လှည်းကူးဈေးရှေ့",
					},
				},
			},
		},
		{
			Query: `
			query{
				busline(name:"နတ်စင်"){
					ID
				}
			}
		`,
			Schema: schema.Schema,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"busline": []interface{}{
						map[string]interface{}{
							"ID": 1,
						},
						map[string]interface{}{
							"ID": 2,
						},
					},
				},
			},
		},

		{
			Query: `
				query{
					busdata(id:3){
						Start					
					}
				}
			`,
			Schema: schema.Schema,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"busdata": map[string]interface{}{
						"Start": "ယုဇနဥယျာဉ်မြို့တော်",
					},
				},
			},
		},

		{
			Query: `
				query{
					busstop(id:4){
						Label
					}
				}
			`,
			Schema: schema.Schema,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"busstop": map[string]interface{}{
						"Label": "ဂိတ်ဟောင်း",
					},
				},
			},
		},
	}
}

func TestQuery(t *testing.T) {
	for _, test := range Tests {
		params := graphql.Params{
			Schema:        test.Schema,
			RequestString: test.Query,
		}
		testGraphql(test, params, t)
	}
}

func testGraphql(test T, p graphql.Params, t *testing.T) {
	result := graphql.Do(p)
	if len(result.Errors) > 0 {
		t.Errorf("wrong result, unexpected errors: %v", result.Errors)
	}

	if !reflect.DeepEqual(result, test.Expected) {
		t.Errorf("wrong result, query: %v, graphql result diff: %v", test.Query, schema.Diff(test.Expected, result))
	}
}
