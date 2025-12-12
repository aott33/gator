package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aott33/gator/internal/database"
	"github.com/google/uuid"
)

func handlerCreateFeed(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("name is required")
	}

	if len(cmd.args) == 1 {
		return errors.New("url is required")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	name := cmd.args[0]
	url := cmd.args[1]
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams {
		ID:			uuid.New(),
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
		Name:		name,
		Url: 		url,
		UserID: 	user.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Feed: %+v created in database\n", feed)

	return nil
}

func handlerGetFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		fmt.Println("Unable to get feeds")
		return err
	}

	for i := range feeds {
		feedName := feeds[i].Name
		feedUrl := feeds[i].Url
		feedUser := feeds[i].Name_2
		fmt.Printf("* %s\n  URL: %s\n  User: %s\n",feedName, feedUrl, feedUser)
	}

	return nil
}
