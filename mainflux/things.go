package mainflux

import (
	"encoding/json"
	"net/http"
)

type Thing struct {
	ID       string                 `json:"id"`
	Name     string                 `json:"name"`
	Key      string                 `json:"key"`
	Metadata map[string]interface{} `json:"metadata"`
}

type Things struct {
	Things []Thing `json:"things"`
	Total  int     `json:"total"`
	Offset int     `json:"offset"`
	Limit  int     `json:"limit"`
}

func GetThings(token Token) (Things, error) {
	req, err := http.NewRequest("GET", "http://localhost/things", nil)
	req.Header.Add("Authorization", token.Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Things{}, err
	}

	var things Things
	if err := json.NewDecoder(resp.Body).Decode(&things); err != nil {
		return Things{}, err
	}
	return things, nil
}
