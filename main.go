package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
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

func testPostArticles(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "test POST endpoint worked!")
}

func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Homepage Endpoint Hit!")
}

func handleRequests()  {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main()  {
	handleRequests()
}