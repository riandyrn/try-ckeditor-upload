package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
)

const addr = ":9765"

const (
	envKeyBaseImageURL = "BASE_IMAGE_URL"
	envKeyBucketName   = "BUCKET_NAME"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./web")))

	s3Uploader := s3manager.NewUploader(session.Must(session.NewSession()))
	http.HandleFunc("/upload", handleUpload(s3Uploader))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))

	log.Printf("server is listening on %v", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("unable to serve server due: %v", err)
	}
}

func handleUpload(s3Uploader *s3manager.Uploader) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			reader, _ := r.MultipartReader()
			p, _ := reader.NextPart()
			defer p.Close()

			result, _ := s3Uploader.Upload(&s3manager.UploadInput{
				Bucket: aws.String(os.Getenv(envKeyBucketName)),
				Key:    aws.String(uuid.New().String() + ".jpg"),
				Body:   p,
				ACL:    aws.String("public-read"),
			})

			b, _ := json.Marshal(resp{URL: result.Location})
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Write(b)

		} else if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Write([]byte(""))
		}
	}
}

type resp struct {
	URL string `json:"url"`
}
