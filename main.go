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

func main() {

	fmt.Println("hello world")
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}
	fmt.Println("Port", portString)

	router := chi.NewRouter()

	// cors config
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	router.Mount("/v1", v1Router)
	v1Router.Get("/heal", handlerReadiness)
  v1Router.Get("/err",handlerError)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	err := srv.ListenAndServe()

	
	if err != nil {
		log.Fatal(err)
	}

}
