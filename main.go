package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Pagedata struct {
	Junk string `json:"junk"`
}

func main() {

	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	http.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", http.FileServer(http.Dir("stylesheets"))))

	http.HandleFunc("/", main_page)

	fmt.Printf("hi\n")

	http.ListenAndServe(":8900", nil)
}

func main_page(the_writer http.ResponseWriter, the_request *http.Request) {

	the_data := Pagedata{
		Junk: "dude that is so cool",
	}

	renderfunction(the_writer, "main.html", the_data)

}

func renderfunction(the_writer http.ResponseWriter, thefilename string, the_data any) {
	template_master, the_error := template.ParseFiles("templates/" + thefilename)

	if the_error != nil {
		http.Error(the_writer, the_error.Error(), http.StatusInternalServerError)
		return
	}

	a_error := template_master.Execute(the_writer, the_data)

	if a_error != nil {
		http.Error(the_writer, the_error.Error(), http.StatusInternalServerError)
		return
	}

}
