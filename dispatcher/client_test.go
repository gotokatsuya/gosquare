package dispatcher

import (
	"encoding/json"
	"testing"
)

func TestDispatchGetRequest(t *testing.T) {
	client := NewClient()
	params := make(map[string]string)
	params["near"] = "New+Delhi"
	params["intent"] = "browse"
	params["radiu"] = "10000"
	params["limit"] = "10"
	params["query"] = "pizza+hut"
	if res, err := client.DispatchGetRequest("venues/search", params); err == nil {
		var data interface{}
		json.Unmarshal(res, &data)
		t.Log("Passed", data)
	} else {
		t.Error("Failed", err)
	}
}
