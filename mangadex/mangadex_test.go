package mangadex

import (
	"testing"

	"github.com/akyoto/assert"
)

func TestGetManga(t *testing.T) {
	manga, err := GetByID("5")

	assert.Equal(t, "Naruto", manga.Title)
	assert.Nil(t, err)
}
