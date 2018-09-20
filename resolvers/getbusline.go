package resolvers

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"io/ioutil"
	"net/http"
	"sort"
)

type BusLine struct {
	ID    int
	Lines []string
}

type BusLines []BusLine

//UnmarshalJSON with arbitary key names
func (bline *BusLines) UnmarshalJSON(b []byte) error {
	tmp := map[int]BusLine{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}
	var blinelist BusLines
	for key, value := range tmp {
		value.ID = key
		blinelist = append(blinelist, value)
	}
	*bline = blinelist
	return nil
}

//GetBusLine function
func GetBusLine(name string) ([]BusLine, error) {
	var id []int
	busline := []BusLine{}
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
	busstop := BusStops{}
	err = json.Unmarshal(data, &busstop)
	if err != nil {
		return nil, fmt.Errorf("%s", "could not unmarshal busstopdatabyid")
	}
	resp, err = http.Get("https://raw.githubusercontent.com/eimg/ybs-data-json/master/lines-of-bus-stops.js")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s: %s", "could not fetch linesofbusstops", resp.Status)
	}
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s", "could not read linesofbusstops")
	}
	buslines := BusLines{}
	err = json.Unmarshal(data, &buslines)
	if err != nil {
		return nil, fmt.Errorf("%s", "could not unmarshal linesofbusstops")
	}
	for _, value := range busstop {
		if name == value.Name {
			id = append(id, value.ID)
		}
	}
	for _, value := range buslines {
		for _, v := range id {
			if v == value.ID {
				busline = append(busline, value)
			}
		}
	}
	sort.Slice(busline, func(i, j int) bool {
		return busline[i].ID < busline[j].ID
	})
	return busline, nil
}

func BusLineResolv(p graphql.ResolveParams) (interface{}, error) {
	name := p.Args["name"]
	v, _ := name.(string)
	return GetBusLine(v)
}
