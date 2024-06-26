package handlers

import (
	"net/http"
	"html/template"
	"strings"
	"io/ioutil"
)

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		text := r.FormValue("text")
		banner := r.FormValue("banner")

		if text == "" || banner == "" {
			http.Error(w, "400 Bad request", http.StatusBadRequest)
			return
		}

		lines, err := readBanner(banner)
		if err!= nil {
			http.Error(w, "500 Internal server error", http.StatusInternalServerError)
			return
		}

		/*text = strings.ReplaceAll(text, "\\n", "\r\n")*/

		// Split the input text into multiple lines
		textLines := strings.Split(text, "\r\n")

		// Initialize the ascii art output
		asciiArt := ""

		// Process each line of input text
		for _, line := range textLines {
			words := strings.Split(line, " ")
			for _, word := range words {
				if word!= "" {
					for i := 0; i < 8; i++ {
						for _, char := range word {
							asciiArt += lines[int(char-' ')*9+1+i] + " "
						}
						asciiArt += "\n"
					}
					asciiArt += "\n"
				}
			}
		}

		tmpl := template.Must(template.ParseFiles("templates/ascii-art.html"))
		err = tmpl.Execute(w, struct {
			Text string
			AsciiArt string
		}{
			Text: text,
			AsciiArt: asciiArt,
		})
		if err!= nil {
			http.Error(w, "500 Internal server error", http.StatusInternalServerError)
		}
	default:
		http.Error(w, "405 Method not allowed", http.StatusMethodNotAllowed)
	}
}


func readBanner(banner string) ([]string, error) {
	path := "banners/" + banner + ".txt"
	data, err := ioutil.ReadFile(path)
	if err!= nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	return lines, nil
}