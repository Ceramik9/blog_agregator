package main

import (
	"fmt"
	"time"
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/Ceramik9/blog_agregator/internal/database"
	"strings"
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
	rssFeed, err := fetchFeed(ctx, feedToScrape.Url.String)
	if err != nil {
		return err
	}
	
	// save posts
	for _, item := range rssFeed.Channel.Item {
		var publishedAt sql.NullTime
		parsedTime, err := time.Parse(time.RFC1123, item.PubDate)
		if err == nil {
			publishedAt = sql.NullTime {
				Time: 	parsedTime,
				Valid: true,
			}
		}
		now := time.Now()
		postParams := database.CreatePostParams {
			ID:          uuid.New(),
			CreatedAt:   now,
			UpdatedAt:   now,
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString {
				String: item.Description,
				Valid:  true,
			},
			PublishedAt: publishedAt,
			FeedID:      feedToScrape.ID,
		}
		_, err = s.db.CreatePost(ctx, postParams)
		if err != nil {
			if strings.Contains(err.Error(), "unique constraint") {
				continue
			} else {
				return fmt.Errorf("error saving post:\n%w", err)
			}
		}
	}

	// remove later
	// print all feed titles
	for i, item := range rssFeed.Channel.Item {
		fmt.Printf("%d. %s\n", i+1, item.Title)
	}
	fmt.Println("===================================================================================")
	return nil
}







