package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var Articles = []Article{
	{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	// add our new DELETE endpoint here
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	// update our global Articles array to include
	// our new Article
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])

	fmt.Fprintf(w, "Key: "+vars["id"])

	value, _ := json.Marshal(Articles[key])
	fmt.Fprintf(w, string(value))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < len(Articles); i++ {

	}
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}
