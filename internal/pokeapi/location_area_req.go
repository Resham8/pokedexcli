package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas() (LocationAreasRes, error) {
	endpoint := "/location-area"
	fullUrl := baseURL + endpoint

	req, err := http.NewRequest("GET", fullUrl, nil)

	if err != nil {
		return LocationAreasRes{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil{
		return  LocationAreasRes{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasRes{}, fmt.Errorf("bad status code: %v",resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasRes{}, err
	}

	loactionAreasResp := LocationAreasRes{}
	err = json.Unmarshal(data,&loactionAreasResp)

	if err != nil {
		return LocationAreasRes{}, err
	}

	return  loactionAreasResp, nil
}