# poe-go

The goal of this library is to provide an entry-level guide to writing tools for PoE in the Go programming language.

It should be simple enough for anyone moderately computer savvy to follow and have their own stash tab indexer running in no time. :moneybag:

## Getting Started

### Prerequisites

Before doing anything else, you'll need to install [Go](https://golang.org/dl/) and [Git](https://git-scm.com/downloads). Both provide installers that you can just spam-click "next" through.

Once both Go and Git are installed, open up "Git Bash" and verify that the "go" and "git" commands are available by typing `go version` and `git version`.

![Go and Git](https://i.imgur.com/Z8jVV7X.png)

### Creating the Project

Go has relatively rigid project organization, so the project directory that we'll be working in needs to be made in your user directory's _go/src_ directory.

Create a new project directory: `mkdir -p ~/go/src/poe-indexing-101`

Make it the current directory: `cd ~/go/src/poe-indexing-101`

![Project Directory](https://i.imgur.com/tf173RC.png)

Create a new file named _main.go_ by copying the contents of [examples/poe-indexing-101/main.go](examples/poe-indexing-101/main.go):

![main.go](https://i.imgur.com/qJb7AoU.png)

Finally, we'll build and run your first indexer.

Install dependencies (This will download and install this library.): `go get -v .`

Run _main.go_: `go run main.go`

![go run](https://i.imgur.com/rXHqln8.png)

Nice! :thumbsup:

### Where to Go from Here

As-is, this example isn't incredibly useful. You probably want to modify the `processStash` function in _main.go_:

```go
func processStash(stash *api.Stash) {
	for _, item := range stash.Items {
		if item.Type == "Ancient Reliquary Key" {
			log.Printf("Ancient Reliquary Key: account = %v, league = %v, note = %v, tab = %v", stash.AccountName, item.League, item.Note, stash.Label)
		}
	}
}
```

You may want to filter by league, show the account's last character name, parse buyouts, play sounds, compose ready-to-whisper messages, etc.

You can refer to [api/item.go](api/item.go) and [api/stash.go](api/stash.go) to see what data is available for you to use.

And if you're new to Go, you should probably read up a bit on [how to write Go code](https://golang.org/doc/code.html).
