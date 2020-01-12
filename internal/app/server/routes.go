package server

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/gorilla/mux"
)

func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tmpl := template.Must(template.ParseFiles("./internal/app/server/templates/index.html"))
		if err := tmpl.ExecuteTemplate(w, "index", SiteData{Title: "Index"}); err != nil {
			s.logger.Fatal(err)
		}
	}
	re := regexp.MustCompile(`\d+`)
	newURL := "/anime/" + string(re.Find([]byte(r.FormValue("id")))) + "/" + r.FormValue("num")
	http.Redirect(w, r, newURL, http.StatusSeeOther)

}

func (s *Server) animeIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	num := vars["num"]
	animeData := s.parseAnime(id)
	tmpl := template.Must(template.ParseFiles("./internal/app/server/templates/anime.html"))
	data := SiteData{
		Title:     "Anime id " + id,
		AnimeData: animeData,
		Number:    num,
	}
	if err := tmpl.ExecuteTemplate(w, "anime", data); err != nil {
		s.logger.Fatal(err)
	}
}

func (s *Server) searchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tmpl := template.Must(template.ParseFiles("./internal/app/server/templates/search.html"))
		if err := tmpl.ExecuteTemplate(w, "search", SiteData{Title: "Search"}); err != nil {
			s.logger.Fatal(err)
		}
	}
	name := r.FormValue("name")
	newURL := "/search/" + name
	http.Redirect(w, r, newURL, http.StatusSeeOther)
}

func (s *Server) searchNameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	resp, err := http.Get("https://shikimori.one/api/animes?search=" + strings.ReplaceAll(name, " ", "%20") + "&limit=20&order=popularity")
	if err != nil {
		s.logger.Warn(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		s.logger.Fatal(err)
	}

	results := []Anime{}
	if err := json.Unmarshal(body, &results); err != nil {
		s.logger.Warn(err)
	}
	tmpl := template.Must(template.ParseFiles("./internal/app/server/templates/searchResults.html"))
	if err := tmpl.ExecuteTemplate(w, "searchRes", struct {
		Title  string
		Animes []Anime
	}{Title: "Index", Animes: results}); err != nil {
		s.logger.Fatal(err)
	}
}

func (s *Server) handler404(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./internal/app/server/templates/404.html"))
	w.WriteHeader(404)
	if err := tmpl.ExecuteTemplate(w, "404", SiteData{Title: "This page not found!"}); err != nil {
		s.logger.Fatal(err)
	}
}
