package main

import (
	"context"
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
	"github.com/fly-apps/go-example-tigris/tigris"
)

//go:embed templates/*
var resources embed.FS

var t = template.Must(template.ParseFS(resources, "templates/*"))

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Region": os.Getenv("FLY_REGION"),
		}

		t.ExecuteTemplate(w, "index.html.tmpl", data)
	})

	http.HandleFunc("/sample-upload-tigris", func(w http.ResponseWriter, r *http.Request) {

		res,_ := tigris.UploadText(context.TODO())

		data := map[string]string{
			"Region": res,
		}
		t.ExecuteTemplate(w, "index.html.tmpl", data)

		
	})

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
