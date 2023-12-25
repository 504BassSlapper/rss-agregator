package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

// build a json rest api server
func main() {

	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("port not found in the environement")
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
	v1Router.Get("/healthz", handlerReadiness)

	// allow all Http methods
	// v1Router.HandleFunc("/healthz", handlerReadiness)
	v1Router.Get("/err", handleErr)
	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	fmt.Printf("Server is starting on port: %s \n", port)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(srv)

}
