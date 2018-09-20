package resolvers

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"io/ioutil"
	"net/http"
)

type BusStopDataById struct {
	ID       int
	Name     string    `json:"name"`
	Label    string    `json:"label"`
	Road     string    `json:"road"`
	Township string    `json:"township"`
	Alias    []string  `json:"alias"`
	Geo      []float64 `json:"geo"`
}

type BusStops []BusStopDataById

//UnmarshalJSON with arbitary key names
func (bstop *BusStops) UnmarshalJSON(b []byte) error {
	tmp := map[int]BusStopDataById{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}
	var bstoplist BusStops
	for key, value := range tmp {
		value.ID = key
		bstoplist = append(bstoplist, value)
	}
	*bstop = bstoplist
	return nil
}

//GetBusStop function
func GetBusStop(id int) (*BusStopDataById, error) {
	busstop := BusStopDataById{}
	resp, err := http.Get("https://raw.githubusercontent.com/eimg/ybs-data-json/master/bus-stop-data-by-id.js")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s: %s", "could not fetch busstopdatabyid", resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s", "could not read busstopdatabyid")
	}
	busstops := BusStops{}
	err = json.Unmarshal(data, &busstops)
	if err != nil {
		return nil, fmt.Errorf("%s", "could not unmarshal busstopdatabyid")
	}
	for _, value := range busstops {
		if id == value.ID {
			busstop = value
		}
	}
	return &busstop, nil
}

func BusStopResolv(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"]
	v, _ := id.(int)
	return GetBusStop(v)
}
