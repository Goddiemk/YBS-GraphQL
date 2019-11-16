package main

import (
	"./schema"
	"encoding/json"
	"github.com/graphql-go/graphql"
	"github.com/stretchr/testify/assert"
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
				busline(name:"နတ်စင်"){
					ID
					Lines
				}
			}
		`,
			Schema: schema.Schema,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"busline": []interface{}{
						map[string]interface{}{
							"ID":    1,
							"Lines": []string{"1", "46"},
						},
						map[string]interface{}{
							"ID":    2,
							"Lines": []string{"1", "46"},
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
						End
						Route			
					}
				}
			`,
			Schema: schema.Schema,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"busdata": map[string]interface{}{
						"Start": "ယုဇနဥယျာဉ်မြို့တော်",
						"End":   "မင်္ဂလာဈေး",
						"Route": []string{
							"ယုဇနဥယျာဉ်မြို့တော်",
							"ပဲခူးမြစ်လမ်း",
							"ဧရာဝဏ်လမ်း",
							"မောင်းမကန်လမ်း",
							"စည်ပင်လမ်း",
							"တောင်မြောက်လမ်းဆုံ",
							"စံပြဈေး",
							"လေးထောင့်ကန်လမ်း",
							"အရှေ့မြင်းပြိုင်ကွင်းလမ်း",
							"တာမွေအဝိုင်း",
							"ဗညားဒလလမ်း",
							"ယုဇနပလာဇာ",
							"မင်္ဂလာဈေး",
							"ကျိုက္ကဆံလမ်း",
						},
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
	actual, _ := json.Marshal(result)
	expected, _ := json.Marshal(test.Expected)
	assert.Exactly(t, string(actual), string(expected))
}
