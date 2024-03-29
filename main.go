package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/504BassSlapper/rss-agregator/handlers"
	"github.com/504BassSlapper/rss-agregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// build a json rest api server
func main() {
	// load .env file (can be skipped in production mode)
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("port not found in the environement")
	}

	db_url := os.Getenv("DB_HOST_URL")
	if db_url == "" {
		log.Fatal("db url is not found in environement")
	}

	conn, err := sql.Open("postgres", db_url)
	if err != nil {
		log.Fatalln("Cannot connect to db:", err)
	}

	queries := database.New(conn)

	// create api config
	apiCfg := handlers.ApiConfig{
		DB: queries,
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{
			"http://*", "https://*",
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	v1Router := chi.NewRouter()

	// only fire on get requests
	v1Router.Get("/healthz", handlers.HandlerReadiness)

	// allow all Http methods
	// v1Router.HandleFunc("/healthz", handlerReadiness)
	v1Router.Get("/err", handlers.HandleErr)
	v1Router.Post("/users", apiCfg.HandlerCreateUser)
	v1Router.Get("/users", apiCfg.MiddleWareAuth(apiCfg.HandlerGetUser))
	v1Router.Post("/feeds", apiCfg.MiddleWareAuth(apiCfg.HandlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.HandlerGetFeeds)
	v1Router.Post("/feed_follows", apiCfg.MiddleWareAuth(apiCfg.HandlerCreateFeedFollow))
	v1Router.Get("/feed_follows", apiCfg.MiddleWareAuth(apiCfg.HandlerGetFeedFollows))

	// mount router
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	fmt.Printf("Server is starting on port: %s \n", port)
	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(srv)

}
