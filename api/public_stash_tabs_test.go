package api

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sampleJSON []byte

func init() {
	b, err := ioutil.ReadFile("testdata/public-stash-tabs.json")
	if err != nil {
		panic(err)
	}
	sampleJSON = b
}

func TestPublicStashTabs(t *testing.T) {
	var tabs PublicStashTabs
	err := json.Unmarshal(sampleJSON, &tabs)
	assert.NoError(t, err)

	assert.Equal(t, "2300-4136-3306-4292-1278", tabs.NextChangeId)
	assert.NotEmpty(t, tabs.Stashes)
}

func BenchmarkPublicStashTabs(b *testing.B) {
	var tabs PublicStashTabs
	err := json.Unmarshal(sampleJSON, &tabs)
	assert.NoError(b, err)

	assert.Equal(b, "2300-4136-3306-4292-1278", tabs.NextChangeId)
	assert.NotEmpty(b, tabs.Stashes)
}
