package models

import (
	"bookmarks/domain/values"
	"time"

	"github.com/jinzhu/gorm"
)

type URL struct {
	gorm.Model
	keyWord     string `gorm:"unique_index;not null;"`
	title       string `gorm:"not null"`
	url         string `gorm:"not null"`
	lastVisitAt time.Time
}

func NewURL(keyWord, url string) (URL, error) {
	page, err := values.NewWebPage(url)
	if err != nil {
		return URL{}, err
	}

	l := URL{
		keyWord: keyWord,
		title:   page.Title,
		url:     page.Url,
	}

	return l, nil
}

func (u URL) Url() string {
	return u.url
}

func (u URL) KeyWord() string {
	return u.keyWord
}

func (u URL) Title() string {
	return u.title
}
