package lib

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/graphql-go/graphql"
)

type BusStop struct {
	ID       int
	Name     string    `json:"name"`
	Label    string    `json:"label"`
	Road     string    `json:"road"`
	Township string    `json:"township"`
	Alias    []string  `json:"alias"`
	Geo      []float64 `json:"geo"`
}

func GetBusStop(id int) (*BusStop, error) {
	var busstop BusStop
	var err error
	var resp *http.Response

	if resp, err = http.Get("https://raw.githubusercontent.com/eimg/ybs-data-json/master/bus-stop-data-by-id.js"); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data []byte
	if data, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}

	var busstops map[int]BusStop
	if err = json.Unmarshal(data, &busstops); err != nil {
		return nil, err
	}

	var busstoplist []BusStop
	for key, value := range busstops {
		value.ID = key
		busstoplist = append(busstoplist, value)
	}

	for _, value := range busstoplist {
		if id == value.ID {
			busstop = value
		}
	}
	return &busstop, nil
}

func BusStopResolver(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["id"]
	return GetBusStop(id.(int))
}
