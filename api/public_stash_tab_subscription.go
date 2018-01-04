package api

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type PublicStashTabSubscriptionResult struct {
	ChangeId        string
	PublicStashTabs *PublicStashTabs
	Error           error
}

type PublicStashTabSubscription struct {
	Channel      chan PublicStashTabSubscriptionResult
	closeChannel chan bool
	host         string
}

// Opens a subscription that begins with the given change id. To subscribe from the beginning, pass
// an empty string.
func OpenPublicStashTabSubscription(firstChangeId string) *PublicStashTabSubscription {
	return OpenPublicStashTabSubscriptionForHost("www.pathofexile.com", firstChangeId)
}

// Opens a subscription for an alternative host. Can be used for beta or foreign servers.
func OpenPublicStashTabSubscriptionForHost(host, firstChangeId string) *PublicStashTabSubscription {
	ret := &PublicStashTabSubscription{
		Channel:      make(chan PublicStashTabSubscriptionResult),
		closeChannel: make(chan bool),
		host:         host,
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
			response, err := http.Get("https://" + s.host + "/api/public-stash-tabs?id=" + url.QueryEscape(nextChangeId))
			if err != nil {
				s.Channel <- PublicStashTabSubscriptionResult{
					ChangeId: nextChangeId,
					Error:    err,
				}
				continue
			}

			tabs := new(PublicStashTabs)
			decoder := json.NewDecoder(response.Body)
			err = decoder.Decode(tabs)
			if err != nil {
				s.Channel <- PublicStashTabSubscriptionResult{
					ChangeId: nextChangeId,
					Error:    err,
				}
				continue
			}

			if len(tabs.Stashes) > 0 {
				s.Channel <- PublicStashTabSubscriptionResult{
					ChangeId:        nextChangeId,
					PublicStashTabs: tabs,
				}
			}

			nextChangeId = tabs.NextChangeId
		}
	}
}
