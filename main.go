package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
)

var (
	titles = []string{
		"Dev",
		"ML",
		"Biz",
		"Sec",
		"QA",
		"Low",
		"Cloud",
		"Prompt",
	}
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("./templates/index.html"))

		rand.Shuffle(len(titles), func(i, j int) { titles[i], titles[j] = titles[j], titles[i] })
		count := rand.Intn(len(titles)-1) + 1
		name := ""
		for i := 0; i < count; i++ {
			name += titles[i]
		}

		data := map[string]string{
			"Name": name + "Ops",
		}
		templ.Execute(w, data)
	})

	log.Println("Server running on 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))

}
