package mangadex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Manga struct {
	CoverURL    string    `json:"cover_url"`
	Description string    `json:"description"`
	Title       string    `json:"title"`
	AltNames    []string  `json:"alt_names"`
	Artist      string    `json:"artist"`
	Author      string    `json:"author"`
	Status      int64     `json:"status"`
	Genres      []int64   `json:"genres"`
	LastChapter string    `json:"last_chapter"`
	LangName    string    `json:"lang_name"`
	LangFlag    string    `json:"lang_flag"`
	Hentai      int64     `json:"hentai"`
	Links       Links     `json:"links"`
	Chapters    []Chapter `json:"chapters"`
}

type Chapter struct {
	ID        int64    `json:"id"`
	Timestamp int64    `json:"timestamp"`
	Hash      string   `json:"hash"`
	Volume    string   `json:"volume"`
	Chapter   string   `json:"chapter"`
	Title     string   `json:"title"`
	LangName  string   `json:"lang_name"`
	LangCode  string   `json:"lang_code"`
	MangaID   int64    `json:"manga_id"`
	Comments  int64    `json:"comments"`
	Server    string   `json:"server"`
	PageArray []string `json:"page_array"`
	LongStrip int64    `json:"long_strip"`
	Status    string   `json:"status"`
}

type Links struct {
	Al    string `json:"al"`
	Ap    string `json:"ap"`
	Kt    string `json:"kt"`
	Mu    string `json:"mu"`
	Amz   string `json:"amz"`
	Mal   string `json:"mal"`
	Engtl string `json:"engtl"`
}

func GetByID(id string) (Manga, error) {

	type ResponseChapter struct {
		Volume    string `json:"volume"`
		Chapter   string `json:"chapter"`
		Title     string `json:"title"`
		LangCode  string `json:"lang_code"`
		Timestamp int64  `json:"timestamp"`
	}

	type Response struct {
		Manga   Manga                      `json:"manga"`
		Chapter map[string]ResponseChapter `json:"chapter"`
	}

	var response Response
	requestURL := "https://mangadex.org/api/manga/" + id

	res, err := http.DefaultClient.Get(requestURL)
	if err != nil {
		return response.Manga, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return response.Manga, fmt.Errorf("could not get %s: %s", requestURL, res.Status)
	}
	data, err := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(data, &response)
	if err != nil {
		return response.Manga, err
	}

	var chapters []Chapter

	for key, element := range response.Chapter {
		id, err := strconv.ParseInt(key, 10, 32)
		if err == nil {
			chapters = append(chapters, Chapter{
				ID:        id,
				Title:     element.Title,
				Volume:    element.Volume,
				Chapter:   element.Chapter,
				LangCode:  element.LangCode,
				Timestamp: element.Timestamp,
			})
		}
	}

	response.Manga.Chapters = chapters

	return response.Manga, nil
}

func GetChapters(id string) (Chapter, error) {
	var chapter Chapter

	requestURL := "https://mangadex.org/api/?type=chapter&id=" + id

	res, err := http.DefaultClient.Get(requestURL)
	if err != nil {
		return chapter, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return chapter, fmt.Errorf("could not get %s: %s", requestURL, res.Status)
	}
	data, err := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(data, &chapter)
	if err != nil {
		return chapter, err
	}

	return chapter, err
}
