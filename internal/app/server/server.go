package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Server ...
type Server struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// New ...
func New(config *Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *Server) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.logger.Info("starting api server")

	s.configureRouter()
	s.logger.Info("router started successful")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *Server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/", s.indexHandler)
	s.router.HandleFunc("/search/{name}", s.searchNameHandler)
	s.router.HandleFunc("/anime/{id}", s.animeIDHandler)
	s.router.HandleFunc("/anime/{id}/{num}", s.animeEpisodeHandler)
	s.router.HandleFunc("/popular", s.popularAnimeHandler)
	s.router.NotFoundHandler = s.router.NewRoute().HandlerFunc(s.handler404).GetHandler()
}

func (s *Server) parseAnime(id string) (RequestBody, Anime) {
	resp, err := http.Get("https://plashiki.online/api/anime/v2/" + id)
	if err != nil {
		s.logger.Warn(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		s.logger.Fatal(err)
	}
	anime := RequestBody{}
	if err := json.Unmarshal(body, &anime); err != nil {
		s.logger.Fatal(err)
	}

	respShiki, err := http.Get("https://shikimori.one/api/animes/" + id)
	if err != nil {
		s.logger.Warn(err)
	}
	bodyShiki, err := ioutil.ReadAll(respShiki.Body)
	if err != nil {
		s.logger.Fatal(err)
	}
	animeShiki := Anime{}
	if err := json.Unmarshal(bodyShiki, &animeShiki); err != nil {
		s.logger.Fatal(err)
	}
	return anime, animeShiki
}
