package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

 func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("duration command needed - eg. 1s, 1m, 1h")
	}
	
	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

	return nil
}

func scrapeFeeds(s *state) error {
	ctx := context.Background()
	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return err
	}

	feedMarked, err := s.db.MarkFeedFetched(ctx, feed.ID)
	if err != nil {
		return err
	}

	feedFetch, err := fetchFeed(ctx, feedMarked.Url)
	if err != nil {
		return err
	}
	
	fmt.Println(feedFetch.Channel.Title)

	for i := range feedFetch.Channel.Item {
		fmt.Println(feedFetch.Channel.Item[i].Title)	
	}


	return nil
}
