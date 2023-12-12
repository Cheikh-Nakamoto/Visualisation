package main

import (
	"fmt"
	"visualizations/models"
	"visualizations/handlers"
	"net/http"
)

func main() {

 models.FetchData(models.URL+"artists", &models.All_Artist)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handlers.IndexHandler)

	http.HandleFunc("/info", handlers.InfoHandler)
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
