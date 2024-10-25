package APITools

import (
	"encoding/json"
	"fmt"
)

var factionList []string

func API_ListFactions(page int) (*ListFactions, error) {

	u := ST_URL + "factions?limit=20"
	if page > 1 {
		u += fmt.Sprintf("&page=%d", page)
	}
	if page <= 0 {
		return nil, APIToolsErr{"Page must be greater than 0"}
	}
	bytes, err := GetRequest(u)

	if err != nil {
		return nil, err
	}

	var data *ListFactions
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func LoadFactionList() (*ListFactions, error) {
	data, err := API_ListFactions(1)
	if err != nil {
		return nil, err
	}
	num := data.Meta.Total
	np := int(num / 20)
	if np == 1 {
		if len(factionList) == 0 {
			for i := 0; i < len(data.Data); i++ {
				factionList = append(factionList, data.Data[i].Symbol)
			}
		}
		return data, nil
	}
	for i := 2; i < np; i++ {
		new_data, err := API_ListFactions(i)
		if err != nil {
			return nil, err
		}
		data.Data = append(data.Data, new_data.Data[i])
	}
	return data, nil
}

func API_GetFaction(symbol string) (*GetFaction, error) {
	u := ST_URL + "factions/" + symbol
	bytes, err := GetRequest(u)
	if err != nil {
		return nil, err
	}
	var data *GetFaction
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
