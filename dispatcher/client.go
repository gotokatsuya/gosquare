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
	OAuthToken string
	Version    string
	APIBaseURL string
	APIVersion string
	Locale     string
	Mode       string
	HTTPClient *http.Client
}

const (
	FoursquareMode = "foursquare"
	SwarmMode      = "swarm"
)

const (
	defaultVersion    = "20160520"
	defaultAPIBaseURL = "https://api.foursquare.com"
	defaultAPIVersion = "v2"
	defaultLocale     = "en"
	defaultMode       = FoursquareMode
)

func newDefaultClient() Client {
	return Client{
		Version:    defaultVersion,
		APIBaseURL: defaultAPIBaseURL,
		APIVersion: defaultAPIVersion,
		Locale:     defaultLocale,
		Mode:       defaultMode,
		HTTPClient: http.DefaultClient,
	}
}

func NewClient() Client {
	return NewClientWithParam(util.GetClientID(), util.GetClientSecret())
}

func NewClientWithParam(clientID, clientSecret string) Client {
	client := newDefaultClient()
	client.ID = clientID
	client.Secret = clientSecret
	return client
}

func NewClientWithToken(token string) Client {
	client := newDefaultClient()
	client.OAuthToken = token
	return client
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
	switch {
	case len(c.OAuthToken) != 0:
		values.Set("oauth_token", c.OAuthToken)
	default:
		values.Set("client_id", c.ID)
		values.Set("client_secret", c.Secret)
	}
	values.Set("v", c.Version)
	values.Set("m", c.Mode)

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
