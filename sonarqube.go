package sonarqube

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

const (
	defaultBaseURL = "http://sonarqube.kubesphere.com/api/"
	userAgent      = "devops server"
)

type authType int

const (
	basicAuth authType = iota
	oAuthToken
	privateToken
)

type Client struct {
	baseURL            *url.URL
	username, password string
	token              string
	authType           authType
	httpClient         *http.Client
	Users              *UsersService
	Projects           *ProjectService
	Groups             *GroupService
}

func NewClient(endpoint, username, password string) (*Client, error) {
	c := &Client{username: username, password: password, authType: basicAuth, httpClient: http.DefaultClient}
	if endpoint == "" {
		c.SetBaseURL(defaultBaseURL)

	} else {
		if err := c.SetBaseURL(endpoint); err != nil {
			return nil, err
		}
	}
	c.Users = &UsersService{client: c}
	c.Projects = &ProjectService{client: c}
	c.Groups = &GroupService{client: c}
	return c, nil
}

func (c *Client) BaseURL() *url.URL {
	u := *c.baseURL
	return &u
}

// SetBaseURL sets the base URL for API requests to a custom endpoint. urlStr
// should always be specified with a trailing slash.
func (c *Client) SetBaseURL(urlStr string) error {
	if u, err := SetBaseURLUtil(urlStr); err != nil {
		return err
	} else {
		c.baseURL = u
	}
	return nil
}

// NewRequest creates an API request. A relative URL path can be provided in
// urlStr, in which case it is resolved relative to the base URL of the Client.
// Relative URL paths should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, path string, opt interface{}) (*http.Request, error) {
	u := *c.baseURL
	// Set the encoded opaque data
	u.Opaque = c.baseURL.Path + path

	if opt != nil {
		q, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		u.RawQuery = q.Encode()
	}

	req := &http.Request{
		Method:     method,
		URL:        &u,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
		Host:       u.Host,
	}

	if method == "POST" || method == "PUT" {
		// SonarQube use RawQuery even method is POST
		// bodyBytes, err := json.Marshal(opt)
		// if err != nil {
		// 	return nil, err
		// }
		// bodyReader := bytes.NewReader(bodyBytes)

		// u.RawQuery = ""
		// req.Body = ioutil.NopCloser(bodyReader)
		// req.ContentLength = int64(bodyReader.Len())
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")

	switch c.authType {
	case basicAuth, oAuthToken:
		req.SetBasicAuth(c.username, c.password)
	case privateToken:
		req.SetBasicAuth(c.token, "")
	}
	req.Header.Set("User-Agent", userAgent)
	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
func (c *Client) Do(req *http.Request, v interface{}) error {
	return webclient.Do(c.httpClient, req, v)
}
