package dispatcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/gotokatsuya/gosquare/util"
)

type Client struct {
	ID         string
	Secret     string
	Version    string
	APIBaseURL string
	APIVersion string
	Locale     string
	HTTPClient *http.Client
}

const (
	defaultVersion    = "20140715"
	defaultAPIBaseURL = "https://api.foursquare.com"
	defaultAPIVersion = "v2"
	defaultLocale     = "en"
)

func NewClient() Client {
	return Client{
		ID:         util.GetClientID(),
		Secret:     util.GetClientSecret(),
		Version:    defaultVersion,
		APIBaseURL: defaultAPIBaseURL,
		APIVersion: defaultAPIVersion,
		Locale:     defaultLocale,
		HTTPClient: http.DefaultClient,
	}
}

func NewClientWithParam(clientID, clientSecret string) Client {
	return Client{
		ID:         clientID,
		Secret:     clientSecret,
		Version:    defaultVersion,
		APIBaseURL: defaultAPIBaseURL,
		APIVersion: defaultAPIVersion,
		Locale:     defaultLocale,
		HTTPClient: http.DefaultClient,
	}
}

func (c *Client) DispatchGetRequest(endpoint string, params map[string]string) ([]byte, error) {
	u, err := url.Parse(c.APIBaseURL)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(c.APIVersion, endpoint)
	urlString := u.String()

	values := url.Values{}
	for k, v := range params {
		values.Set(k, v)
	}
	values.Set("client_id", c.ID)
	values.Set("client_secret", c.Secret)
	values.Set("v", c.Version)

	req, err := http.NewRequest("GET", urlString+"?"+values.Encode(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Language", c.Locale)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusBadRequest {
		return body, nil
	}
	return nil, fmt.Errorf("Dispatcher.Client.DispatchGetRequest: code:%d body:%s", resp.StatusCode, body)
}
