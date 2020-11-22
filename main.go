package main

import (
	"bytes"
	"fmt"
	"github.com/collinux/watermark"
	"github.com/gorilla/mux"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

//GET request for article with ID
func GetArticleWithId(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	file, header, err := r.FormFile("d")
	if err != nil {
		panic(err)
	}
	zd := len(header.Header["Content-Type"])
	zf := len("[image/jpeg]")
	fmt.Printf("File name %s\n", header.Filename)
	fmt.Println("i Ass 1")
	if zd == zf {
		w.WriteHeader(418)
		return
	}

	buf.Reset()
	file1, err := os.Create(header.Filename)
	if err != nil {
		fmt.Println(err)

	}

	img, _, err := image.Decode(file)
	opt := jpeg.Options{
		Quality: 90,
	}

	err = jpeg.Encode(file1, img, &opt)
	if err != nil {
		fmt.Println("i Ass")
	}
	fmt.Println(header.Header["Content-Type"])
	logo := watermark.Watermark{Source: "logoPic/g.png"}
	logo.Apply(header.Filename)

	fmt.Println("Done.")
	filename1 := strings.Replace(header.Filename, ".jpg", "", -4) + "_watermark.jpg"
	//_watermark.jpg
	fmt.Println("Read request: " + filename1)
	file2, err := ioutil.ReadFile(filename1)
	w.Write(file2)
	os.Remove(filename1)
	os.Remove(header.Filename)
	return
}

func main() {
	fmt.Println("REST API V2.0 worked....")
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/logo", GetArticleWithId).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", myRouter))
}
