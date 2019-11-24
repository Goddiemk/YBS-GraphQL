package lib

import (
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/stretchr/testify/assert"
)

var testSchema, _ = BaseSchema()

func TestBusStop(t *testing.T) {
	var testquery = "query{busstop(id:4){Label}}"
	var params = graphql.Params{
		Schema:        testSchema,
		RequestString: testquery,
	}

	result := graphql.Do(params)
	assert.Equal(t, result.Data, map[string]interface{}{"busstop": map[string]interface{}{"Label": "ဂိတ်ဟောင်း"}})
}
