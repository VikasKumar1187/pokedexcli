package pokeapi

import (
	"context"
	"encoding/json"
	"net/http"
)

func (c *Client) ListLocations(ctx context.Context, pageURL *string) (*ShallowLocations, error) {

	url := BaseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return &ShallowLocations{}, err
	}

	// Add request headers - in this case we want to specify that we accept JSON.
	//You could also add any o11y(observability) headers such as trace headers or request IDs.
	req.Header.Add("Accept", "application/json")

	// actually perform the http request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return &ShallowLocations{}, err
	}
	// ALWAYS close the request body - otherwise you end up in a world of pain
	defer resp.Body.Close()

	// Let's unmarshal the response body into a struct
	var shallowLocations ShallowLocations
	err = json.NewDecoder(resp.Body).Decode(&shallowLocations)
	if err != nil {
		return &ShallowLocations{}, err
	}

	return &shallowLocations, nil

}
