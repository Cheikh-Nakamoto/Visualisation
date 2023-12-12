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
				image := "static/img/500.png"
				ParseData(w, 500, "internalServer", "errorBase", image)
				return
			}
			ParseData(w, 200, "index", "base", models.All_Artist)
		} else {
			image := "static/img/405.png"
			ParseData(w, 405, "internalServer", "errorBase", image)
		}
	} else {
		image := "static/img/404.png"
		ParseData(w, 404, "notFound", "errorBase", image)
	}
}

func ParseData(w http.ResponseWriter, code int, tmpl, base string, data interface{}) {
	t, err := template.ParseFiles("templates/" + tmpl + ".tmpl")
	if err != nil {
		image := "static/img/500.png"
		ParseData(w, 500, "internalServer", "errorBase", image)
		return
	}
	tmp, err := t.ParseFiles("templates/" + base + ".tmpl")
	if err != nil {
		image := "static/img/500.png"
		ParseData(w, 500, "internalServer", "errorBase", image)
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
				image := "static/img/500.png"
				ParseData(w, 500, "internalServer", "errorBase", image)
				return
			}
			if err != nil {
				image := "static/img/404.png"
				ParseData(w, 400, "badRequest", "errorBase", image)
				return
			}

			if ID <= 0 || ID > len(models.All_Artist) {
				image := "static/img/400.png"
				ParseData(w, 400, "internalServer", "errorBase", image)
				return
			}
			if err = models.SearchId(ID); err != nil {
				image := "static/img/500.png"
				ParseData(w, 500, "internalServer", "errorBase", image)
				return
			}

			ParseData(w, 200, "info", "base", models.ArtistPrint)

		} else {
			image := "static/img/405.png"
			ParseData(w, 405, "internalServer", "errorBase", image)
			return
		}
	} else {
		image := "static/img/404.png"
		ParseData(w, 404, "notFound", "errorBase", image)
		return

	}
}
