package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Pranay-Pandey/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	fmt.Println("Hello, World!")

	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT environment variable not set")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable not set")
	}

	conn, conn_err := sql.Open("postgres", dbURL)
	if conn_err != nil {
		log.Fatal("Can't connect to the database ", conn_err)
	}

	dbQueries := database.New(conn)
	apiCfg := apiConfig{
		DB: dbQueries,
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/error", handlerErr)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.handlerGetUsers)
	v1Router.Get("/user", apiCfg.authMiddleware(apiCfg.handleGetUserByAPIKEY))
	v1Router.Post("/feed", apiCfg.authMiddleware(apiCfg.handlerCreateFeed))
	v1Router.Delete("/feed/{feedId}", apiCfg.authMiddleware(apiCfg.handlerDeleteFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)
	v1Router.Post("/follow", apiCfg.authMiddleware(apiCfg.handlerCreateFollow))
	v1Router.Get("/follow", apiCfg.authMiddleware(apiCfg.handlerGetFollows))
	v1Router.Delete("/follow/{feedFollowId}", apiCfg.authMiddleware(apiCfg.handlerDeleteFollow))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	const collectionConcurrency = 10
	const collectionInterval = time.Minute
	go startScraping(dbQueries, collectionConcurrency, collectionInterval)

	log.Println("Server listening on port " + portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
