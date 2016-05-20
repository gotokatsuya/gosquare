package venues

import (
	"encoding/json"

	"github.com/gotokatsuya/gosquare/dispatcher"
	"github.com/gotokatsuya/gosquare/model"
)

type ExploreRequest struct {
	LatLng         string `json:"ll,omitempty"`
	Near           string `json:"near,omitempty"`
	LatLngAcc      string `json:"llAcc,omitempty"`
	Alt            string `json:"alt,omitempty"`
	AltAcc         string `json:"altAcc,omitempty"`
	Radius         string `json:"radius,omitempty"`
	Section        string `json:"section,omitempty"`
	Query          string `json:"query,omitempty"`
	Limit          string `json:"limit,omitempty"`
	Offset         string `json:"offser,omitempty"`
	Novelty        string `json:"novelty,omitempty"`
	FriendVisits   string `json:"friendVisits,omitempty"`
	Time           string `json:"time,omitempty"`
	Day            string `json:"day,omitempty"`
	VenuePhotos    string `json:"venuePhotos,omitempty"`
	LastVenue      string `json:"lastVenue,omitempty"`
	OpenNow        string `json:"openNow,omitempty"`
	SortByDistance string `json:"sortByDistance,omitempty"`
	Price          string `json:"price,omitempty"`
	Saved          string `json:"saved,omitempty"`
	Specials       string `json:"specials,omitempty"`
}

func NewExploreRequest() ExploreRequest {
	return ExploreRequest{}
}

func (req *ExploreRequest) getParams() (map[string]string, error) {
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

type ExploreResponse struct {
	Response exploreResponseBody `json:"response"`
}

func (resp *ExploreResponse) GetVenues() (venues []model.Venue) {
	for _, group := range resp.Response.Groups {
		for _, item := range group.Items {
			venues = append(venues, item.Venue)
		}
	}
	return venues
}

type exploreResponseBody struct {
	Groups []exploreResponseGroup `json:"groups"`
}

type exploreResponseGroup struct {
	Items []exploreResponseItem `json:"items"`
}

type exploreResponseItem struct {
	Venue model.Venue `json:"venue"`
}

func Explore(client dispatcher.Client, req ExploreRequest) (*ExploreResponse, error) {
	params, err := req.getParams()
	if err != nil {
		return nil, err
	}
	body, err := client.DispatchGetRequest("venues/explore", params)
	if err != nil {
		return nil, err
	}
	resp := new(ExploreResponse)
	if err := json.Unmarshal(body, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
