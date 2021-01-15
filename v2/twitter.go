package twitter

import (
	"net/http"

	"github.com/dghubble/sling"
)

const (
	twitterAPI = "https://api.twitter.com/2/"
)

// Client is a Twitter client for making Twitter 2 API requests.
type Client struct {
	sling  *sling.Sling
	Tweets *TweetsService
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base(twitterAPI)
	return &Client{
		sling:  base,
		Tweets: newTweetsService(base.New()),
	}
}

// Bool returns a new pointer to the given bool value.
func Bool(v bool) *bool {
	ptr := new(bool)
	*ptr = v
	return ptr
}

// Float returns a new pointer to the given float64 value.
func Float(v float64) *float64 {
	ptr := new(float64)
	*ptr = v
	return ptr
}
