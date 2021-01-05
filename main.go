package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Article struct {
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json: "content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request)  {
	articles := Articles{
		Article{Title: "Health living", Desc: "Consume some of fruits and vegetables", Content: "Keep healthy!"},
	}
	fmt.Println("Endpoint Hit: All Articles endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Homepage Endpoint Hit!")
}

func handleRequests()  {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main()  {
	handleRequests()
}