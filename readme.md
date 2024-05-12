# RSS Aggregator Go - README

## Description

RSS Aggregator Go is a simple RSS feed aggregator written in Go. It fetches RSS feeds from a list of URLs and stores them in a SQLite database. The feeds are then served as a JSON API.

It is a simple project to demonstrate how to build a RESTful API in Go. It uses go routines to fetch feeds concurrently and a Postgres database to store the feeds. It has various endpoints to fetch feeds, add feeds, and delete feeds, follow a feed, and unfollow a feed.

## Installation

1. Clone the repository
2. Run `go mod tidy` to install the dependencies
3. Make a .env file in the root directory and add the following environment variables:
    - `DB_URL` - the URL for the database with all the database table setup
    - `PORT` - the port for the server
4. Run `go run main.go` to start the server
