package main

import (
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)


type Bahane struct {
	Title  string
}

func showBahane(w http.ResponseWriter, r *http.Request) {
	bahane := getBahane()
	random := rand.Intn(100) + 23
	//fp := path.Join("/templates", "index.html")
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bahan := Bahane{bahane[random]}

	if err := tmpl.Execute(w, bahan); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}


func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", showBahane)

	http.ListenAndServe(":8000", nil)
}





func getBahane() []string {

	fileIO, err := os.OpenFile("file.txt", os.O_RDWR, 0600)
	if err != nil {
		panic(err)
	}
	defer fileIO.Close()
	rawBytes, err := ioutil.ReadAll(fileIO)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(rawBytes), "\n")

	return lines

}

