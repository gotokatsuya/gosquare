package venues

import (
	"testing"

	"github.com/gotokatsuya/gosquare/dispatcher"
)

func TestSearchVenues(t *testing.T) {
	client := dispatcher.NewClient()
	req := NewSearchRequest()
	req.Near = "New+Delhi"
	if v, err := Search(client, req); err == nil {
		t.Log("Passed", v)
	} else {
		t.Error("Failed", err)
	}
}
