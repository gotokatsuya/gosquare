package venues

import (
	"testing"

	"github.com/gotokatsuya/gosquare/dispatcher"
)

func TestExploreVenues(t *testing.T) {
	client := dispatcher.NewClient()
	req := NewExploreRequest()
	req.LatLng = "40.7,-74"
	if v, err := Explore(client, req); err == nil {
		t.Log("Passed", v)
	} else {
		t.Error("Failed", err)
	}
}
