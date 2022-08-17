package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

const addr = ":9765"

const envKeyBaseImageURL = "BASE_IMAGE_URL"

func main() {
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.HandleFunc("/upload", handleUpload)
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))

	log.Printf("server is listening on %v", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("unable to serve server due: %v", err)
	}
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		reader, _ := r.MultipartReader()
		p, _ := reader.NextPart()

		srcData, _ := ioutil.ReadAll(p)
		fileName := uuid.New().String() + ".jpg"
		ioutil.WriteFile(fmt.Sprintf("./images/%v", fileName), srcData, 0755)

		b, _ := json.Marshal(resp{URL: fmt.Sprintf("%v/%v", os.Getenv(envKeyBaseImageURL), fileName)})
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(b)

	} else if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Write([]byte(""))
	}
}

type resp struct {
	URL string `json:"url"`
}
