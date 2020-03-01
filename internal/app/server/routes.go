package server

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tmpl := template.Must(template.ParseFiles("./internal/app/server/templates/index.html"))
		if err := tmpl.ExecuteTemplate(w, "index", SiteData{Title: "ROMBICK"}); err != nil {
			s.logger.Fatal(err)
		}
	} else {
		r.Header.Add("Status", "303")
		name := r.FormValue("name")
		newURL := "/search/" + name
		http.Redirect(w, r, newURL, http.StatusSeeOther)
	}
}

func (s *Server) animeEpisodeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	num, _ := strconv.Atoi(vars["num"])
	animeData, _ := s.parseAnime(id)
	tmpl := template.Must(template.ParseFiles("./internal/app/server/templates/episode.html"))
	data := struct {
		Title     string
		AnimeData RequestBody
		Number    string
		Previous  int
		Next      int
		ID        string
	}{
		Title:     "Anime id " + id,
		AnimeData: animeData,
		Number:    strconv.Itoa(num),
		Next:      num + 1,
		Previous:  num - 1,
		ID:        id,
	}
	if err := tmpl.ExecuteTemplate(w, "anime", data); err != nil {
		s.logger.Fatal(err)
	}
}

func (s *Server) animeIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	animeData, anime := s.parseAnime(id)
	tmpl := template.Must(template.ParseFiles("./internal/app/server/templates/anime.html"))
	data := SiteData{
		Title:     "Anime id " + id,
		AnimeData: animeData,
		Anime:     anime,
		Jap:       anime.Japanese[0],
	}
	if err := tmpl.ExecuteTemplate(w, "anime", data); err != nil {
		s.logger.Fatal(err)
	}
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
	tmpl := template.Must(template.ParseFiles("./internal/app/server/templates/animelist.html"))
	if err := tmpl.ExecuteTemplate(w, "results", struct {
		Title  string
		Animes []Anime
	}{Title: "Results", Animes: results}); err != nil {
		s.logger.Fatal(err)
	}
}

func (s *Server) popularAnimeHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://shikimori.one/api/animes?order=popularity&limit=30")
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
	tmpl := template.Must(template.ParseFiles("./internal/app/server/templates/animelist.html"))
	if err := tmpl.ExecuteTemplate(w, "results", struct {
		Title  string
		Animes []Anime
	}{Title: "Results", Animes: results}); err != nil {
		s.logger.Fatal(err)
	}
}

func (s *Server) ongoingAnimeHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://shikimori.one/api/animes?order=popularity&limit=30&status=ongoing")
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
	tmpl := template.Must(template.ParseFiles("./internal/app/server/templates/animelist.html"))
	if err := tmpl.ExecuteTemplate(w, "results", struct {
		Title  string
		Animes []Anime
	}{Title: "Results", Animes: results}); err != nil {
		s.logger.Fatal(err)
	}
}

func (s *Server) mostRatedAnimeHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://shikimori.one/api/animes?order=ranked&limit=30")
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
	tmpl := template.Must(template.ParseFiles("./internal/app/server/templates/animelist.html"))
	if err := tmpl.ExecuteTemplate(w, "results", struct {
		Title  string
		Animes []Anime
	}{Title: "Results", Animes: results}); err != nil {
		s.logger.Fatal(err)
	}
}

func (s *Server) movieAnimeHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://shikimori.one/api/animes?kind=movie&order=ranked&limit=30")
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
	tmpl := template.Must(template.ParseFiles("./internal/app/server/templates/animelist.html"))
	if err := tmpl.ExecuteTemplate(w, "results", struct {
		Title  string
		Animes []Anime
	}{Title: "Results", Animes: results}); err != nil {
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
