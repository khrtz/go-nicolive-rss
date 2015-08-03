package main

import (
       "fmt"
        rss "github.com/jteeuwen/go-pkg-rss"
        "os"
)

func main() {
	timeout := 5
	uri := "http://live.nicovideo.jp/recent/rss?tab=live"
	
	feed := rss.New(timeout, true, chanHandler, itemHandler)
	err := feed.Fetch(uri, nil)

	if err != nil {
		fmt.Fprintf(os.Stderr, "[e] %s: %s", uri, err)
		return
	}
}

func chanHandler(feed *rss.Feed, newchannels []*rss.Channel) {
	fmt.Printf("%d new channel(s) in %s\n", len(newchannels), feed.Url)
}

func itemHandler(feed *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
	fmt.Printf("%d new item(s) in %s\n", len(newitems), feed.Url)
	for _, item := range newitems {
		fmt.Println(item.Title)
		fmt.Println("")
	}
}
