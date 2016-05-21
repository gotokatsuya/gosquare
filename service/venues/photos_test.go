package venues

import (
	"testing"

	"github.com/gotokatsuya/gosquare/dispatcher"
)

func TestPhotosVenues(t *testing.T) {
	client := dispatcher.NewClient()
	req := NewPhotosRequest()
	req.VenueID = "43695300f964a5208c291fe3"
	if v, err := Photos(client, req); err == nil {
		t.Log("Passed", v)
	} else {
		t.Error("Failed", err)
	}
}
