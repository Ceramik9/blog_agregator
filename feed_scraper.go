package main

import (
	"fmt"
	"time"
	"context"
	"database/sql"
	"github.com/Ceramik9/blog_agregator/internal/database"
)

func scrapeFeeds(s *state) error {

	// create context
	ctx := context.Background()

	// get feed to scrape
	feedToScrape, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return err
	}

	// mark feed fetched
	timeNow := sql.NullTime {
		Time:  time.Now(),
		Valid: true,
	}
	feedToMark := database.MarkFeedFetchedParams {
		UpdatedAt:     timeNow.Time,
		LastFetchedAt: timeNow,
		ID:            feedToScrape.ID,
	}
	err = s.db.MarkFeedFetched(ctx, feedToMark)
	if err != nil {
		return err
	}

	// fetch feed
	feed, err := fetchFeed(ctx, feedToScrape.Url.String)
	if err != nil {
		return err
	}
	
	// print all feed titles
	for i, item := range feed.Channel.Item {
		fmt.Printf("%d. %s\n", i+1, item.Title)
	}
	fmt.Println("===================================================================================")
	return nil
}







