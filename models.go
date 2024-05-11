package main

import (
	"time"

	"github.com/Pranay-Pandey/rssagg/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"userId"`
}

func DBUserToUser(DBuser database.AppUser) User {
	return User{
		ID:        DBuser.ID,
		CreatedAt: DBuser.CreatedAt,
		UpdatedAt: DBuser.UpdatedAt,
		Name:      DBuser.Name,
		ApiKey:    DBuser.ApiKey,
	}
}

func DBFeedToFeed(DBFeed database.Feed) Feed {
	return Feed{
		ID:        DBFeed.ID,
		CreatedAt: DBFeed.CreatedAt,
		UpdatedAt: DBFeed.UpdatedAt,
		Name:      DBFeed.Name,
		Url:       DBFeed.Url,
		UserID:    DBFeed.UserID,
	}
}

type Follow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
}

func DBFollowToFollow(Dbdata database.FeedsFollow) Follow {
	return Follow{
		ID:        Dbdata.ID,
		CreatedAt: Dbdata.CreatedAt,
		UpdatedAt: Dbdata.UpdatedAt,
		FeedID:    Dbdata.FeedID,
		UserID:    Dbdata.UserID,
	}
}


func DBFollowsToFollows(Dbdata []database.FeedsFollow) []Follow {
	feeds := []Follow{}
	for _, feed := range(Dbdata) {
		feeds = append(feeds, DBFollowToFollow(feed))
	}

	return feeds
}
