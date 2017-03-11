package api

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type PublicStashTabSubscriptionResult struct {
	PublicStashTabs *PublicStashTabs
	Error           error
}

type PublicStashTabSubscription struct {
	Channel      chan PublicStashTabSubscriptionResult
	closeChannel chan bool
}

// Opens a subscription that begins with the given change id. To subscribe from the beginning, pass
// an empty string.
func OpenPublicStashTabSubscription(firstChangeId string) *PublicStashTabSubscription {
	ret := &PublicStashTabSubscription{
		Channel:      make(chan PublicStashTabSubscriptionResult),
		closeChannel: make(chan bool),
	}
	go ret.run(firstChangeId)
	return ret
}

func (s *PublicStashTabSubscription) Close() {
	s.closeChannel <- true
}

func (s *PublicStashTabSubscription) run(firstChangeId string) {
	defer close(s.Channel)

	nextChangeId := firstChangeId

	const requestInterval = time.Second
	var lastRequestTime time.Time

	for {
		waitTime := requestInterval - time.Now().Sub(lastRequestTime)
		if waitTime > 0 {
			time.Sleep(waitTime)
		}

		select {
		case <-s.closeChannel:
			return
		default:
			lastRequestTime = time.Now()
			response, err := http.Get("https://www.pathofexile.com/api/public-stash-tabs?id=" + url.QueryEscape(nextChangeId))
			if err != nil {
				s.Channel <- PublicStashTabSubscriptionResult{
					Error: err,
				}
				continue
			}

			tabs := new(PublicStashTabs)
			decoder := json.NewDecoder(response.Body)
			err = decoder.Decode(tabs)
			if err != nil {
				s.Channel <- PublicStashTabSubscriptionResult{
					Error: err,
				}
				continue
			}

			nextChangeId = tabs.NextChangeId

			if len(tabs.Stashes) > 0 {
				s.Channel <- PublicStashTabSubscriptionResult{
					PublicStashTabs: tabs,
				}
			}
		}
	}
}
