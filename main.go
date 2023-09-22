package main

import (
    "fmt"
    "net/http"
    "hash/fnv"
    "encoding/base64"
    "github.com/gorilla/mux"
)

var urlMap = make(map[string]string)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", shortenURL).Methods("POST")
    r.HandleFunc("/{shortURL}", redirectURL)

    http.Handle("/", r)
    fmt.Println("Listening on :8080...")
    http.ListenAndServe(":8080", nil)
}

func shortenURL(w http.ResponseWriter, r *http.Request) {
    longURL := r.FormValue("url")
    if longURL == "" {
        http.Error(w, "URL is required", http.StatusBadRequest)
        return
    }

    hash := hashURL(longURL)
    shortURL := base64.StdEncoding.EncodeToString(hash)[:8]
    urlMap[shortURL] = longURL

    fmt.Fprintf(w, "Shortened URL: http://localhost:8080/%s\n", shortURL)
}

func redirectURL(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    shortURL := vars["shortURL"]
    longURL, ok := urlMap[shortURL]
    if !ok {
        http.Error(w, "Short URL not found", http.StatusNotFound)
        return
    }

    http.Redirect(w, r, longURL, http.StatusFound)
}

func hashURL(url string) []byte {
    h := fnv.New32a()
    h.Write([]byte(url))
    return h.Sum(nil)
}