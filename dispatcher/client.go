package dispatcher

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	clientID     string
	clientSecret string
	v            string
	apiBaseURL   string
	locale       string
	HTTPClient   *http.Client
}

const (
	defaultVersion    = "20160520"
	defaultApiBaseURL = "https://api.foursquare.com/v2/"
	defaultLocale     = "en"
)

func NewClient(clientID, clientSecret string) Client {
	return Client{
		clientID:     clientID,
		clientSecret: clientSecret,
		v:            defaultVersion,
		apiBaseURL:   defaultApiBaseURL,
		locale:       defaultLocale,
		HTTPClient:   http.DefaultClient,
	}
}

func (c *Client) SetVersion(v string) {
	c.v = v
}

func (c *Client) SetAPIBaseURL(apiBaseURL string) {
	c.apiBaseURL = apiBaseURL
}

func (c *Client) SetLocale(locale string) {
	c.locale = locale
}

func (c *Client) SetHTTPClient(httpClient *http.Client) {
	c.HTTPClient = httpClient
}

func (c *Client) DispatchGetRequest(endpoint string, params map[string]string) ([]byte, error) {
	var reqURL bytes.Buffer
	reqURL.WriteString(c.apiBaseURL + endpoint)
	values := url.Values{}
	for k, v := range params {
		values.Set(k, v)
	}
	values.Set("client_id", c.clientID)
	values.Set("client_secret", c.clientSecret)
	values.Set("v", c.v)
	reqURL.WriteString(values.Encode())

	req, err := http.NewRequest("GET", reqURL.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept-Language", c.locale)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= http.StatusOK && resp.StatusCode <= http.StatusBadRequest {
		return body, nil
	}
	return nil, fmt.Errorf("client.getRequest: code:%d body:%s", resp.StatusCode, body)
}
