package twitter

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dghubble/sling"
)

type Tweets struct {
	Data     []Tweet       `json:"data"`
	Errors   []TweetError  `json:"errors"`
	Includes *TweetInclude `json:"includes"`
}

type Tweet struct {
	ID               string                 `json:"id"`
	Attachments      *TweetAttachments      `json:"attachments"`
	AuthorID         string                 `json:"author_id"`
	CreatedAt        string                 `json:"created_at"`
	NonPublicMetrics *TweetNonPublicMetrics `json:"non_public_metrics"`
	OrganicMetrics   *TweetOrganicMetrics   `json:"organic_metrics"`
	PublicMetrics    *TweetPublicMetrics    `json:"public_metrics"`
	Text             string                 `json:"text"`
}

type TweetError struct {
	Detail       string `json:"detail"`
	ResourceID   string `json:"resource_id"`
	ResourceType string `json:"resource_type"`
	Title        string `json:"title"`
	Section      string `json:"section"`
	Type         string `json:"type"`
}

// IDInt64 returns the ID as int64
func (t Tweet) IDInt64() (int64, error) {
	n, err := strconv.ParseInt(t.ID, 10, 64)
	if err != nil {
		return 0, err
	}
	return n, nil
}

// AuthorIDInt64 returns the AuthorID as int64
func (t Tweet) AuthorIDInt64() (int64, error) {
	n, err := strconv.ParseInt(t.AuthorID, 10, 64)
	if err != nil {
		return 0, err
	}
	return n, nil
}

// CreatedAtTime returns the time a tweet was created.
func (t Tweet) CreatedAtTime() (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05.000Z", t.CreatedAt)
}

type TweetAttachments struct {
	MediaKeys []string `json:"media_keys"`
}

type TweetMedia struct {
	DurationMS       int64                  `json:"duration_ms"`
	Height           int64                  `json:"height"`
	MediaKey         string                 `json:"media_key"`
	OrganicMetrics   *VideoOrganicMetrics   `json:"organic_metrics"`
	NonPublicMetrics *VideoNonPublicMetrics `json:"non_public_metrics"`
	PreviewImageURL  string                 `json:"preview_image_url"`
	PublicMetrics    *VideoPublicMetrics    `json:"public_metrics"`
	Type             string                 `json:"type"`
	URL              string                 `json:"url"`
	Width            int64                  `json:"width"`
}

type TweetInclude struct {
	Media *[]TweetMedia `json:"media"`
}

type TweetOrganicMetrics struct {
	ReplyCount        int `json:"reply_count"`
	LikeCount         int `json:"like_count"`
	RetweetCount      int `json:"retweet_count"`
	URLLinkClicks     int `json:"url_link_clicks"`
	UserProfileClicks int `json:"user_profile_clicks"`
	ImpressionCount   int `json:"impression_count"`
}

type TweetNonPublicMetrics struct {
	UserProfileClicks int `json:"user_profile_clicks"`
	ImpressionCount   int `json:"impression_count"`
	URLLinkClicks     int `json:"url_link_clicks"`
}

type TweetPublicMetrics struct {
	RetweetCount int `json:"retweet_count"`
	ReplyCount   int `json:"reply_count"`
	LikeCount    int `json:"like_count"`
	QuoteCount   int `json:"quote_count"`
}

type VideoOrganicMetrics struct {
	ViewCount int `json:"view_count"`
}

type VideoNonPublicMetrics struct {
	ViewCount int `json:"view_count"`
}

type VideoPublicMetrics struct {
	ViewCount int `json:"view_count"`
}

// TweetsService provides methods for accessing Twitter tweets API endpoints.
type TweetsService struct {
	sling *sling.Sling
}

// newTweetsService returns a new TweetsService.
func newTweetsService(sling *sling.Sling) *TweetsService {
	return &TweetsService{
		sling: sling.Path("tweets"),
	}
}

