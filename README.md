# Golang Blog Aggregator
We're going to build an RSS feed aggregator in Go! 
We'll call it "Gator", you know, because aggreGATOR 🐊.
Anyhow, it's a CLI tool that allows users to:

- Add RSS feeds from across the internet to be collected
- Store the collected posts in a PostgreSQL database
- Follow and unfollow RSS feeds that other users have added
- View summaries of the aggregated posts in the terminal, with a link to the full post
- RSS feeds are a way for websites to publish updates to their content. You can use this project to keep up with your favorite blogs, news sites, podcasts, and more!


## Learning Goals

- Learn how to integrate a Go application with a PostgreSQL database
- Practice using SQL skills to query and migrate a database
- Learn how to write a long-running service that continuously fetches new posts from RSS feeds and stores them in the database

