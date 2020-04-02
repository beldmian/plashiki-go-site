package server

// Translation provides translations for listing
type Translation struct {
	ID      int    `json:"id,omitempty"`
	Quality string `json:"quality,omitempty"`
	Author  string `json:"author,omitempty"`
	Name    string `json:"name,omitempty"`
	URL     string `json:"url,omitempty"`
}

// TranslationFromReq ...
type TranslationFromReq struct {
	ID     int    `json:"id,omitempty"`
	Author string `json:"author,omitempty"`
	URL    string `json:"url,omitempty"`
}

// TranslationResponse ...
type TranslationResponse struct {
	Ok        bool               `json:"ok,omitempty"`
	Result    TranslationFromReq `json:"result,omitempty"`
	ServeTime float32            `json:"serve_time,omitempty"`
}

// Author provides authors for listing
type Author struct {
	Kind         string        `json:"kind,omitempty"`
	Lang         string        `json:"lang,omitempty"`
	MetaTag      string        `json:"meta_tag,omitempty"`
	Name         string        `json:"name,omitempty"`
	Translations []Translation `json:"translations,omitempty"`
}

// AnimeEpisode provides anime parsing from json
type AnimeEpisode struct {
	Authors []Author `json:"authors,omitempty"`
	Players []string `json:"players,omitempty"`
}

// RequestBody provides return of pasrse
type RequestBody struct {
	Ok        bool                    `json:"ok,omitempty"`
	Result    map[string]AnimeEpisode `json:"result,omitempty"`
	ServeTime float32                 `json:"serve_time,omitempty"`
}

// SiteData provide data for site pages
type SiteData struct {
	Title     string
	AnimeData RequestBody
	Anime     Anime
	Number    string
	Jap       string
}

// Anime provides return from search anime on shikimori
type Anime struct {
	ID       int      `json:"id,omitempty"`
	Name     string   `json:"name,omitempty"`
	Russian  string   `json:"russian,omitempty"`
	Japanese []string `json:"japanese,omitempty"`
	Status   string   `json:"status,omitempty"`
	Score    string   `json:"score,omitempty"`
	Image    Image    `json:"image,omitempty"`
	Episodes int      `json:"episodes,omitempty"`
}

// Image provides image type for shikimori serach result
type Image struct {
	Original string `json:"original,omitempty"`
	Preview  string `json:"preview,omitempty"`
}
