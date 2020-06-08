package main

import (
	"html/template"
	"log"
	"net/http"
)

func http_Server(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("HTML_pages/index.html") //parse the html file homepage.html
	if err != nil {                                        // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, nil) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {         // if there is an error
		log.Print("template executing error: ", err) //log it
	}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

}

func main() {
	http.HandleFunc("/", http_Server)
	http.ListenAndServe(":8000", nil)
}
