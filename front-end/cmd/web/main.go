package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

const webPort = "8084"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "test.page.gohtml")
	})

	fmt.Printf("Starting front end service on port %s\n", webPort)
	err := http.ListenAndServe(fmt.Sprintf(":%s", webPort), nil)
	if err != nil {
		log.Panic(err)
	}
}

//go:embed templates
var templateFS embed.FS

func render(w http.ResponseWriter, t string) {

	partials := []string{
		"templates/base.layout.gohtml",
		"templates/header.partial.gohtml",
		"templates/footer.partial.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("templates/%s", t))
	templateSlice = append(templateSlice, partials...)

	tmpl, err := template.ParseFS(templateFS, templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data struct {
		BrokerURL        template.URL
		BrokerHandleURL  template.URL
		BrokerLogGRPCURL template.URL
	}

	data.BrokerURL = template.URL(os.Getenv("BROKER_URL"))
	data.BrokerHandleURL = template.URL(os.Getenv("BROKER_HANDLE_URL"))
	data.BrokerLogGRPCURL = template.URL(os.Getenv("BROKER_LOG_GRPC_URL"))

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
