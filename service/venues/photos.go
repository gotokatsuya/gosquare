package venues

import (
	"encoding/json"
	"fmt"

	"github.com/gotokatsuya/gosquare/dispatcher"
	"github.com/gotokatsuya/gosquare/model"
)

type PhotosRequest struct {
	VenueID string `json:"venueId,omitempty"`
	Group   string `json:"group,omitempty"`
	Limit   string `json:"limit,omitempty"`
	Offset  string `json:"offset,omitempty"`
}

func NewPhotosRequest() PhotosRequest {
	return PhotosRequest{}
}

func (req *PhotosRequest) getParams() (map[string]string, error) {
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

type PhotosResponse struct {
	Response PhotosResponseBody `json:"response"`
}

func (resp *PhotosResponse) GetPhotos() (venues []model.Photo) {
	return resp.Response.Photos.Items
}

type PhotosResponseBody struct {
	Photos PhotosResponsePhotos `json:"photos"`
}

type PhotosResponsePhotos struct {
	Count int           `json:"count"`
	Items []model.Photo `json:"items"`
}

func Photos(client dispatcher.Client, req PhotosRequest) (*PhotosResponse, error) {
	params, err := req.getParams()
	if err != nil {
		return nil, err
	}
	body, err := client.DispatchGetRequest(fmt.Sprintf("venues/%s/photos", req.VenueID), params)
	if err != nil {
		return nil, err
	}
	resp := new(PhotosResponse)
	if err := json.Unmarshal(body, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
