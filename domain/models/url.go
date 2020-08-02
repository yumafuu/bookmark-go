package models

import (
	"bookmarks/domain/values"
	"time"

	"github.com/jinzhu/gorm"
)

type URL struct {
	gorm.Model
	KeyWord     string `gorm:"unique_index;not null;"`
	Title       string `gorm:"not null"`
	Url         string `gorm:"not null"`
	LastVisitAt time.Time
}

func NewURL(keyWord, url string) (URL, error) {
	p, err := values.NewWebPage(url)
	if err != nil {
		return URL{}, err
	}

	l := URL{
		KeyWord: keyWord,
		Title:   p.Title,
		Url:     p.Url,
	}

	return l, nil
}
