package venues

import (
	"encoding/json"

	"github.com/gotokatsuya/gosquare/dispatcher"
	"github.com/gotokatsuya/gosquare/model"
)

type SearchRequest struct {
	LatLng     string `json:"ll,omitempty"`
	Near       string `json:"near,omitempty"`
	LatLngAcc  string `json:"llAcc,omitempty"`
	Alt        string `json:"alt,omitempty"`
	AltAcc     string `json:"altAcc,omitempty"`
	Query      string `json:"query,omitempty"`
	Limit      string `json:"limit,omitempty"`
	Intent     string `json:"intent,omitempty"`
	Radius     string `json:"radius,omitempty"`
	SW         string `json:"sw,omitempty"`
	NE         string `json:"ne,omitempty"`
	CateogryID string `json:"cateogryId,omitempty"`
	URL        string `json:"url,omitempty"`
	ProviderId string `json:"providerId,omitempty"`
	LinkedId   string `json:"linkedId,omitempty"`
}

func NewSearchRequest() SearchRequest {
	return SearchRequest{}
}

func (req *SearchRequest) getParams() (map[string]string, error) {
	params := make(map[string]string)
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &params); err != nil {
		return nil, err
	}
	return params, nil
}

type SearchResponse struct {
	Response searchResponseBody `json:"response"`
}

func (resp *SearchResponse) GetVenues() (venues []model.Venue) {
	return resp.Response.Venues
}

type searchResponseBody struct {
	Venues []model.Venue `json:"venues"`
}

func Search(client dispatcher.Client, req SearchRequest) (*SearchResponse, error) {
	params, err := req.getParams()
	if err != nil {
		return nil, err
	}
	body, err := client.DispatchGetRequest("venues/search", params)
	if err != nil {
		return nil, err
	}
	resp := new(SearchResponse)
	if err := json.Unmarshal(body, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
