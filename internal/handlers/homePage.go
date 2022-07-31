package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	ascii "ascii-art-web/internal/ascii"
)

type Ascii struct {
	AsciiFont string
	Texted    string
}

var files = []string{
	"./ui/html/home.html",
	"./ui/html/404.html",
	"./ui/html/405.html",
	"./ui/html/500.html",
	"./ui/html/400.html",
}

// GETHandler func receives only GET request and displays main page.
func GETHandler(w http.ResponseWriter, r *http.Request) {
	status := checkMethodAndPath(r, "/", http.MethodGet)
	if status != 200 {
		http.Error(w, fmt.Sprint(status)+" "+http.StatusText(status), status)
		return
	}
	templates(w, "./ui/html/home.html", nil)
}

// Post handler responces only post request and processes data we receive through FromValue
// checks if text is correct and do not contain cyrilic alphabet, correct new lines,checks if font name is correct
// and if file that contains format font haven't been modified through HashSum & ConverFont func.
func POSTHandler(w http.ResponseWriter, r *http.Request) {
	checkMethodAndPath(r, "ascii-art", http.MethodPost)
	text := r.FormValue("formtext") // receives text to format
	font := r.FormValue("fonts")    // receives font's name
	btn := r.FormValue("download")
	html, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "500InternalServerError.html", http.StatusInternalServerError)
		return
	}
	status := CheckUserInput(text, font)
	if status != 200 {
		w.WriteHeader(status)
		html.ExecuteTemplate(w, fmt.Sprint(status)+".html", status) //+http.StatusText(status)
		fmt.Println("error with code status")
		return
	}
	// Converting ascii output result  and saves it in string
	result := ascii.OutputAscii(text, font)
	if btn == "download" {
		w.Header().Set("Content-Disposition", "attachment; filename=ASCII-ART.txt")
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Length", strconv.Itoa(len(result)))
		w.Write([]byte(result))
		return // http.StatusOK
	}
	ascii := Ascii{
		AsciiFont: result,
		Texted:    text,
	}
	templates(w, "./ui/html/home.html", ascii)
}
