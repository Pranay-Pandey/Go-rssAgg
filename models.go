package main

import (
	"time"
	"database/sql"

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


type Post struct {
	ID          uuid.UUID  `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Title       string     `json:"title"`
	Url         string     `json:"url"`
	Description *string    `json:"description"`
	PublishedAt *time.Time `json:"published_at"`
	FeedID      uuid.UUID  `json:"feed_id"`
}

func databasePostToPost(post database.Post) Post {
	return Post{
		ID:          post.ID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Title:       post.Title,
		Url:         post.Url,
		Description: nullStringToStringPtr(post.Description),
		PublishedAt: nullTimeToTimePtr(post.PublishedAt),
		FeedID:      post.FeedID,
	}
}

func databasePostsToPosts(posts []database.Post) []Post {
	result := make([]Post, len(posts))
	for i, post := range posts {
		result[i] = databasePostToPost(post)
	}
	return result
}

func nullTimeToTimePtr(t sql.NullTime) *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}

func nullStringToStringPtr(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	}
	return nil
}