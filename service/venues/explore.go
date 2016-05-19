package venues

import (
	"encoding/json"

	"github.com/gotokatsuya/gosquare/dispatcher"
	"github.com/gotokatsuya/gosquare/model"
)

type ExploreRequest struct {
	LatLng string `json:"ll"`
	Near   string `json:"near"`
}

func NewExploreRequest(ll string) ExploreRequest {
	return ExploreRequest{
		LatLng: ll,
	}
}

func (req *ExploreRequest) getParams() (params map[string]string) {
	if len(req.LatLng) != 0 {
		params["ll"] = req.LatLng
	}
	if len(req.Near) != 0 {
		params["near"] = req.Near
	}
	return params
}

type ExploreResponse struct {
	Response exploreBody `json:"response"`
}

func (resp *ExploreResponse) GetVenues() (venues []model.Venue) {
	for _, group := range resp.Response.Groups {
		for _, item := range group.Items {
			venues = append(venues, item.Venue)
		}
	}
	return venues
}

type exploreBody struct {
	Groups []exploreGroup `json:"groups"`
}

type exploreGroup struct {
	Items []exploreItem `json:"items"`
}

type exploreItem struct {
	Venue model.Venue `json:"venue"`
}

func Explore(client dispatcher.Client, req ExploreRequest) (*ExploreResponse, error) {
	body, err := client.DispatchGetRequest("venues/explore?", req.getParams())
	if err != nil {
		return nil, err
	}
	resp := new(ExploreResponse)
	if err := json.Unmarshal(body, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
