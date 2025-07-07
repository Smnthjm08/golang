package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OriginalUrl  string    `json:"original_url"`
	ShortenedUrl string    `json:"shortened_url"`
	CreationDate time.Time `json:"creation_date"`
}

// d972323 ->{
// 	ID: "d972323"
// 	OriginalUrl: "https://github.com"
// 	ShortenedUrl: "d972323"
// 	CreationDate: "2025-07-07"
// }

var urlDB = make(map[string]URL)

func generateShortUrl(OriginalUrl string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalUrl))

	data := hasher.Sum(nil)
	hash := hex.EncodeToString(data)
	fmt.Println("hash", hash)
	return hash[:8]
}

func createUrl(OriginalUrl string) string {
	shortURL := generateShortUrl(OriginalUrl)
	id := shortURL
	urlDB[id] = URL{
		ID:           id,
		OriginalUrl:  OriginalUrl,
		ShortenedUrl: shortURL,
		CreationDate: time.Now(),
	}

	fmt.Println("urlDB", urlDB)
	return shortURL
}

func getURL(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")
	}
	return url, nil
}

func RootPageUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
	// return w.Write()
}

func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	shortURL := createUrl(data.URL)

	// fmt.Fprintf(w, shortURL)
	response := struct {
		ShortURL string `json:"short_url"`
	}{ShortURL: shortURL}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func redirectUrlHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]

	url, err := getURL(id)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusNotFound)
		// =====
		return
	}
	fmt.Println("url:", url)

	http.Redirect(w, r, url.OriginalUrl, http.StatusFound)
}

func main() {
	// fmt.Println("starting url shortner...")
	// OriginalUrl := "https://github.com/Smnthjm08/"
	// createUrl(OriginalUrl)

	// register the handler function to handle all requests to the toot URL ("/")
	http.HandleFunc("/", RootPageUrl)
	http.HandleFunc("/url-shortner", ShortURLHandler)
	http.HandleFunc("/redirect/", redirectUrlHandler)

	// start the HTTP server on port 8080
	fmt.Println("Starting server on port 3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error on statrting server:", err)
	}

}
