package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	srv := &http.Server{
		Addr:    ":8000",
		Handler: route(),
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}

type HomePage struct {
	Count string
}

func route() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		count, err := count("count.txt")
		if err != nil {
			log.Println("Error count function failed: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		homepage := template.Must(template.ParseFiles("tmpl/homepage.html"))
		if err := homepage.Execute(w, &HomePage{
			Count: count,
		}); err != nil {
			log.Println("Error executing template: ", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
	return mux
}

func count(path string) (string, error) {
	file, err := os.Open("count.txt")
	if err != nil {
		return "", err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	if len(data) == 0 {
		data = []byte("0")
	}

	count, _ := strconv.Atoi(string(data))
	count++

	if err := os.WriteFile("count.txt", []byte(strconv.Itoa(count)), 0666); err != nil {
		return "", err
	}

	return strconv.Itoa(count), nil
}
