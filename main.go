package main

import (
	"fmt"
	"github.com/collinux/watermark"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Article struct {
	Id string `json:"Id"`
}

type ErrorMessage struct {
	Message string `json:"Message"`
}
type Option struct {
	Message string `json:"Message"`
}

//GET request for /articles
func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hint: getAllArticles worked.....")
	files, err := ioutil.ReadDir("download/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		logo := watermark.Watermark{Source: "logoPic/g.png"}
		logo.Apply("download/" + file.Name())
		fmt.Println(logo)
		if err != nil {
			panic(err)
		}
	}
}

//GET request for article with ID
func GetArticleWithId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I did it")
	//reqBody, _, err := image.Decode(r.Body)
	reqBody1, err := ioutil.ReadAll(r.Body)
	file, err := os.Open(reqBody1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reqBody1)
}

func main() {
	fmt.Println("REST API V2.0 worked....")
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/download", GetAllArticles).Methods("GET")
	myRouter.HandleFunc("/logo", GetArticleWithId).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", myRouter))
}
