package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

… // Existing code from above
func handleRequests() {
    
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/all", returnAllArticles)
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
    fmt.Println("Rest API v2.0 - Mux Routers")
    handleRequests()
}
myRouter.HandleFunc("/article/{id}", returnSingleArticle)
type Article struct {
	Id 			int		 `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}
func returnSingleArticle(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    key := vars["id"]
    
    fmt.Fprintf(w, "Key: " + key)
}
func returnSingleArticle(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    key := vars["key"]
    var1 := vars["var1"]
    var2 := vars["var2"]
    
    fmt.Println("Var 1: " + var1)
    fmt.Println("Var 2: " + var2)
    fmt.Fprintf(w, "Key: " + key)
}

func handleRequests() {
    
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/all", returnAllArticles)
    myRouter.HandleFunc("/article/{key}/{var1}/{var2}/", returnSingleArticle)
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}