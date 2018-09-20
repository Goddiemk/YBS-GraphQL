package resolvers

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"io/ioutil"
	"net/http"
)

type BusInfo struct {
	BusName  int
	Start    string   `json:"start"`
	End      string   `json:"end"`
	Route    []string `json:"route"`
	BusStops []string `json:"busStops"`
}

type BusData []BusInfo

//UnmarshalJSON with arbitary key names
func (bdata *BusData) UnmarshalJSON(b []byte) error {
	tmp := map[int]BusInfo{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}
	var bdatalist BusData
	for key, value := range tmp {
		value.BusName = key
		bdatalist = append(bdatalist, value)
	}
	*bdata = bdatalist
	return nil
}

//GetBusInfo function
func GetBusInfo(id int) (*BusInfo, error) {
	businfo := BusInfo{}
	resp, err := http.Get("https://raw.githubusercontent.com/eimg/ybs-data-json/master/bus-stop-names-of-lines.js")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s: %s", "could not fetch busstopnamesoflines", resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s", "could not read busstopnamesoflines")
	}
	busdata := BusData{}
	err = json.Unmarshal(data, &busdata)
	if err != nil {
		return nil, fmt.Errorf("%s", "could not unmarshal busstopnamesoflines")
	}
	for _, value := range busdata {
		if id == value.BusName {
			businfo = value
		}
	}
	return &businfo, nil
}

func BusInfoResolv(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["id"]
	v, _ := id.(int)
	return GetBusInfo(v)
}
