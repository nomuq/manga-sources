package myanimelist

import (
	"github.com/nokusukun/jikan2go/manga"
)

func GetByID(id int64) (manga.Manga, error) {
	return manga.GetManga(manga.Manga{MalID: malID})
}
