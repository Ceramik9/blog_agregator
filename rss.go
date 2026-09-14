package main

import (
	"context"
	"encoding/xml"
	"net/http"
	"html"
	"fmt"
	"io"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	
	// request data
	req, err := http.NewRequest("GET", feedURL, nil)
	if err != nil {
		return nil, err
	}
	
	// set header
	req.Header.Set("User-Agent", "gator")
	
	// request data
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// check response status code
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-OK HTTP status: %v\n", res.Status)
	}

	// read response body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	
	// unmarshal the xml
	var feed RSSFeed
	err = xml.Unmarshal(data, &feed)
	if err != nil {
		return nil, err
	}
	
	// clean up the html
	cleanRSSFeed(&feed)

	return &feed, nil
}

func cleanRSSFeed(feed *RSSFeed) {
	
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	for _, item := range feed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}
}





