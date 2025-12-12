package main

import (
	"context"
	"fmt"
)

 func handlerAgg(s *state, cmd command) error {
	rssfeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")		
	if err != nil {
		fmt.Println("Error!\nUnable to process rss feed")
		return nil
	}

	fmt.Printf("%+v\n",rssfeed)
	
	return nil
}
