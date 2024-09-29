package main

import (
	"html/template"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	http.HandleFunc("/", homeHandler)
	http.Handle("/metrics", promhttp.Handler())

	log.Info("Сервер cv-devops запущен на порту 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Ошибка при запуске сервера cv-devops: ", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("resume.html")
	if err != nil {
		log.Error("Ошибка загрузки шаблона: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		log.Error("Ошибка выполнения шаблона: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
