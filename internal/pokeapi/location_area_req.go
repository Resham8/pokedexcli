package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasRes, error) {
	endpoint := "/location-area"
	fullUrl := baseURL + endpoint

	if pageURL != nil {
		fullUrl = *pageURL
	}

	data, ok := c.cache.Get(fullUrl)

	if ok {		
		locationAreasResp := LocationAreasRes{}
		err := json.Unmarshal(data, &locationAreasResp)

		if err != nil {
			return LocationAreasRes{}, err
		}

		return locationAreasResp, nil
	}	

	req, err := http.NewRequest("GET", fullUrl, nil)

	if err != nil {
		return LocationAreasRes{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasRes{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasRes{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasRes{}, err
	}

	locationAreasResp := LocationAreasRes{}
	err = json.Unmarshal(data, &locationAreasResp)

	if err != nil {
		return LocationAreasRes{}, err
	}

	c.cache.Add(fullUrl, data)

	return locationAreasResp, nil
}
