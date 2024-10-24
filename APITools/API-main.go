package APITools

import (
	"encoding/json"
	"net/mail"
	"slices"
	"strings"
)

var st_URL = "http://spacetraders.io/api/v2/"

var factions = []string{
	"COSMIC",
	"VOID",
	"GALACTIC",
	"QUANTUM",
	"DOMINION",
	"ASTRO",
	"CORSIARS",
	"OBSIDIAN",
	"AEGIS",
	"UNITED",
	"SOLITARY",
	"COBALT",
	"OMEGA",
	"ECHO",
	"LORDS",
	"CULT",
	"ANCIENTS",
	"SHADOW",
	"ETHEREAL",
}

type APIToolsErr struct {
	err string
}

func (a APIToolsErr) Error() string {
	return a.err
}

func API_GetStatus() (map[string]interface{}, error) {
	bytes, err := GetRequest(st_URL)

	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	err = json.Unmarshal(bytes, &data)
	return data, err
} // API_GetStatus

func API_RegisterNewAgent(username string, faction string, email string) (map[string]interface{}, error) {
	if len(username) < 3 || len(username) > 14 {
		return nil, APIToolsErr{"Username must be between 3 and 14 characters"}
	}
	if !slices.Contains(factions, faction) {
		return nil, APIToolsErr{"Faction must be one of the following: " + strings.Join(factions, ", ")}
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, APIToolsErr{"Email is not valid"}
	}

	//TODO: build post request
	//need to deal with the case where email is blank (it is allowed to be blank)
	data := map[string]string{
		"username": username,
		"faction":  faction,
		"email":    email,
	}
	jsonPost, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	//convert jsonPost to string
	strData := string(jsonPost)
	resp, err := PostRequest(st_URL+"agents/register", strData)
	if err != nil {
		return nil, err
	}

	var data2 map[string]interface{}
	err = json.Unmarshal(resp, &data2)
	if err != nil {
		return nil, err
	}

	return data2, nil
}
