package handlers

import (
	"visualizations/models"

	"net/http"
	"strconv"
	"text/template"
)


// The above code is a Go function that handles HTTP requests for the index page and the info page,
// parsing data and rendering templates accordingly.

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" || r.URL.Path == "/handleData" {

		if r.Method == http.MethodGet {
			if len(models.All_Artist) < 1 {
				ParseData(w, 500, "internalServer", "errorBase", nil)
				return
			}
			ParseData(w, 200, "index", "base", models.All_Artist)
		} else {
			ParseData(w, 405, "notAllowed", "errorBase", nil)
		}
	} else {
		ParseData(w, 404, "notFound", "errorBase", nil)
	}
}

func ParseData(w http.ResponseWriter, code int, tmpl, base string, data interface{}) {
	t, err := template.ParseFiles("templates/" + tmpl + ".tmpl")
	if err != nil {
		ParseData(w, 500, "internalServer", "errorBase", nil)
		return
	}
	tmp, err := t.ParseFiles("templates/" + base + ".tmpl")
	if err != nil {
		ParseData(w, 500, "internalServer", "errorBase", nil)
		return
	}
	w.WriteHeader(code)
	tmp.Execute(w, data)
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/info" {
		if r.Method == http.MethodGet {

			id := r.URL.Query().Get("id")
			ID, err := strconv.Atoi(id)
			if len(models.All_Artist) < 1 {
				ParseData(w, 500, "internalServer", "errorBase", nil)
				return
			}
			if err != nil {
				ParseData(w, 400, "badRequest", "errorBase", nil)
				return
			}

			if err = models.SearchId(ID); err != nil {
				ParseData(w, 500, "internalServer", "errorBase", nil)
				return
			}

			if models.ArtistPrint.ID == 0 {
				ParseData(w, 500, "internalServer", "errorBase", nil)
				return
			}

			ParseData(w, 200, "info", "base", models.ArtistPrint)

		} else {
			ParseData(w, 405, "notAllowed", "errorBase", nil)
			return
		}
	} else {
		ParseData(w, 404, "notFound", "errorBase", nil)
		return

	}
}
