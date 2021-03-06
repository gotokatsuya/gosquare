# gosquare
  
Foursquare API client library for Go.
  
https://developer.foursquare.com/docs/
  
## Installation

```bash
$ go get github.com/gotokatsuya/gosquare
```

## Usage

```bash
export FOURSQUARE_CLIENT_ID=""
export FOURSQUARE_CLIENT_SECRET=""
```

### Use Service (Recommend)
  
Service is useful that uses request and response entities are defined.
  
```go

import (
    "fmt"
    
	"github.com/gotokatsuya/gosquare/dispatcher"
    "github.com/gotokatsuya/gosquare/service/venues"
)

func VenuesExplore() error {
	client := dispatcher.NewClient()
	req := venues.NewExploreRequest()
	req.LatLng = "40.7,-74"
	res, err := venues.Explore(client, req)
    if err != nil {
		return err
	}
    for _, v := range res.GetVenues() {
	    fmt.Println(v.Name)
    }
    return nil
}
```

### Use Dispatcher

```go

import (
    "fmt"
	"encoding/json"
    
	"github.com/gotokatsuya/gosquare/dispatcher"
 

func VenuesSearch() error {
	client := dispatcher.NewClient()
	params := make(map[string]string)
	params["near"] = "New+Delhi"
	params["intent"] = "browse"
	params["radiu"] = "10000"
	params["limit"] = "10"
	params["query"] = "pizza+hut"
	res, err := client.DispatchGetRequest("venues/search", params)
    if err != nil {
		return err
	}
	var data interface{}
	if err := json.Unmarshal(res, &data); err != nil {
		return err
	}
    fmt.Println(data)
    return nil
}
```