// TweetsShowParams are the parameters for TweetsService.Show
type TweetsShowParams struct {
	IDs                          []int64
	IncludeMedia                 *bool
	IncludeTweetPublicMetrics    *bool
	IncludeTweetNonPublicMetrics *bool
	IncludeTweetOrganicMetrics   *bool
	IncludeMediaPublicMetrics    *bool
	IncludeMediaNonPublicMetrics *bool
	IncludeMediaOrganicMetrics   *bool
}

// Show returns the requested Tweet.
// https://developer.twitter.com/en/docs/twitter-api/tweets
func (s *TweetsService) Show(params *TweetsShowParams) (Tweets, *http.Response, error) {
	tweets := new(Tweets)
	if params == nil {
		return *tweets, nil, nil
	}
	if len(params.IDs) == 0 {
		return *tweets, nil, nil
	}

	type params2 struct {
		IDs         string  `url:"ids,omitempty"`
		TweetFields *string `url:"tweet.fields,omitempty"`
		MediaFields *string `url:"media.fields,omitempty"`
		Expansions  *string `url:"expansions,omitempty"`
	}

	_true := true
	_false := false
	if params.IncludeMedia == nil {
		params.IncludeMedia = &_false
	}
	if params.IncludeTweetPublicMetrics == nil {
		params.IncludeTweetPublicMetrics = &_true
	}
	if params.IncludeTweetNonPublicMetrics == nil {
		params.IncludeTweetNonPublicMetrics = &_false
	}
	if params.IncludeTweetOrganicMetrics == nil {
		params.IncludeTweetOrganicMetrics = &_false
	}
	if params.IncludeMediaPublicMetrics == nil {
		params.IncludeMediaPublicMetrics = &_true
	}
	if params.IncludeMediaNonPublicMetrics == nil {
		params.IncludeMediaNonPublicMetrics = &_false
	}
	if params.IncludeMediaOrganicMetrics == nil {
		params.IncludeMediaOrganicMetrics = &_false
	}

	_params := new(params2)
	_ids := []string{}
	for _, i := range params.IDs {
		_ids = append(_ids, fmt.Sprintf("%v", i))
	}
	_params.IDs = strings.Join(_ids, ",")
	_tweetFields := []string{"id", "created_at", "author_id"}
	if *params.IncludeTweetPublicMetrics {
		_tweetFields = append(_tweetFields, "public_metrics")
	}
	if *params.IncludeTweetNonPublicMetrics {
		_tweetFields = append(_tweetFields, "non_public_metrics")
	}
	if *params.IncludeTweetOrganicMetrics {
		_tweetFields = append(_tweetFields, "organic_metrics")
	}
	if len(_tweetFields) > 0 {
		__tweetFields := strings.Join(_tweetFields, ",")
		_params.TweetFields = &__tweetFields
	}
	_mediaFields := []string{}
	if *params.IncludeMedia {
		_mediaFields = append(_mediaFields, "duration_ms")
		_mediaFields = append(_mediaFields, "height")
		_mediaFields = append(_mediaFields, "preview_image_url")
		_mediaFields = append(_mediaFields, "media_key")
		_mediaFields = append(_mediaFields, "type")
		_mediaFields = append(_mediaFields, "url")
		_mediaFields = append(_mediaFields, "width")
	}
	if *params.IncludeMediaPublicMetrics {
		_mediaFields = append(_mediaFields, "public_metrics")
	}
	if *params.IncludeMediaNonPublicMetrics {
		_mediaFields = append(_mediaFields, "non_public_metrics")
	}
	if *params.IncludeMediaOrganicMetrics {
		_mediaFields = append(_mediaFields, "organic_metrics")
	}
	if len(_mediaFields) > 0 {
		__mediaFields := strings.Join(_mediaFields, ",")
		_params.MediaFields = &__mediaFields
		_expansions := "attachments.media_keys"
		_params.Expansions = &_expansions
	}

	apiError := new(APIError)
	resp, err := s.sling.New().Get("").QueryStruct(_params).Receive(tweets, apiError)

	return *tweets, resp, relevantError(err, *apiError)
}
