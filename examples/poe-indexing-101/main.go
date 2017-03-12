package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/ccbrown/poe-go/api"
)

// To begin receiving newly updated items immediately, we need to get a recent change id. poe.ninja
// can provide us with that.
func getRecentChangeId() (string, error) {
	resp, err := http.Get("http://api.poe.ninja/api/Data/GetStats")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var stats struct {
		// There are a few more fields in the response, but we only care about this.
		NextChangeId string `json:"nextChangeId"`
	}
	if err := json.Unmarshal(body, &stats); err != nil {
		return "", err
	}

	return stats.NextChangeId, nil
}

// This is where we look through stashes for items of interest. For this example, we'll just log
// reliquary key activity. You might want to parse buyouts, play sounds, compose messages that you
// can send in whispers, etc.
func processStash(stash *api.Stash) {
	for _, item := range stash.Items {
		if item.Type == "Ancient Reliquary Key" {
			log.Printf("Ancient Reliquary Key: account = %v, league = %v, note = %v, tab = %v", stash.AccountName, item.League, item.Note, stash.Label)
		}
	}
}

func main() {
	log.Printf("requesting a recent change id from poe.ninja...")
	recentChangeId, err := getRecentChangeId()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("starting with change id %v", recentChangeId)

	subscription := api.OpenPublicStashTabSubscription(recentChangeId)

	// If we get an interrupt signal (e.g. from control+c on the terminal), handle it gracefully.
	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, os.Interrupt)
		<-ch
		log.Printf("shutting down")
		subscription.Close()
	}()

	// Loop forever over results.
	for result := range subscription.Channel {
		if result.Error != nil {
			log.Printf("error: %v", result.Error.Error())
			continue
		}
		for _, stash := range result.PublicStashTabs.Stashes {
			processStash(&stash)
		}
	}
}
