package main

import (
	_ "embed"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

//go:embed index.html
var indexHTML string

func main() {
	filePath := "target/manifest.json"
	// Check if the correct number of arguments is provided
	if len(os.Args) == 2 {
		filePath = os.Args[1]
	}

	// Read the file's content into a byte slice
	byteValue, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Define an instance of the root struct
	var manifest WritableManifest

	// Unmarshal the byte slice into the Go struct
	err = json.Unmarshal(byteValue, &manifest)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := manifest.CreateMermaidFCGraph()

		// Parse the HTML template from the file
		tmpl, err := template.New("index").Parse(indexHTML)
		if err != nil {
			log.Fatalf("Failed to parse template file: %v", err)
		}

		// Prepare the template data
		templateData := struct {
			Data string
		}{
			Data: data, // Safely insert JSON into the <script> tag
		}

		// Render the template with the data
		err = tmpl.Execute(w, templateData)
		if err != nil {
			log.Fatalf("Failed to execute template: %v", err)
		}
	})
	http.HandleFunc("/er", func(w http.ResponseWriter, r *http.Request) {
		data := manifest.CreateMermaidERGraph()

		log.Print(data)
		// Parse the HTML template from the file
		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			log.Fatalf("Failed to parse template file: %v", err)
		}

		// Prepare the template data
		templateData := struct {
			Data string
		}{
			Data: data, // Safely insert JSON into the <script> tag
		}

		// Render the template with the data
		err = tmpl.Execute(w, templateData)
		if err != nil {
			log.Fatalf("Failed to execute template: %v", err)
		}
	})

	// Start the server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
